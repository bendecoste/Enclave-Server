package main

import (
	"io"
	"log"
	"net/http"

	"github.com/mdlayher/vsock"
)

func main() {
	handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, Enclave!\n")
	}

	listener, err := vsock.Listen(2345, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.Serve(listener, http.HandlerFunc(handler)))
}
