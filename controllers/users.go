package controllers

import (
    "github.com/gin-gonic/gin"
    dao "../daos"
    mod "../models"
)

func GetUsers (c *gin.Context) {
    users, _ := dao.GetUsers()

    c.JSON(200, users)
}

func GetUser (c *gin.Context) {
    userId := c.Params.ByName("userId")
    user := dao.GetUser(userId)

    if user.Id == "" {
        c.JSON(404, mod.Get404("User", userId))
    } else {
        c.JSON(200, user)
    }
}

func CreateUser (c *gin.Context) {
    var json mod.User
    c.Bind(&json)

    user, err := dao.CreateUser(json)

    if err == nil {
        c.JSON(201, user)
    } else {
        c.JSON(500, mod.Message{"Error creating User"})
    }
}

func UpdateUser (c *gin.Context) {
    var json mod.User
    json.Id = c.Params.ByName("userId")

    c.Bind(&json)

    user, err := dao.UpdateUser(json)

    if err == nil {
        c.JSON(200, user)
    } else {
        c.JSON(500, mod.Message{"Error updating User"})
    }
}

func DeleteUser (c *gin.Context) {
    userId := c.Params.ByName("userId")

    err := dao.DeleteUser(userId)

    if err == nil {
        c.JSON(200, mod.Message{"Deleted User"})
    } else {
        c.JSON(500, mod.Message{"Error Deleting User"})
    }
}
