package daos

import (
    "fmt"
    db "github.com/dancannon/gorethink"
    mod "../models"
)

func GetUsers() (*mod.Users, error) {
    var count float64
    session := GetSession()
    users := []mod.User{}
    usersMod := new (mod.Users)

    response, err := db.Table("users").Run(session)

    if err != nil {
        fmt.Println("e:", err)
    }
    err = response.All(&users)

    resp, err := db.Table("users").Count().Run(session)
    resp.Next(&count)

    usersMod.Users = users
    usersMod.Total = count

    return usersMod, err
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
