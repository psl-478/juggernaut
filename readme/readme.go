package read

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func Documentation(c *gin.Context) {
	c.FileAttachment("./README.md", "JUGGERNAUT_Documentation.md")
}

func Contributors(c *gin.Context) {
	c.HTML(http.StatusOK, "contributors.html", gin.H{})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func Rota(c *gin.Context) {
	c.HTML(http.StatusOK, "rota.html", gin.H{})
}

func Game(c *gin.Context) {
	c.HTML(http.StatusOK, "game.html", gin.H{})
}
