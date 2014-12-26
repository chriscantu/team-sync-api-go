package daos

import (
    "fmt"
    db "github.com/dancannon/gorethink"
    mod "../models"
)

func GetUsers() (*[]mod.User, error) {
    session := GetSession()
    dbUsers := []mod.User{}

    response, err := db.Table("users").OrderBy(db.Asc("id")).Run(session)

    if err != nil {
        fmt.Println("e:", err)
    }

    err = response.All(&dbUsers)

    return &dbUsers, err
}

func GetUser(userId string) (mod.User, error) {
    session := GetSession()
    user := mod.User{}
    response, err := db.Table("users").Get(userId).Run(session)

    if err != nil {
        fmt.Println("e:", err)
    }

    response.Next(&user)

    return user, err
}

func CreateUser(user mod.User) (mod.User, error) {
    session := GetSession()

    response, err := db.Table("users").Insert(user).RunWrite(session)

    if err != nil {
        fmt.Println("e:", err)
    }

    user.Id = response.GeneratedKeys[0]

    return user, err
}

func UpdateUser(user mod.User) (mod.User, error) {
    session := GetSession()

    _, err := db.Table("users").Get(user.Id).Update(user).RunWrite(session)

    return user, err
}

func DeleteUser(userId string) (error) {
    session := GetSession()
    _, err := db.Table("users").Get(userId).Delete().RunWrite(session)

    return err
}
