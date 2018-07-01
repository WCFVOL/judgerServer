package result

import (
	"net/http"
	"strings"
	"fmt"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Result struct {
	Result int
	Error int
	Exit_code int
	Signal int
	Memory int
	Read_time int
	Cpu_time int
}
type sendJson struct {
	Id int
	Length int
	Memory int
	Time int
	Result int
}
func Send(res Result,id, length int) {
	seJson := sendJson{id,length,res.Memory,res.Cpu_time,res.Result}
	sendStr,_ := json.Marshal(seJson)
	fmt.Println(string(sendStr))
	resp, err := http.Post("123.235.209.242:8080/admin/set_result",
		"application/json",
		strings.NewReader(string(sendStr)))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(body))
}