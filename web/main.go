package main

import (
	"log"
	_ "silent-cxl.top/app/bootstrap"
)

func main() {
	if err := routers.Router(); err != nil {
		log.Fatal(err)
	}
}
