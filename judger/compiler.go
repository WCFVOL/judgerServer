package judger

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

//Compiler compiler code
func Compiler(submission Submission) string {
	proc := exec.Command("g++", "/root/user_code/"+strconv.Itoa(submission.Id)+".cpp", "-o", "/root/user_code/"+strconv.Itoa(submission.Id), "-O2")
	var out bytes.Buffer
	proc.Stdout = &out
	err := proc.Run()
	if err != nil {
		fmt.Println(err.Error())
		return "error"
	}
	return "ok"
}
