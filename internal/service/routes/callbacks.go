package routes

import "github.com/gin-gonic/gin"

func Landing(c *gin.Context) {
	c.HTML(200, "landing.html", gin.H{})
}
