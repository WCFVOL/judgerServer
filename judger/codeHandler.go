package judger

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"crypto/md5"
)

type Submission struct {
	Id        int
	Code      string
	UserId    int
	Date      int64
	ProblemId int
	Language  int
}

func Handler(str string) {
	var submission Submission
	err := json.Unmarshal([]byte(str), &submission)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(str)
	log.Print(submission)
	WriteCode(submission) // 根据submissionId将代码写入相应文件
	compilerResult := Compiler(submission)
	if compilerResult == "error" {
		//TODO compiler error
		return
	}
	judgeResult := Judger(1000, 2000, 128*1024*1024, 200, 10000, 32*1024*1024, 0, 0, 0,
		"/root/user_code/"+strconv.Itoa(submission.Id), "/root/problem_in/"+strconv.Itoa(submission.ProblemId)+".in", "/root/user_result/"+strconv.Itoa(submission.Id)+".out", "/root/user_result/"+strconv.Itoa(submission.Id)+".error", "judger.log", "c_cpp",
		[]string{""}, []string{"foo=bar"})
	log.Println(judgeResult)
	outputByte, err := ioutil.ReadFile("/root/user_result/" + strconv.Itoa(submission.Id) + ".out")
	stdoutByte, err := ioutil.ReadFile("/root/std_result/" + strconv.Itoa(submission.ProblemId) + ".out")

	md5 := md5.New()
	md5.Write(outputByte)
	output := string(md5.Sum(nil))
	md5.Reset()
	md5.Write(stdoutByte)
	stdout := string(md5.Sum(nil))
	if output == stdout {
		fmt.Println("AC")
	} else {
		fmt.Println(output)
		fmt.Println(stdout)
		fmt.Println("WRONG ANSWER")
	}
}
