package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/net"
)

func getSystemInfo() map[string]interface{} {
	cpuInfo, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(1, true)
	memoryInfo, _ := mem.VirtualMemory()
	diskUsage, _ := disk.Usage("/")
	netIOCounters, _ := net.IOCounters(false)

	systemInfo := map[string]interface{}{
		"cpu_count": len(cpuInfo),
		"cpu_percent": cpuPercent,
		"memory_info": map[string]interface{}{
			"total":     memoryInfo.Total,
			"available": memoryInfo.Available,
			"percent":   memoryInfo.UsedPercent,
		},
		"disk_usage": map[string]interface{}{
			"total":   diskUsage.Total,
			"used":    diskUsage.Used,
			"free":    diskUsage.Free,
			"percent": diskUsage.UsedPercent,
		},
		"net_io_counters": netIOCounters[0],
	}

	return systemInfo
}

func main() {
	router := gin.Default()
	systemInfo := getSystemInfo()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "We are star stuff which has taken its destiny into its own hands.",
		})
	})

	router.GET("/data", func(c *gin.Context) {
		c.JSON(http.StatusOK, systemInfo)
	})

	router.Run(":81")
}



