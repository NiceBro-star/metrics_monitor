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

func GetSysLoad(hostinfo HOSTINFO) []string {
	cmd := `
		top -b -n 1 | head -n 5 | awk 'NR==1{print $12,$13,$14}'
	`
	host := hostinfo.Host
	port := hostinfo.Port
	username := hostinfo.Username
	password := hostinfo.Password
	stdout, stderr, err := SSHExecPipe(host, port, username, password, cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}
	//fmt.Printf("STDOUT:\n%s\n", stdout)
	fz := strings.Split(stdout, ", ")
	res := make([]string, len(fz))
	for i, line := range fz {
		res[i] = strings.TrimSpace(line)
	}
	//fmt.Print(res)
	return res
}

func GetCPUUsedPercent(hostinfo HOSTINFO) string {
	cmd := `
		top -b -n 1 | head -n 5 | awk 'NR==3{print $2,$4}'
	`
	host := hostinfo.Host
	port := hostinfo.Port
	username := hostinfo.Username
	password := hostinfo.Password
	stdout, stderr, err := SSHExecPipe(host, port, username, password, cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}
	us_sy_used := strings.Split(stdout, " ")
	us_sy_used_trim := make([]string, len(us_sy_used))
	for i, line := range us_sy_used {
		us_sy_used_trim[i] = strings.TrimSpace(line)
	}
	us, _ := strconv.ParseFloat(us_sy_used_trim[0], 64)
	sy, _ := strconv.ParseFloat(us_sy_used_trim[1], 64)
	res := strconv.FormatFloat(us+sy, 'f', 2, 64) + "%"
	return res
}

func GetMEMUsedPercent(hostinfo HOSTINFO) string {
	cmd := `
		top -b -n 1 | head -n 5 | awk 'NR==4{sub(/\+total,/, "", $4); printf "%s ", $4} NR==5{print $9}'
	`
	host := hostinfo.Host
	port := hostinfo.Port
	username := hostinfo.Username
	password := hostinfo.Password
	stdout, stderr, err := SSHExecPipe(host, port, username, password, cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}
	mem_used := strings.Split(stdout, " ")
	mem_used_trim := make([]string, len(mem_used))
	for i, line := range mem_used {
		mem_used_trim[i] = strings.TrimSpace(line)
	}
	total, _ := strconv.ParseFloat(mem_used_trim[0], 64)
	awail, _ := strconv.ParseFloat(mem_used_trim[1], 64)
	mem_calc := (total - awail) / total
	//total_ := strconv.FormatFloat(total, 'f', 2, 64) + "%"
	return strconv.FormatFloat(mem_calc, 'f', 2, 64) + "%"
}

func MonitorTemplate1(hostinfo HOSTINFO) MONITOR_TEMPLATE1 {
	res1 := GetSysLoad(hostinfo)
	res2 := GetCPUUsedPercent(hostinfo)
	res3 := GetMEMUsedPercent(hostinfo)
	res := MONITOR_TEMPLATE1{
		res1,
		res2,
		res3,
	}
	return res
}
