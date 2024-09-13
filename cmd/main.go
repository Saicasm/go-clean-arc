package main

import (
	"log"

	"github.com/clean-arc/internal/constants"
	"github.com/clean-arc/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRoutes(r)
	port := constants.ServerPort
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
