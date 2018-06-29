package judger

import (
	"encoding/json"
	"log"
	"strconv"
)

type Submission struct {
	Id int
	Code string
	UserId int
	Date int64
	ProblemId int
	Language int
}

func Handler(str string) {
	var submission Submission;
	err := json.Unmarshal([]byte(str),&submission)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(str)
	log.Print(submission)
	WriteCode(submission) // 根据submissionId将代码写入相应文件
	compilerResult:=Compiler(submission)
	if compilerResult == "error" {
		//TODO compiler error
		return
	}
	judgeResult := Judger(1000, 2000, 128*1024*1024, 200, 10000, 32*1024*1024, 0, 0, 0,
		"/root/user_result/"+strconv.Itoa(submission.Id), strconv.Itoa(submission.ProblemId)+".in", "/root/user_result/"+strconv.Itoa(submission.Id)+".out", "/root/user_result/echo.out", "judger.log", "c_cpp",
		[]string{""}, []string{"foo=bar"})
	log.Println(judgeResult)
}
