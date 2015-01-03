package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    mod "../models"
)

func GetTeams() (*mod.Teams, error) {
    teams := []mod.Team{}
    teamsMod := new (mod.Teams)

    response, err := GetList("teams")

    err = response.All(&teams)
    count, err := GetCount("teams")

    teamsMod.Teams = teams
    teamsMod.Total = count

    return teamsMod, err
}

func GetTeam(teamId string) (mod.Team) {
    team := mod.Team{}

    response := GetRec("teams", teamId)
    response.Next(&team)

    return team
}

func CreateTeam(team mod.Team) (mod.Team, error) {
    session := GetSession()

    response, err := db.Table("teams").Insert(team).RunWrite(session)

    if err != nil {
        log.Panic(err)
    }

    team.Id = response.GeneratedKeys[0]

    return team, err
}

func UpdateTeam(team mod.Team) (mod.Team, error) {
    session := GetSession()

    _, err := db.Table("teams").Get(team.Id).Update(team).RunWrite(session)

    return team, err
}

func DeleteTeam(teamId string) (error) {
    return DeleteRec("teams", teamId)
}
