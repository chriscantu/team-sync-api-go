package controllers

import (
    "github.com/gin-gonic/gin"
    dao "../daos"
    mod "../models"
)

func GetTeams (c *gin.Context) {
    teams, err := dao.GetTeams()

    if err != nil {
        c.JSON(500, mod.Message{"System Error retrieving Teams"})
    } else {
        c.JSON(200, teams)
    }
}

func GetTeam (c *gin.Context) {
    teamId := c.Params.ByName("teamId")
    team, err := dao.GetTeam(teamId)

    if err != nil {
        c.JSON(500, mod.Message{"Error retrieving Team"})
    } else if team.Id == "" {
        c.JSON(404, mod.Get404("Team", teamId))
    } else {
        c.JSON(200, team)
    }
}

func CreateTeam (c *gin.Context) {
    var json mod.Team
    c.Bind(&json)

    team, err := dao.CreateTeam(json)

    if err == nil {
        c.JSON(201, team)
    } else {
        c.JSON(500, mod.Message{"Error creating Team"})
    }
}

func UpdateTeam (c *gin.Context) {
    var json mod.Team
    json.Id = c.Params.ByName("teamId")

    c.Bind(&json)

    team, err := dao.UpdateTeam(json)

    if err == nil {
        c.JSON(200, team)
    } else {
        c.JSON(500, mod.Message{"Error updating Team"})
    }
}

func DeleteTeam (c *gin.Context) {
    teamId := c.Params.ByName("teamId")

    err := dao.DeleteTeam(teamId)

    if err == nil {
        c.JSON(200, mod.Message{"Deleted Team"})
    } else {
        c.JSON(500, mod.Message{"Error Deleting Team"})
    }
}
