package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iamme24cl/clickipo_server_golang/db"
)



func main() {
    // connect to DB first
    env := new(Env)
    var err error
    env.DB, err = ConnectDB()
    if err != nil {
        log.Fatalf("failed to start the server: %v", err)
    }

    router := gin.Default()

    router.Run("localhost:8080")
}

