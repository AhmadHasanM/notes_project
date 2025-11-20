package controllers

import (
	"context"
	"net/http"
	"os"
	"time"

	"notes_project/config"
	"notes_project/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// fungsi untuk membaca JWT_SECRET saat dibutuhkan
func getJWTKey() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET belum di-set")
	}
	return []byte(secret)
}

// Register
func Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	// cek apakah email sudah ada
	var exists bool
	err := config.DB.QueryRow(context.Background(),
		"SELECT EXISTS(SELECT 1 FROM users WHERE email=$1)", input.Email).Scan(&exists)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal cek email"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email sudah terdaftar"})
		return
	}

	// hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// insert user dengan username
	_, err = config.DB.Exec(context.Background(),
		"INSERT INTO users (username, email, password, created_at) VALUES ($1, $2, $3, $4)",
		input.Username, input.Email, string(hashedPassword), time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal daftar user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user berhasil didaftarkan"})
}

// Login
func Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	var user models.User
	err := config.DB.QueryRow(context.Background(),
		"SELECT id, username, email, password FROM users WHERE email=$1", input.Email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email atau password salah"})
		return
	}

	// cek password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email atau password salah"})
		return
	}

	// buat JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(getJWTKey())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal buat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
