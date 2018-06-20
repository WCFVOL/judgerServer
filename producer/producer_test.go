package producer

import (
	"testing"
)

func TestSend(t *testing.T) {
	result:=Send("fuck u")
	t.Log("%s",result)
}