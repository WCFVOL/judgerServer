package result

import "testing"

func TestSend(t *testing.T) {
	res := Result{0,0,0,0,10000,3,2}
	Send(res,6,1000)
}
