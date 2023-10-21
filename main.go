package main

import (
	"fmt"
	"net/http"
)

func catchReq(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("0xNull"))
}

func main() {
	http.HandleFunc("/", catchReq)
	port := ":8080"
	fmt.Printf("01101001 01101110 01101001 01100011 01101001 01100001 01101110 01100100 01101111 00100000 01110011 01100101 01110010 01110110 01101001 01100100 01101111 01110010 %s\n", port)
	http.ListenAndServe(port, nil)
}
