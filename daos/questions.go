package daos

import (
    mod "../models"
    )

func GetQuestions() ([]mod.Question, error) {
    questions := []mod.Question{}

    response, err := GetList("questions")
    err = response.All(&questions)

    return questions, err
}

func GetQuestion(questionId string) (mod.Question, error) {
    question := mod.Question{}

    response, err := GetRec("users", questionId)
    response.Next(&question)

    return question, err
}

func CreateQuestion(question mod.Question) (mod.Question, error) {
    response, err := CreateRec("questions", question)
    question.Id = response.GeneratedKeys[0]

    return question, err
}

func UpdateQuestion(question mod.Question) (mod.Question, error) {
    err := UpdateRec("questions", question.Id, question)

    return question, err
}

func DeleteQuestion(questionId string) (error) {
    return DeleteRec("questions", questionId)
}
