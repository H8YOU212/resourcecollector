package collector

import (
	"context"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/host"
)

type Monitor struct {
	nvmlEnabled bool
	Configuration Configuration
}

func NewMonitor() *Monitor {
	m := &Monitor{}
	return m
}



func (m *Monitor) GPUModule() {

}

func (m *Monitor) CPUModule(ctx context.Context) (CPUInfo, error) {
	cpus, err := cpu.InfoWithContext(ctx)
	if err != nil || len(cpus) == 0{
		return CPUInfo{}, err
	}
	
	cores, _ := cpu.Counts(false)
	threads, _ := cpu.Counts(true)

	percentSlice, _ := cpu.PercentWithContext(ctx, time.Second, false)

	pct := 0.0
	if len(percentSlice) > 0{
		pct = percentSlice[0]
	}

	freqs := []int64{}
	model := cpus[0].ModelName
	flags := cpus[0].Flags

	for _, cpu := range cpus{
		if cpu.Mhz > 0{
			freqs = append(freqs, int64(cpu.Mhz))
		}
	}

	return CPUInfo{
		ModelName: model,
		Cores: cores,
		Threads: threads,
		Percent: pct,
		Frequencies: freqs,
		Flags: flags,
	}, nil
}

func (m *Monitor) RAMModule(ctx context.Context) ( MemoryInfo, error ){
	memoryinfo := m.Configuration.RAMInfo
	vm, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil{
		return MemoryInfo{}, err
	}
	
	memoryinfo.Available = vm.Available
	memoryinfo.Total = vm.Total
	memoryinfo.Used = vm.Used
	memoryinfo.UsedPercent = vm.UsedPercent 

	return memoryinfo, nil 
}

func (m *Monitor) NETModule(ctx context.Context) ( []NetInfo, error ) {
	counters, err := net.IOCountersWithContext(ctx, true)
	if err != nil{
		return nil, err
	}

	var out []NetInfo
	for _, c := range counters{
		out = append(out, NetInfo{
			Name: c.Name,
			BytesSent: c.BytesSent,
			BytesRecv: c.BytesRecv,
			PacketsSent: c.PacketsSent,
			PacketsRecv: c.PacketsRecv,
		})
	}
	return out, nil
}

func (m *Monitor) DiskModule(ctx context.Context) ([]DiskPartition, error) {
	parts, err := disk.PartitionsWithContext(ctx, true)
	if err != nil{
		return nil, err
	}

	var out []DiskPartition

	for _, p := range parts{
		usage, err := disk.UsageWithContext(ctx, p.Mountpoint)
		if err != nil{
			continue
		}
		out = append(out, DiskPartition{
			Device: p.Device,
			Mountpoint: p.Mountpoint,
			Fstype: p.Fstype,
			Total: usage.Total,
			Free: usage.Free,
			Used: usage.Used,
			UsedPct: usage.UsedPercent,
		})
	}

	return out, nil
}

func (m *Monitor) SystemModule(ctx context.Context) (SystemInfo, error) {
	host, err := host.InfoWithContext(ctx)
	if err != nil{
		return SystemInfo{}, err
	}
	return SystemInfo{
		Hostname: host.Hostname,
		OS: host.OS,
		Platform: host.Platform,
		PlatformVer: host.PlatformVersion,
		UptimeSec: uint64(host.Uptime),
	}, err
}


func (m *Monitor) GetTemps(ctx context.Context) ([]TempInfo, error) {
    temps, err := host.SensorsTemperaturesWithContext(ctx)
    if err != nil {
        return nil, nil
    }
    var out []TempInfo
    for _, t := range temps {
        out = append(out, TempInfo{
            Sensor: t.SensorKey,
            Value:  t.Temperature,
        })
    }
    return out, nil
}
