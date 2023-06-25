package main

import (
	"log"
	_ "silent-cxl.top/app/bootstrap"
	"silent-cxl.top/router"
)

func main() {
	if err := router.Router(); err != nil {
		log.Fatal(err)
	}
}
