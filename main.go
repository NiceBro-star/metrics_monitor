package main

import (
	"fmt"
	"metrics_monitor/api"
)

func main() {
	host1 := api.HOST1
	host2 := api.HOST2
	host3 := api.HOST3
	host4 := api.HOST4

	monitorRes1 := api.MonitorTemplate1(host1)
	monitorRes2 := api.MonitorTemplate1(host2)
	monitorRes3 := api.MonitorTemplate1(host3)
	monitorRes4 := api.MonitorTemplate1(host4)

	fmt.Println("IP | LB | CPU | MEM")
	fmt.Print(host1.Host + ": ")
	fmt.Println(monitorRes1)
	fmt.Print(host2.Host + ": ")
	fmt.Println(monitorRes2)
	fmt.Print(host3.Host + ": ")
	fmt.Println(monitorRes3)
	fmt.Print(host4.Host + ": ")
	fmt.Println(monitorRes4)
}
