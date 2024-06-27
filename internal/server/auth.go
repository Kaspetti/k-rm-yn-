package server

import (
	"crypto/subtle"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func login(c *gin.Context) {
    pass := c.PostForm("password")
    uname := c.PostForm("username")

    adminUname := os.Getenv("ADMIN_UNAME")
    adminPass := os.Getenv("ADMIN_PWD_HASH")

    if subtle.ConstantTimeCompare([]byte(uname), []byte(adminUname)) != 1 {
        c.JSON(http.StatusUnauthorized, gin.H {
            "code": http.StatusUnauthorized,
            "message": "invalid credentials",
        })
        return
    }

    if bcrypt.CompareHashAndPassword([]byte(adminPass), []byte(pass)) != nil || uname != adminUname {
        c.JSON(http.StatusUnauthorized, gin.H {
            "code": http.StatusUnauthorized,
            "message": "invalid credentials",
        })
        return
    }

    sessionToken := uuid.New().String()
    c.SetCookie("session", sessionToken, 3600, "/admin", "karmoy.kaspeti.com", true, true)
    c.Redirect(http.StatusFound, "/admin")
}
