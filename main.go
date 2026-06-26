package main

import (
	"fmt"
	"metrics_monitor/api"
	"sync"
)

func main() {
	fmt.Println("IP | LB | CPU | MEM")
	var wg sync.WaitGroup
	for _, host := range api.HOSTS {
		wg.Add(1)
		go func(h api.HOSTINFO) {
			defer wg.Done()
			monitorRes := api.MonitorTemplate1(h)
			fmt.Print(h.Host + ": ")
			fmt.Println(monitorRes)
		}(host)
	}
	wg.Wait()
}
