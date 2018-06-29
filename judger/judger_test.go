package judger

import (
	"fmt"
	"testing"
)

//maxCpuTime, maxRealTime, maxMemory, maxProcessNumber, maxOutputSize, maxStack, gid, uid, memoryLimitCheckOnly int,
//	exePath, inputPath, outputPath, errorPath, logPath, seccompRuleName string,
//	args, env []string
func TestJudger(t *testing.T) {
	result := Judger(1000, 2000, 128*1024*1024, 200, 10000, 32*1024*1024, 0, 0, 0,
		"../main", "in.txt", "out.txt", "echo.out", "judger.log", "c_cpp",
		[]string{""}, []string{"foo=bar"})
	fmt.Println(result)
	t.Log(result)
}
