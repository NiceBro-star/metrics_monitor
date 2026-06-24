package main

import (
	"fmt"
	"metrics_monitor/api"
)

func main() {
	// 示例：带管道、awk、grep 的复杂命令
	cmd := `
		uname -a 2>1
	`

	stdout, stderr, err := api.SSHExecPipe("172.21.72.2", "root", "Admin@9000", cmd)
	if err != nil {
		fmt.Printf("执行出错: %v\n", err)
	}
	if stderr != "" {
		fmt.Printf("STDERR:\n%s\n", stderr)
	}
	fmt.Printf("STDOUT:\n%s\n", stdout)
}
