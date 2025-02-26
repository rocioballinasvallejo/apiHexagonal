package core

import (
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

type ProcessInfo struct {
	PID      int32   `json:"pid"`
	Name     string  `json:"name"`
	CPUUsage float64 `json:"cpu_usage"`
	MemUsage float64 `json:"mem_usage"`
}

// GetSystemMetrics obtiene las métricas de CPU y RAM
func GetSystemMetrics() (*SystemMetrics, error) {
	// Obtener uso de CPU en un intervalo de 1 segundo
	cpuPercentages, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo uso de CPU: %v", err)
	}

	// Obtener uso de RAM
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo uso de RAM: %v", err)
	}

	// Retornar métricas
	return &SystemMetrics{
		CPUUsage: cpuPercentages[0], // Primer valor del slice
		RAMUsage: vmStat.UsedPercent,
	}, nil
}

// GetSystemProcesses obtiene la lista de procesos que más recursos consumen
func GetSystemProcesses() ([]ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, fmt.Errorf("error obteniendo lista de procesos: %v", err)
	}

	var processInfoList []ProcessInfo
	for _, p := range processes {
		name, err := p.Name()
		if err != nil {
			continue
		}

		cpu, err := p.CPUPercent()
		if err != nil {
			cpu = 0
		}

		mem, err := p.MemoryPercent()
		if err != nil {
			mem = 0
		}

		// Solo incluir procesos que consuman más del 2% de CPU o memoria
		if cpu > 2.0 || mem > 2.0 {
			processInfoList = append(processInfoList, ProcessInfo{
				PID:      p.Pid,
				Name:     name,
				CPUUsage: cpu,
				MemUsage: float64(mem),
			})
		}
	}

	// Ordenar procesos por uso de CPU (de mayor a menor)
	sort.Slice(processInfoList, func(i, j int) bool {
		return processInfoList[i].CPUUsage > processInfoList[j].CPUUsage
	})

	// Limitar a los 5 procesos que más consumen
	if len(processInfoList) > 5 {
		processInfoList = processInfoList[:5]
	}

	return processInfoList, nil
}
