package controllers

import (
    "github.com/gin-gonic/gin"
    dao "../daos"
    mod "../models"
    )

func GetReports (c *gin.Context) {
    reports, err := dao.GetReports()

    if err != nil {
        c.JSON(500, mod.Message{"System Error retrieving Reports"})
    } else {
        c.JSON(200, reports)
    }
}

func GetReport (c *gin.Context) {
    reportId := c.Params.ByName("reportId")
    report := dao.GetReport(reportId)

    if report.Id == "" {
        c.JSON(404, mod.Get404("Report", reportId))
    } else {
        c.JSON(200, report)
    }
}

func CreateReport (c *gin.Context) {
    var json mod.Report
    c.Bind(&json)

    report, err := dao.CreateReport(json)

    if err == nil {
        c.JSON(201, report)
    } else {
        c.JSON(500, mod.Message{"Error creating Report"})
    }
}

func UpdateReport (c *gin.Context) {
    var json mod.Report
    json.Id = c.Params.ByName("reportId")

    c.Bind(&json)

    report, err := dao.UpdateReport(json)

    if err == nil {
        c.JSON(200, report)
    } else {
        c.JSON(500, mod.Message{"Error updating Report"})
    }
}

func DeleteReport (c *gin.Context) {
    reportId := c.Params.ByName("reportId")

    err := dao.DeleteReport(reportId)

    if err == nil {
        c.JSON(200, mod.Message{"Deleted Report"})
    } else {
        c.JSON(500, mod.Message{"Error Deleting Report"})
    }
}
