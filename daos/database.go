package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    )

var session *db.Session

func InitDb() {
    sessionTmp, err := db.Connect(db.ConnectOpts{
        Address: "localhost:28015",
        Database: "teamsync",
    })

    if err != nil {
        log.Fatal(err)
        return
    }
    session = sessionTmp
}

func GetSession() *db.Session {
    return session
}

func GetRec(table string, id string) (*db.Cursor, error) {

    response, err := db.Table(table).Get(id).Run(session)

    if err != nil {
        log.Panic(err)
    }

    return response, err
}

func GetList(table string) (*db.Cursor, error) {
    response, err := db.Table(table).Run(session)

    if err != nil {
        log.Panic(err)
    }

    return response, err
}


func GetCount(table string) (float64, error) {
    var count float64
    response, err := db.Table("users").Count().Run(session)

    if err != nil {
        log.Panic(err)
    }

    response.Next(&count)

    return count, err
}

func DeleteRec(table string, id string) (error) {
    _, err := db.Table(table).Get(id).Delete().RunWrite(session)

    if err != nil {
        log.Panic(err)
    }

    return err
}
