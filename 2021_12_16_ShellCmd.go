package f2

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ShellCmd(name string, arg ...string) (outStr, errStr string, err error) {
	cmd := exec.Command(name, arg...)
	log.Printf(name, arg)
	var stdout, stderr []byte
	var errStdout, errStderr error
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	cmd.Start()
	go func() {
		stdout, errStdout = __2021_12_16_copyAndCapture(os.Stdout, stdoutIn)
	}()
	go func() {
		stderr, errStderr = __2021_12_16_copyAndCapture(os.Stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		//log.Fatalf()
		//err = errors.New("cmd.Run() failed with %s\n", err)
		return
	}
	if errStdout != nil || errStderr != nil {
		err = errors.New("failed to capture stdout or stderr\n")
		//log.Fatalf("failed to capture stdout or stderr\n")
		return
	}
	outStr, errStr = string(stdout), string(stderr)
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
	return
}

func __2021_12_16_copyAndCapture(w io.Writer, r io.Reader) ([]byte, error) {
	var out []byte
	buf := make([]byte, 1024, 1024)
	for {
		n, err := r.Read(buf[:])
		if n > 0 {
			d := buf[:n]
			out = append(out, d...)
			if runtime.GOOS == "windows" {
				d, _ = GbkToUtf8(d)
			}
			os.Stdout.Write(d)
		}
		if err != nil {
			// Read returns io.EOF at the end of file, which is not an error for us
			if err == io.EOF {
				err = nil
			}
			return out, err
		}
	}
	// never reached
	panic(true)
	return nil, nil
}
