package judger

import (
	"testing"
	"awesomeProject/judgerServer/result"
	"encoding/json"
	"fmt"
)

func TestHandler(t *testing.T) {
	str := `{
    "cpu_time": 0,
    "real_time": 2,
    "memory": 3145728,
    "signal": 0,
    "exit_code": 0,
    "error": 0,
    "result": 0
}`
var res result.Result
json.Unmarshal([]byte(str),&res)
	fmt.Println(res)
}
