package server

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


func uploadFile(c *gin.Context) {
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    c.JSON(http.StatusNotImplemented, gin.H {
        "code": http.StatusNotImplemented,
        "message": "This endpoint is not functional yet",
    })
}
