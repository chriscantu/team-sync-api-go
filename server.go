package main

import (
    "github.com/gin-gonic/gin"
    ctrl "./controllers"
    daos "./daos"
)

func main() {
    r := gin.Default()

    daos.InitDb()

    r.GET("/", ctrl.Hello)
    r.GET("/users", ctrl.GetUsers)
    r.GET("/users/:userId", ctrl.GetUser)
    r.POST("/users", ctrl.CreateUser)
    r.PUT("/users/:userId", ctrl.UpdateUser)
    r.DELETE("/users/:userId", ctrl.DeleteUser)

    r.Run(":9000")
}
