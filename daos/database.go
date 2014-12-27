package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    )

var sessions []*db.Session

func InitDb() {
    session, err := db.Connect(db.ConnectOpts{
        Address: "localhost:28015",
        Database: "teamsync",
    })

    if err != nil {
        log.Fatal(err)
        return
    }

    sessions = append(sessions, session)
}

func GetSession() *db.Session {
    return sessions[0]
}
