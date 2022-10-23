package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func server(port string) {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "root\n")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hey\n")
	})

	err := http.ListenAndServe(":" + port, nil)

  if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}