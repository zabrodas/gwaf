package main

import (
	"fmt"
	"net/http"
)

func rqHello(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello!!!")
}

func main() {
	http.HandleFunc("/hello", rqHello)
    http.ListenAndServe("0.0.0.0:80",nil)
}
