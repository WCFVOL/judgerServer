package problem

import (
	"encoding/json"
	"os"
	"strconv"
)

type AddFile struct {
	ProblemId int
	CaseId    int
	Input     string
	Output    string
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func AddInputOutput(str string) {
	//TODO
	var addFile AddFile
	json.Unmarshal([]byte(str), &addFile)

	if len(addFile.Output) == 0 {
		if ok, _ := PathExists("/root/problem_in/" + strconv.Itoa(addFile.ProblemId)); !ok {
			os.Mkdir("/root/problem_in/"+strconv.Itoa(addFile.ProblemId), os.ModePerm)
		}
		file, _ := os.Create("/root/problem_in/" + strconv.Itoa(addFile.ProblemId) + "/" + strconv.Itoa(addFile.CaseId) + ".in")
		file.Write([]byte(addFile.Input))
	} else {
		if ok, _ := PathExists("/root/std_result/" + strconv.Itoa(addFile.ProblemId)); !ok {
			os.Mkdir("/root/std_result/"+strconv.Itoa(addFile.ProblemId), os.ModePerm)
		}
		file, _ := os.Create("/root/std_result/" + strconv.Itoa(addFile.ProblemId) + "/" + strconv.Itoa(addFile.CaseId) + ".out")
		file.Write([]byte(addFile.Output))
	}
}
