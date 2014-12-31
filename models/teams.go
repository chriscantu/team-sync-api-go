package models

type Team struct {
    Id string `json:"id"`
    Admin string `json:"admin"`
    Users []string `json:"users" required`
}

type Teams struct {
    Teams []Team `json:"teams" required`
    Total float64 `json:"total"`
}
