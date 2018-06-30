package judger

import "testing"

func TestCompiler(t *testing.T) {
	str := Submission{Date: 1530115200000, Code: "123", Language: 1, Id: 4, ProblemId: 1, UserId: 1}
	println(Compiler(str))
}
