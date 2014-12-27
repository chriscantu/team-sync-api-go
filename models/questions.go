package models

type Question struct {
    Id string `json:"id"`
    Question string `json:"question" required`
}
