package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    mod "../models"
)

func GetUser(userId string) (mod.User) {
    user := mod.User{}

    response := GetRec("users", userId)
    response.Next(&user)

    return user
}

func CreateUser(user mod.User) (mod.User, error) {
    session := GetSession()

    response, err := db.Table("users").Insert(user).RunWrite(session)

    if err != nil {
        log.Panic(err)
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
    return DeleteRec("users", userId)
}
