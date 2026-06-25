package api

import (
	"bytes"
	"time"

	"golang.org/x/crypto/ssh"
)

// SSHExecPipe 通过密码认证，启动 bash 执行命令（支持管道）
func SSHExecPipe(host, port, user, password, cmd string) (string, string, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 测试/内网用；生产建议校验 HostKey
		Timeout:         10 * time.Second,
	}

	client, err := ssh.Dial("tcp", host+":"+port, config)
	if err != nil {
		return "", "", err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", "", err
	}
	defer session.Close()

	// 分别捕获 stdout / stderr
	var stdoutBuf, stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	stdin, err := session.StdinPipe()
	if err != nil {
		return "", "", err
	}

	// 启动 bash（交互式 shell）
	if err := session.Start("bash"); err != nil {
		return "", "", err
	}

	// 写入命令 + 退出，完全等价于在终端里手动输入
	stdin.Write([]byte(cmd + "\n"))
	stdin.Write([]byte("exit\n"))
	stdin.Close()

	err = session.Wait()
	return stdoutBuf.String(), stderrBuf.String(), err
}
