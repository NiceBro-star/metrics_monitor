package api

import (
	"fmt"
	"strconv"
	"strings"
)

type HOSTINFO struct {
	Host     string
	Port     string
	Username string
	Password string
}

type MONITOR_TEMPLATE1 struct {
	LoadBalance []string
	CPUUsed     string
	MEMUsed     string
}

func GetSysLoad(hostInfo HOSTINFO) []string {
	cmd := `
		uptime | sed 's/.*load average: //'
	`
	host := hostInfo.Host
	port := hostInfo.Port
	username := hostInfo.Username
	password := hostInfo.Password
	stdout, stderr, err := SSHExecPipe(host, port, username, password, cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}

	fz := strings.Split(stdout, ", ")
	res := make([]string, len(fz))
	for i, line := range fz {
		res[i] = strings.TrimSpace(line)
	}

	return res
}

func GetCPUUsedPercent(hostInfo HOSTINFO) string {
	cmd := `
		top -b -n 1 | head -n 5 | awk 'NR==3{print $2,$4}'
	`
	host := hostInfo.Host
	port := hostInfo.Port
	username := hostInfo.Username
	password := hostInfo.Password
	stdout, stderr, err := SSHExecPipe(host, port, username, password, cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}
	UsSyUsed := strings.Split(stdout, " ")
	UsSyUsedTrimSpace := make([]string, len(UsSyUsed))
	for i, line := range UsSyUsed {
		UsSyUsedTrimSpace[i] = strings.TrimSpace(line)
	}
	us, _ := strconv.ParseFloat(UsSyUsedTrimSpace[0], 64)
	sy, _ := strconv.ParseFloat(UsSyUsedTrimSpace[1], 64)
	res := strconv.FormatFloat(us+sy, 'f', 2, 64) + "%"
	return res
}

func GetMEMUsedPercent(hostInfo HOSTINFO) string {
	cmd := `
		free -b | awk 'NR==2{print $2, $7}'
	`
	host := hostInfo.Host
	port := hostInfo.Port
	username := hostInfo.Username
	password := hostInfo.Password
	stdout, stderr, err := SSHExecPipe(host, port, username, password, cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}
	memUsed := strings.Split(stdout, " ")
	memUsedTrimSpace := make([]string, len(memUsed))
	for i, line := range memUsed {
		memUsedTrimSpace[i] = strings.TrimSpace(line)
	}
	totalMem, _ := strconv.ParseFloat(memUsedTrimSpace[0], 64)
	availMem, _ := strconv.ParseFloat(memUsedTrimSpace[1], 64)
	usedMemPercent := (totalMem - availMem) / totalMem * 100

	return strconv.FormatFloat(usedMemPercent, 'f', 2, 64) + "%"
}

func MonitorTemplate1(hostInfo HOSTINFO) MONITOR_TEMPLATE1 {
	res := MONITOR_TEMPLATE1{
		GetSysLoad(hostInfo),
		GetCPUUsedPercent(hostInfo),
		GetMEMUsedPercent(hostInfo),
	}
	return res
}
