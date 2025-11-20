package models

type RegisterInput struct {
    Username string `json:"username" binding:"required"` // tambahkan username
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type LoginInput struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type User struct {
    ID        int    `json:"id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    Password  string `json:"password"`
    CreatedAt string `json:"created_at"`
}
