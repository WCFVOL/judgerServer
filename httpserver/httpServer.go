package httpserver

import (
	"awesomeProject/judgerServer/work"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

//Run run httpserver and bind ip:port
func Run(ip, port string) {
	r := gin.Default()
	r.POST("/task", getTask)
	r.Run(ip + ":" + port)
}

func getTask(c *gin.Context) {
	value := c.PostForm("task")
	var task work.Task
	err := json.Unmarshal([]byte(value), &task)
	if err != nil {
		log.Println(err.Error())
	}
	work.AddTask(task)
	//fmt.Printf(value)

}
