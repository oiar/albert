package albert

import (
	"testing"
)

func compare(res, response []bool) bool {
	for i := 0; i < len(res); i++ {
		if res[i] == response[i] {
			continue
		}
		return false
	}
	return true
}

func TestService(t *testing.T) {
	//var m sync.Mutex
	go func() {
		//m.Lock()
		StartServer()
		//m.Unlock()
	}()

	// request 1
	con1 := []string{"one", "two", "three"}
	response, err := PostWithJson(con1)
	if err != nil {
		t.Fatal(err)
	}

	res1 := []bool{false,false,false}
	if compare(res1, response) {
		t.Errorf("Response was incorrect, got: %t, want: %t.", response, res1)
	}

	// request 2
	con2 := []string{"one", "four", "five"}
	response, err = PostWithJson(con2)
	if err != nil {
		t.Fatal(err)
	}

	res2 := []bool{true,false,false}
	if compare(res2, response) {
		t.Errorf("Response was incorrect, got: %t, want: %t.", response, res2)
	}

	// request 3
	con3 := []string{"two", "six", "three", "seven"}
	response, err = PostWithJson(con3)
	if err != nil {
		t.Fatal(err)
	}

	res3 := []bool{true,false,true,false}
	if compare(res3, response) {
		t.Errorf("Response was incorrect, got: %t, want: %t.", response, res3)
	}
}
