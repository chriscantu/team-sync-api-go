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

    r.GET("/questions", ctrl.GetQuestions)
    r.GET("/questions/:questionId", ctrl.GetQuestion)
    r.POST("/questions", ctrl.CreateQuestion)
    r.PUT("/questions/:questionId", ctrl.UpdateQuestion)
    r.DELETE("/questions/:questionId", ctrl.DeleteQuestion)

    r.Run(":9000")
}
