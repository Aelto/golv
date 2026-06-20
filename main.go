package main

import (
	"fmt"
	"golv/pages/home"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")

	home.ViewHome.Register()

	log.Fatal(http.ListenAndServe(":3000", nil))
}
