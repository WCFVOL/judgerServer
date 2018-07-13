package problem

import "testing"

func TestAddInputOutput(t *testing.T) {
	str:="{\"problemId\":1, \"caseId\":1, \"input\":\"1\"}"
	AddInputOutput(str)
}
