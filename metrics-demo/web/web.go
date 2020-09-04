package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

var t = metrics.GetOrRegisterTimer("account.create.latency", nil)

func main() {
	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	engine := gin.New()
	// engine.Use(timer)
	engine.GET("/test", timer, test)
	engine.GET("/info", info)

	engine.Run(":8080")
}

func timer(context *gin.Context) {
	t.Time(func() {
		context.Next()
	})
}

func info(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Metrics":     getMetricsInfo(),
		"CpuPercent":  GetCpuPercent(),
		"MemPercent":  GetMemPercent(),
		"DiskPercent": GetDiskPercent(),
		"Check":       "pass",
	})
}

func GetCpuPercent() string {
	percent, _ := cpu.Percent(time.Second, false)
	return fmt.Sprintf("%0.2f", percent[0])
}

func GetMemPercent() string {
	memInfo, _ := mem.VirtualMemory()
	return fmt.Sprintf("%0.2f", memInfo.UsedPercent)
}

func GetDiskPercent() string {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return fmt.Sprintf("%0.2f", diskInfo.UsedPercent)
}

func getMetricsInfo() MetricsInfo {
	return MetricsInfo{
		Count:        t.Count(),
		Min:          time.Duration(t.Min()) / time.Millisecond,
		Max:          time.Duration(t.Max()) / time.Millisecond,
		Mean:         time.Duration(t.Mean()) / time.Millisecond,
		The75:        time.Duration(t.Percentile(0.75)) / time.Millisecond,
		The95:        time.Duration(t.Percentile(0.95)) / time.Millisecond,
		The99:        time.Duration(t.Percentile(0.99)) / time.Millisecond,
		The999:       time.Duration(t.Percentile(0.999)) / time.Millisecond,
		The1MinRate:  fmt.Sprintf("%0.2f", t.Rate1()),
		The5MinRate:  fmt.Sprintf("%0.2f", t.Rate5()),
		The15MinRate: fmt.Sprintf("%0.2f", t.Rate15()),
	}
}

func test(context *gin.Context) {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5) + 1
	fmt.Println(n)
	time.Sleep(time.Second * time.Duration(n))

	context.JSON(http.StatusOK, gin.H{"msg": "success",
		"sleep": n,
	})
}

type MetricsInfo struct {
	Count        int64         `json:"Count"`
	Min          time.Duration `json:"Min"`
	Max          time.Duration `json:"Max"`
	Mean         time.Duration `json:"Mean"`
	The75        time.Duration `json:"75%"`
	The95        time.Duration `json:"95%"`
	The99        time.Duration `json:"99%"`
	The999       time.Duration `json:"99.9%"`
	The1MinRate  string        `json:"1-min rate"`
	The15MinRate string        `json:"15-min rate"`
	The5MinRate  string        `json:"5-min rate"`
}
