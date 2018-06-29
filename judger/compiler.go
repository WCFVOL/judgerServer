package judger

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Compiler(name string, arg ...string) string {
	proc := exec.Command(name, arg...)
	var out bytes.Buffer
	proc.Stdout = &out
	err := proc.Run()
	if err != nil {
		fmt.Println(err)
	}
	return out.String()
}
