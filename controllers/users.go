package controllers

import (
    "github.com/gin-gonic/gin"
    dao "../daos"
    mod "../models"
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
