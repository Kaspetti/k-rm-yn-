package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kaspetti/k-rm-yn-/internal/data"
	"github.com/gin-gonic/gin"
)


func uploadFile(c *gin.Context) {
    isValidSession, errorMessage := validSession(c)
    if !isValidSession {
        c.Redirect(http.StatusSeeOther, fmt.Sprintf("/login?auth_status=%s", errorMessage))
        return
    }

    file, err := c.FormFile("file")
    if err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=file_error")
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
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=count_error")
        return
    }

    lat := c.PostForm("latitude")
    lon := c.PostForm("longitude")
    desc := c.PostForm("description")

    d, err := data.GetData("data.json")
    if err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=data_read_error")
        return
    }

    latFloat, err := strconv.ParseFloat(lat, 64)
    if err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=invalid_lat")
        return
    }

    lonFloat, err := strconv.ParseFloat(lon, 64)
    if err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=invalid_lon")
        return
    }

    d = append(d, data.KarmoySticker {
        ID: imageCounts,
        Latitude: latFloat,
        Longitude: lonFloat,
        Description: desc,
    })

    if err := c.SaveUploadedFile(file, fmt.Sprintf("./static/images/%d.jpg", imageCounts)); err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=file_upload_error")
        return
    } 

    if err := data.SaveData("data.json", d); err != nil {
        log.Println(err)
        c.Redirect(http.StatusSeeOther, "/admin?upload_status=data_write_error")
        return
    }

    c.Redirect(http.StatusSeeOther, "/admin?upload_status=success")
}


func fileCount(path string) (uint16, error){
    var i uint16 = 0
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
