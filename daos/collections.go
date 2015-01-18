package daos

import (
    "log"
    mod "../models"
)

func GetCollection(members interface{}, table string) interface{} {
    collection := new (mod.Collection)

    response, err := GetList(table)
    if err != nil {
        log.Panic(err)
    }

    err = response.All(members)
    if err != nil {
        log.Panic(err)
    }

    count, err := GetCount(table)
    if err != nil {
        log.Panic(err)
    }

    collection.Members = members
    collection.Total = count

    return collection
}
