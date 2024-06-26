package routes

import "github.com/gin-gonic/gin"

func InitRoute(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId")
	r.GET("/getUserByEmail/:userEmail")
	r.POST("createUser")
	r.PUT("/updateUser/:userId")
	r.DELETE("/deleteUser/:userId")
}
