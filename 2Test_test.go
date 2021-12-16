package f2_test

import (
	"f2"
	"testing"
)

func TestM1(t *testing.T) {
	f2.ShellCmdTimeout(15, "ping", "-t", "www.baidu.com")
}
