package daos

import (
    mod "../models"
    )

func GetQuestions() (*mod.Questions, error) {
    questions := []mod.Question{}
    quesMod := new(mod.Questions)

    response, err := GetList("questions")
    err = response.All(&questions)

    count, err :=GetCount("questions")
    quesMod.Questions = questions
    quesMod.Total = count

    return quesMod, err
}

func GetQuestion(questionId string) (mod.Question) {
    question := mod.Question{}

    response := GetRec("questions", questionId)
    response.Next(&question)

    return question
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
