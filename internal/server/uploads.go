package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)


func uploadFile(c *gin.Context) {
    file, err := c.FormFile("file")
    if err != nil {
        // TODO: HANDLE ERROR!
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=server_error")
        return
    }

    extension := file.Filename[len(file.Filename)-4:len(file.Filename)]
    if extension != ".jpg" {
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=invalid_filetype")
        return
    }

    imageCounts, err := fileCount("./static/images/")
    if err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=server_error")
        return
    }

    fmt.Println(imageCounts)

    c.JSON(http.StatusNotImplemented, gin.H {
        "code": http.StatusNotImplemented,
        "message": "This endpoint is not functional yet",
    })
}


func fileCount(path string) (int, error){
    i := 0
    files, err := os.ReadDir(path)
    if err != nil {
        return 0, err
    }
    for _, file := range files {
        if !file.IsDir() { 
            i++
        }
    }
    return i, nil
}
