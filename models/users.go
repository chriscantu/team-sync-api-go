package models

type User struct {
    FirstName string `json:"firstName" required`
    LastName string `json:"lastName" required`
    Email string `json:"email" required`
}

type Users struct {
    Users []User `json:"users" required`
    Total int `json:"total"`
}
