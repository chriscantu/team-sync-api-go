package controllers

import (
    "github.com/gin-gonic/gin"
    dao "../daos"
    mod "../models"
    "strconv"
)

func GetUsers (c *gin.Context) {
    users, err := dao.GetUsers()

    if err != nil {
        c.JSON(500, mod.SystemError{"System Error"})
    } else {
        c.JSON(200, users)
    }
}

func GetUser (c *gin.Context) {
    userId := c.Params.ByName("userId")
    user, err := dao.GetUser(userId)

    if err != nil {
        c.JSON(500, mod.SystemError{"Unable to retrieve User"})
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
        c.JSON(500, mod.SystemError{"Error creating User"})
    }
}

func UpdateUser (c *gin.Context) {
    var json mod.User
    c.Bind(&json)

    user, err := dao.UpdateUser(json)

    if err == nil {
        c.JSON(200, user)
    } else {
        c.JSON(500, mod.SystemError{"Error updating User"})
    }
}

func DeleteUser (c *gin.Context) {
    userId := c.Params.ByName("userId")
    id,err := strconv.Atoi(userId)

    dao.DeleteUser(id)

    if err == nil {
        c.JSON(200, mod.SystemError{"Deleted User"})
    } else {
        c.JSON(500, mod.SystemError{"Error Deleting USer"})
    }
}
