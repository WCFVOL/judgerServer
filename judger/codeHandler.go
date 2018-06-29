package judger

import (
	"encoding/json"
	"log"
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

}
