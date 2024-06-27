<<<<<<< HEAD
package routes

import (
	"crud-go/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.updateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
=======
package routes

import (
	"crud-go/src/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.RouterGroup) {

	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.updateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
>>>>>>> d1ab3d39f55c8a54c6a868e23831d3dc0e22b327
