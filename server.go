package server

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Web!\n"))
}

func main() {
	fmt.Printf("Starting server at port 2020\n")
	if err := http.ListenAndServe(":2020", nil); err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)

}
