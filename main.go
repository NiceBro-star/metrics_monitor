package main

import (
	"fmt"
	"metrics_monitor/api"
)

func main() {
	host1 := api.HOSTINFO{
		Host:     "172.21.72.2",
		Port:     "22",
		Username: "root",
		Password: "Admin@9000",
	}

	monitorRes := api.MonitorTemplate1(host1)
	fmt.Println(monitorRes.LoadBalance)
	fmt.Println(monitorRes.CPUUsed)
	fmt.Println(monitorRes.MEMUsed)

}
