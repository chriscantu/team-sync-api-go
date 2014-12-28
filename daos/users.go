package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    mod "../models"
)

func GetUsers() (*mod.Users, error) {
    users := []mod.User{}
    usersMod := new (mod.Users)

    response, err := GetList("users")

    err = response.All(&users)
    count, err := GetCount("users")

    usersMod.Users = users
    usersMod.Total = count

    return usersMod, err
}

func GetUser(userId string) (mod.User, error) {
    user := mod.User{}

    response, err := GetRec("users", userId)
    response.Next(&user)

    return user, err
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
