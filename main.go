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

	hostList := make([]api.HOSTINFO, 0, 4)
	hostList = append(hostList, host1, host2, host3, host4)

	fmt.Println("IP | LB | CPU | MEM")
	for _, host := range hostList {
		monitorRes := api.MonitorTemplate1(host)

		fmt.Print(host.Host + ": ")
		fmt.Println(monitorRes)
	}
}
