package controllers

import (
    "github.com/gin-gonic/gin"
    mod "../models"
)

func Hello (c *gin.Context) {
    c.JSON(200, mod.Hello{"world"})
}
