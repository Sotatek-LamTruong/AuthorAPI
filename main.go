package main

import (
	"book-author/pkg/app"
	"fmt"
	"log"
)

func main() {

	r := app.SetupApp()

	const port = "3000"

	fmt.Println("App running on port " + port)
	log.Fatalln(r.Run(fmt.Sprintf(":%s", port)))
}
