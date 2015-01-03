package daos

import (
    "log"
    db "github.com/dancannon/gorethink"
    mod "../models"
    "time"
    )

func GetReports() (*mod.Reports, error) {
    reports := []mod.Report{}
    reportsMod := new (mod.Reports)

    response, err := GetList("reports")

    err = response.All(&reports)
    count, err := GetCount("reports")

    reportsMod.Reports = reports
    reportsMod.Total = count

    return reportsMod, err
}

func GetReport(reportId string) (mod.Report) {
    report := mod.Report{}

    response := GetRec("reports", reportId)
    response.Next(&report)

    return report
}

func CreateReport(report mod.Report) (mod.Report, error) {
    session := GetSession()

    report.Date = time.Now()
    response, err := db.Table("reports").Insert(report).RunWrite(session)

    if err != nil {
        log.Panic(err)
    }

    report.Id = response.GeneratedKeys[0]

    return report, err
}

func UpdateReport(report mod.Report) (mod.Report, error) {
    session := GetSession()

    _, err := db.Table("reports").Get(report.Id).Update(report).RunWrite(session)

    return report, err
}

func DeleteReport(reportId string) (error) {
    return DeleteRec("reports", reportId)
}
