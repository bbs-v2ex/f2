package go_f2

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// 删除输出的\x00和多余的空格
func trimOutput(buffer bytes.Buffer) string {
	return strings.TrimSpace(string(bytes.TrimRight(buffer.Bytes(), "\x00")))
}

// 实时打印输出
func traceOutput(out *bytes.Buffer) {
	offset := 0
	t := time.NewTicker(time.Second * 3)
	defer t.Stop()
	//show := ""
	for {
		<-t.C
		result := bytes.TrimRight((*out).Bytes(), "\x00")
		size := len(result)
		rows := bytes.Split(bytes.TrimSpace(result), []byte{'\n'})
		nRows := len(rows)
		newRows := rows[offset:nRows]

		if result[size-1] != '\n' {
			newRows = rows[offset : nRows-1]
		}

		if len(newRows) < offset {
			continue
		}
		for _, row := range newRows {
			log.Println(string(row))
		}
		offset += len(newRows)
	}
}

// 运行Shell命令，设定超时时间（秒）
func ShellCmdTimeout(timeout int, cmd string, args ...string) (stdout, stderr string, e error) {
	if len(cmd) == 0 {
		e = fmt.Errorf("cannot run a empty command")
		return
	}
	var out, err bytes.Buffer
	command := exec.Command(cmd, args...)
	command.Stdout = &out
	command.Stderr = &err
	command.Start()
	// 启动routine等待结束
	done := make(chan error)
	go func() { done <- command.Wait() }()
	// 启动routine持续打印输出
	go traceOutput(&out)
	// 设定超时时间，并select它
	after := time.After(time.Duration(timeout) * time.Second)
	select {
	case <-after:
		command.Process.Signal(syscall.SIGINT)
		time.Sleep(time.Second)
		command.Process.Kill()
		log.Printf("运行命令（%s）超时，超时设定：%v 秒。",
			fmt.Sprintf(`%s %s`, cmd, strings.Join(args, " ")), timeout)
	case <-done:
	}
	stdout = trimOutput(out)
	stderr = trimOutput(err)
	return
}