package main

import (
	"github.com/gin-gonic/gin"
	"github.com/marcobinda/MedSupply/db"
)

func main() {
	db.InitDB()

	server := gin.Default()

	server.Run()
}
