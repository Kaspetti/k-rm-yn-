package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kaspetti/k-rm-yn-/internal/data"
	"github.com/gin-gonic/gin"
)


func StartServer(ip, port string) error {
    r := gin.Default() 
    r.SetTrustedProxies(nil)

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
            log.Printf("Error: %v\n", err)
            c.JSON(http.StatusInternalServerError, gin.H {
                "message": "An error occured on our end when fetching the data",
            })
        }

        c.JSON(http.StatusOK, karmoyStickers)
    })

    log.Printf("Start listening on: %s:%s\n", ip, port)
    return r.Run(fmt.Sprintf("%s:%s", ip, port))
}
