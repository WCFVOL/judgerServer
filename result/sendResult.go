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
	Real_time int
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
	req, err := http.NewRequest("POST","http://localhost:8080/admin/set_result",
		strings.NewReader(string(sendStr)))
	req.Header.Set("Content-Type", "application/json; encoding=utf-8")
	req.Header.Set("token", "isAdmin")
	client := &http.Client{}
	resp, err := client.Do(req)
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