package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CMessage struct {
	Content []string `json:"content"`
}

func PostWithJson(con []string) ([]bool, error) {
	mes := CMessage{con}
	b, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("Errors in MarshalMes", err)
		return nil, err
	}

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	resp, err := c.Post("https://127.0.0.1:80/send","application/json", bytes.NewBuffer(b))
	if err != nil {
		fmt.Println("Errors in Post", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Errors in ReadAll", err)
		return nil, err
	}

	//fmt.Printf("Resonse: %s\n", string(body))
	cb := convertByte(body)
	return cb, nil
}

func convertByte(data []byte) []bool {
	res := make([]bool, len(data) * 8)
	for i := range res {
		res[i] = data[i/8]&(0x80>>byte(i&0x7)) != 0
	}
	return res
}
