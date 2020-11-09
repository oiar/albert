package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Message struct {
	Content []string `json:"content"`
}

type Store struct {
	Items []string
	mux sync.Mutex
}

func (sto *Store) receiveMessage(w http.ResponseWriter, r *http.Request) {
	var mes Message
	err := json.NewDecoder(r.Body).Decode(&mes)
	if err != nil {
		_ = r.Body.Close()
		fmt.Println("Errors in NewDecoder", err)
	}

	var res = make([]bool, len(mes.Content))
	sto.mux.Lock()
	for i := 0; i < len(mes.Content); i++ {
		if isContain(sto.Items, mes.Content[i]) {
			res[i] = true
			continue
		}
		sto.Items = append(sto.Items, mes.Content[i])
	}
	sto.mux.Unlock()

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Println("Errors in NewEncoder", err)
	}
}

func isContain(items []string, item string) bool {
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			return true
		}
	}
	return false
}

func StartServer() {
	s := &http.Server{
		Addr: ":80",
		Handler: nil,
	}

	sto := &Store{Items: []string{}}
	http.HandleFunc("/send", sto.receiveMessage)

	err := s.ListenAndServeTLS("./cert/localhost.crt", "./cert/localhost.key")
	if err != nil {
		fmt.Println("ListenAndServeTLS", err)
	}
}