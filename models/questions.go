package models

type Question struct {
    Id string `json:"id"`
    Question string `json:"question" required`
}

type Questions struct {
    Questions []Question `json:"users" required`
    Total float64 `json:"total"`
}
