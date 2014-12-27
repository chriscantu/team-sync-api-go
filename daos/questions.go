package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    mod "../models"
    )

func GetQuestions() ([]mod.Question, error) {
    session := GetSession()
    questions := []mod.Question{}

    response, err := db.Table("questions").Run(session)

    if err != nil {
        log.Panic(err)
    }

    err = response.All(&questions)

    return questions, err
}

func GetQuestion(questionId string) (mod.Question, error) {
    session := GetSession()
    question := mod.Question{}
    response, err := db.Table("questions").Get(questionId).Run(session)

    if err != nil {
        log.Panic(err)
    }

    response.Next(&question)

    return question, err
}

func CreateQuestion(question mod.Question) (mod.Question, error) {
    session := GetSession()

    response, err := db.Table("questions").Insert(question).RunWrite(session)

    if err != nil {
        log.Panic(err)
    }

    question.Id = response.GeneratedKeys[0]

    return question, err
}

func UpdateQuestion(question mod.Question) (mod.Question, error) {
    session := GetSession()

    _, err := db.Table("questions").Get(question.Id).Update(question).RunWrite(session)

    return question, err
}

func DeleteQuestion(questionId string) (error) {
    session := GetSession()
    _, err := db.Table("questions").Get(questionId).Delete().RunWrite(session)

    return err
}
