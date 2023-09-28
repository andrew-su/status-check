package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	endpoint, ok := os.LookupEnv("STATUS_ENDPOINT")
	if !ok {
		panic(fmt.Errorf("STATUS_ENDPOINT environment variable not set"))
	}
	go func() {
		for {
			if err := checkStatus(endpoint); err != nil {
				panic(err)
			}

			time.Sleep(1 * time.Second)
		}
	}()

	http.ListenAndServe("0.0.0.0:8080", http.HandlerFunc(getRoot))
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello\n")
}

func checkStatus(ep string) error {
	resp, err := http.Get(ep)
	if err != nil {
		panic(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	if string(b) != "ok" {
		return fmt.Errorf("got '%s'", string(b))
	}
	return nil
}
