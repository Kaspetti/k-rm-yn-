// server is responsible for all server related functionality on karmoy stickers
// creates endpoints for necessary files and data
package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
    r.GET("/login", func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html", gin.H {
            "title": "Karmøy Login",
        })
    })
    r.GET("/admin", func(c *gin.Context) {
        mode := os.Getenv("GIN_MODE")
        if !validSession(c) && mode == "release" {
            c.Redirect(http.StatusSeeOther, "/login?auth_status=no_session")
            return
        }

        c.HTML(http.StatusOK, "admin.html", gin.H {
            "title": "Karmøy Admin",
        })
    })
    r.GET("/stats", func(c *gin.Context) {
        mode := os.Getenv("GIN_MODE")
        if !validSession(c) && mode == "release" {
            c.Redirect(http.StatusSeeOther, "/login?auth_status=no_session")
            return
        }

        c.JSON(http.StatusOK, gin.H {
            "stats": "not found",
        })
    })

    api := r.Group("/api")
    {
        // Get the data from the csv file using the data package
        // and send it to the user
        api.GET("/data", func(c *gin.Context) {
            karmoyStickers, err := data.GetData("data.json")
            if err != nil {
                log.Printf("Error: %v\n", err)
                c.JSON(http.StatusInternalServerError, gin.H {
                    "message": "An error occured on our end when fetching the data",
                })
            }

            c.JSON(http.StatusOK, karmoyStickers)
        })

        api.POST("/upload", uploadFile)

        api.POST("/login", login)
    }

    log.Printf("Start listening on: %s:%s\n", ip, port)
    return r.Run(fmt.Sprintf("%s:%s", ip, port))
}


func validSession(c *gin.Context) bool {
    sessionCookie, err := c.Cookie("session")
    // Redirect to login page if no session cookie is found
    if err != nil {
        return false
    }

    tokenMutex.RLock()
    validToken := sessionToken
    expirationTime := tokenExpiration
    tokenMutex.RUnlock()

    if sessionCookie != validToken {
        return false
    }

    if time.Now().After(expirationTime) {
        return false
    }

    return true
}
