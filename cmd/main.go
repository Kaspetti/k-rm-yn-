package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Kaspetti/k-rm-yn-/internal/data"
	"github.com/gin-gonic/gin"
)



func main() {
    r := gin.Default() 
    r.SetTrustedProxies(nil)

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

    r.Static("/static", "./static")

    r.LoadHTMLGlob("./templates/*")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H {
            "title": "Karmøy Stickers",
        })
    })


    r.GET("/api/data", func(c *gin.Context) {
        karmoyStickers, err := data.GetData("data.csv")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "message": "An error occured on our end when fetching the data",
            })
        }

        c.JSON(http.StatusOK, karmoyStickers)
    })

    r.Run(fmt.Sprintf("%s:%s", ip, port))
}
