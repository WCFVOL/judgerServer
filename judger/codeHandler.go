package judger

import (
	"awesomeProject/judgerServer/result"
	"crypto/md5"
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

type Submission struct {
	Id        int
	Code      string
	Username  string
	Date      int64
	ProblemId int
	Language  int
	TestCase  int
	TimeLimit int
	MemLimit  int
}

func Handler(str string) {
	var submission Submission
	err := json.Unmarshal([]byte(str), &submission)
	if err != nil {
		log.Println(err.Error())
	}
	var res result.Result
	res.Result = 8
	result.Send(res, submission.Id, 0)
	log.Println(str)
	log.Print(submission)
	WriteCode(submission) // 根据submissionId将代码写入相应文件
	compilerResult := Compiler(submission)
	if compilerResult == "error" {
		res := result.Result{6, 0, 0, 0, 0, 0, 0}
		result.Send(res, submission.Id, 0)
		return
	}
	res.Result = 9
	result.Send(res, submission.Id, 0)
	for i := 0; i < submission.TestCase; i++ {
		judgeResult := Judger(submission.TimeLimit, submission.TimeLimit*2, submission.MemLimit*1024*1024, 200, 10000, 32*1024*1024, 0, 0, 0,
			"/root/user_code/"+strconv.Itoa(submission.Id), "/root/problem_in/"+strconv.Itoa(submission.ProblemId)+"/"+strconv.Itoa(i)+".in", "/root/user_result/"+strconv.Itoa(submission.Id)+".out", "/root/user_result/"+strconv.Itoa(submission.Id)+".error", "judger.log", "c_cpp",
			[]string{""}, []string{"foo=bar"})
		log.Println(string(judgeResult))
		err = json.Unmarshal(judgeResult, &res)
		if err != nil {
			log.Println(err.Error())
		}
		if res.Result != 0 {
			result.Send(res, submission.Id, 0)
			return
		}
		outputByte, _ := ioutil.ReadFile("/root/user_result/" + strconv.Itoa(submission.Id) + ".out")
		stdoutByte, _ := ioutil.ReadFile("/root/std_result/" + strconv.Itoa(submission.ProblemId) + "/" + strconv.Itoa(i) + ".out")

		md5 := md5.New()
		md5.Write(outputByte)
		output := string(md5.Sum(nil))
		md5.Reset()
		md5.Write(stdoutByte)
		stdout := string(md5.Sum(nil))
		if output == stdout {
			res.Result = 0
		} else {
			res.Result = -1
			break
		}
	}

	result.Send(res, submission.Id, len(submission.Code))
}
