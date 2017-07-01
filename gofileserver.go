package main

import (
	"log"
	"net/http"
	// "os"
)

func main() {
	//os.Mkdir("file", 0777)
	http.Handle("/anchor/", http.StripPrefix("/anchor/", http.FileServer(http.Dir("/root/test/dailydeeds/operation/Anchor"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
