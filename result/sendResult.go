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
	resp, err := http.Post("http://106.15.183.211:8080/admin/set_result",
		"application/json",
		strings.NewReader(string(sendStr)))
	if err != nil {
		log.Println(err.Error())
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(string(body))
}