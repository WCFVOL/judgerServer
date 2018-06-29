package judger

import "testing"

func TestHandler(t *testing.T) {
	str := `{"date":1530115200000,"code":"123","language":1,"id":4,"problemId":1,"userId":1}`
	Handler(str)
}
