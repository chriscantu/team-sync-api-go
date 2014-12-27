package controllers

import (
    "github.com/gin-gonic/gin"
    dao "../daos"
    mod "../models"
)

func GetQuestions (c *gin.Context) {
    questions, err := dao.GetQuestions()

    if err != nil {
        c.JSON(500, mod.Message{"System Error"})
    } else {
        c.JSON(200, questions)
    }
}

func GetQuestion (c *gin.Context) {
    questionId := c.Params.ByName("questionId")
    question, err := dao.GetQuestion(questionId)

    if err != nil {
        c.JSON(500, mod.Message{"Unable to retrieve Question"})
    } else if question.Id == "" {
        c.JSON(404, mod.Message{"Unable to find Question by Id: " + questionId})
    } else {
        c.JSON(200, question)
    }
}

func CreateQuestion (c *gin.Context) {
    var json mod.Question
    c.Bind(&json)

    question, err := dao.CreateQuestion(json)

    if err == nil {
        c.JSON(201, question)
    } else {
        c.JSON(500, mod.Message{"Error creating Question"})
    }
}

func UpdateQuestion (c *gin.Context) {
    var json mod.Question
    json.Id = c.Params.ByName("questionId")

    c.Bind(&json)

    question, err := dao.UpdateQuestion(json)

    if err == nil {
        c.JSON(200, question)
    } else {
        c.JSON(500, mod.Message{"Error updating Question"})
    }
}

func DeleteQuestion (c *gin.Context) {
    questionId := c.Params.ByName("questionId")

    err := dao.DeleteQuestion(questionId)

    if err == nil {
        c.JSON(200, mod.Message{"Deleted Question"})
    } else {
        c.JSON(500, mod.Message{"Error Deleting USer"})
    }
}
