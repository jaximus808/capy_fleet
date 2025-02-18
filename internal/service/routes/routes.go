package routes

import "github.com/gin-gonic/gin"

func CreateRoutes(r *gin.Engine) {
	GET_Routes := map[string]func(*gin.Context){
		"/":   Landing,
		"/ws": create_websocket,
	}

	for route, callback := range GET_Routes {
		r.GET(route, callback)
	}
}
