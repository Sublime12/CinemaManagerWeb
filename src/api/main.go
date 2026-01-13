package main

import (
	"fmt"
	"os"
	"api/core_router"
)

const dns = "host=db user=cinema_manager password=cinema_manager dbname=cinema_manager port=5432 sslmode=disable"
const redisDns = "redis:6379"

func main() {

	router, err := core_router.Engine(dns, redisDns)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating router: ", err)
		return
	}

	router.Run(":8080")
}


