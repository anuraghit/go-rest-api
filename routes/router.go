package routes

import (
	"gihub.com/anuraghit/go-assin/controller"
	"github.com/julienschmidt/httprouter"
)

func SetUp(r *httprouter.Router) {
	r.GET("/", controller.Welcome)
	r.GET("/api/v1/user/:id", controller.GetUser)
	r.GET("/api/v1/user", controller.GetUsers)
	r.POST("/api/v1/user", controller.CreateUser)
	r.DELETE("/api/v1/user/:id", controller.DeleteUser)
	r.PATCH("/api/v1/user", controller.UpdateUser)
}
