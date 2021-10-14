package main

import (
	"fmt"
	"net/http"

	"./assets"
	"golang.org/x/text/message"
)

func handlerHallo(w http.ResponseWriter, r *http.Request) {
	var message = "Wellcomww"
	w.Write([]byte(message))
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

func main() {
	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)
	// http.HandleFunc("/sapa", handlerHallo)

	//mencoba menggunakan routing static
	http.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("assets"))))

	var address = "localhost:9000"
	fmt.Print("ur server", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	message := 

}
