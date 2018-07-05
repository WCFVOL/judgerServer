package result

import "testing"

func TestSend(t *testing.T) {
	res := Result{0,0,0,0,10000,3,2}
	Send(res,13,1000)
}
