package judger

import (
	"os/exec"
	"fmt"
	"strconv"
	"bytes"
)

var (
	RESULT_SUCCESS                  = 0
	RESULT_WRONG_ANSWER             = -1
	RESULT_CPU_TIME_LIMIT_EXCEEDED  = 1
	RESULT_REAL_TIME_LIMIT_EXCEEDED = 2
	RESULT_MEMORY_LIMIT_EXCEEDED    = 3
	RESULT_RUNTIME_ERROR            = 4
	RESULT_SYSTEM_ERROR             = 5

	ERROR_INVALID_CONFIG      = -1
	ERROR_FORK_FAILED         = -2
	ERROR_PTHREAD_FAILED      = -3
	ERROR_WAIT_FAILED         = -4
	ERROR_ROOT_REQUIRED       = -5
	ERROR_LOAD_SECCOMP_FAILED = -6
	ERROR_SETRLIMIT_FAILED    = -7
	ERROR_DUP2_FAILED         = -8
	ERROR_SETUID_FAILED       = -9
	ERROR_EXECVE_FAILED       = -10
	ERROR_SPJ_ERROR           = -11
)

func Judger(maxCpuTime, maxRealTime, maxMemory, maxProcessNumber, maxOutputSize, maxStack, gid, uid, memoryLimitCheckOnly int,
	exePath, inputPath, outputPath, errorPath, logPath, seccompRuleName string,
	args, env []string) string {
	tmp:="--args="
	for _,v:=range args {
		tmp+=v+" "
	}
	strenv:="--env="
	for _,v:=range args {
		strenv+=v+" "
	}
	proc := exec.Command("/root/Judger-newnew/Judger-newnew/output/libjudger.so","--max_cpu_time=" + strconv.Itoa(maxCpuTime),
		"--max_real_time=" + strconv.Itoa(maxRealTime),
		"--max_memory=" + strconv.Itoa(maxMemory) ,
		"--max_stack=" + strconv.Itoa(maxStack) ,
		"--max_output_size=" + strconv.Itoa(maxOutputSize) ,
		"--max_process_number=" + strconv.Itoa(maxProcessNumber) ,
		"--exe_path=" + string(exePath) ,
		"--input_path=" + string(inputPath) ,
		"--output_path=" + string(outputPath) ,
		"--error_path=" + string(errorPath) ,
		"--log_path=" + string(logPath) ,
		"--seccomp_rule_name=" + string(seccompRuleName) ,
		"--uid=" + strconv.Itoa(uid) ,
		"--gid=" + strconv.Itoa(gid) ,
		"--memory_limit_check_only=" + strconv.Itoa(memoryLimitCheckOnly),tmp,strenv,
	)
	var out bytes.Buffer
	proc.Stdout = &out
	err:=proc.Run()
	if err!=nil {
		fmt.Println(err)
	}
	return out.String()
}
