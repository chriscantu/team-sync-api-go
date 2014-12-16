package daos

import (
    "encoding/json"
    "io/ioutil"
    "path/filepath"
    "fmt"
    mod "../models"
)

func GetUsers() (*mod.Users, error) {
    filePath, _ :=  filepath.Abs("json/users.json")
    file, e := ioutil.ReadFile(filePath)

    if e != nil {
        fmt.Println("e:", e)
    }

    var users mod.Users
    response := json.Unmarshal(file, &users)

    return &users, response
}
