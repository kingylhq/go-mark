package example

import (
	"fmt"
	"net/http"
	"time"
)

func hello2(w http.ResponseWriter, req *http.Request) {

	ctx := req.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():

		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func Context() {

	// http://localhost:8090/hello2
	http.HandleFunc("/hello2", hello)
	http.ListenAndServe(":8091", nil)
}
