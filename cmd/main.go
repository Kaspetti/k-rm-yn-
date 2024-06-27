package main

import (
	"log"
	"os"

	"github.com/Kaspetti/k-rm-yn-/internal/server"
)



func main() {
    ip, exists := os.LookupEnv("IP")
    if !exists {
        log.Printf("IP not in environment. Using '127.0.0.1'")
        ip = "127.0.0.1"
    }

    port, exists := os.LookupEnv("PORT")
    if !exists {
        log.Printf("PORT not in environment. Using '6969'")
        port = "6969"
    }

    if err := server.StartServer(ip, port); err != nil {
        log.Printf("Error while starting server: %v\n", err)
        return
    }
}
