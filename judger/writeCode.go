package judger

import (
	"os"
	"strconv"
	"log"
)

func WriteCode(submission Submission) {
	file, error := os.OpenFile("../user_code/"+strconv.Itoa(submission.Id)+".cpp", os.O_RDWR|os.O_CREATE, 0766);
	if error!=nil {
		log.Println(error.Error())
	}
	defer file.Close()
	size,error2 := file.Write([]byte(submission.Code))
	if error2 != nil {
		log.Println(error2.Error())
	}
	log.Printf("size: %d",size)
}
