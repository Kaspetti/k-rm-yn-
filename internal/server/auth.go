package server

import (
	"crypto/subtle"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


const (
    tokenMaxAge = 1 * time.Hour
)


var (
    sessionToken    string
    tokenExpiration time.Time
    tokenMutex      sync.RWMutex
)


func login(c *gin.Context) {
    pass := c.PostForm("password")
    uname := c.PostForm("username")

    adminUname := os.Getenv("ADMIN_UNAME")
    adminPass := os.Getenv("ADMIN_PWD_HASH")

    if subtle.ConstantTimeCompare([]byte(uname), []byte(adminUname)) != 1 {
        c.Redirect(http.StatusSeeOther, "/login?auth_status=invalid_credentials")
        return
    }

    if bcrypt.CompareHashAndPassword([]byte(adminPass), []byte(pass)) != nil || uname != adminUname {
        c.Redirect(http.StatusSeeOther, "/login?auth_status=invalid_credentials")
        return
    }

    tokenMutex.Lock()
    sessionToken = uuid.New().String()
    tokenExpiration = time.Now().Add(tokenMaxAge)
    tokenMutex.Unlock()

    c.SetCookie("session", sessionToken, 3600, "/admin", "karmoy.kaspeti.com", true, true)
    c.Redirect(http.StatusFound, "/admin")
}
