package judger

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

func Compiler(submission Submission) string {
	proc := exec.Command("g++","user_code/"+strconv.Itoa(submission.Id)+".cpp","-o","user_code/"+strconv.Itoa(submission.Id),"-O2")
	var out bytes.Buffer
	proc.Stdout = &out
	err := proc.Run()
	if err != nil {
		fmt.Println(err.Error())
		return "error"
	}
	return "ok"
}
