package controllers

import (
	"context"
	"net/http"
	"strconv"
	"notes_project/config"
	"time"
	"os"

	"github.com/gin-gonic/gin"
)

// Ambil semua notes
func GetNotes(c *gin.Context) {
	rows, err := config.DB.Query(context.Background(), "SELECT id, title, content, image_url, created_at, updated_at FROM notes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal ambil notes"})
		return
	}
	defer rows.Close()

	var notes []map[string]interface{}
	for rows.Next() {
		var n map[string]interface{} = make(map[string]interface{})
		var id int
		var title, content, imageURL string
		var createdAt, updatedAt time.Time
		rows.Scan(&id, &title, &content, &imageURL, &createdAt, &updatedAt)
		n["id"] = id
		n["title"] = title
		n["content"] = content
		n["image_url"] = imageURL
		n["created_at"] = createdAt
		n["updated_at"] = updatedAt
		notes = append(notes, n)
	}

	c.JSON(http.StatusOK, notes)
}

// Ambil note by ID
func GetNoteByID(c *gin.Context) {
	id := c.Param("id")
	var n map[string]interface{} = make(map[string]interface{})
	var title, content, imageURL string
	var createdAt, updatedAt time.Time

	err := config.DB.QueryRow(context.Background(),
		"SELECT title, content, image_url, created_at, updated_at FROM notes WHERE id=$1", id).
		Scan(&title, &content, &imageURL, &createdAt, &updatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "note tidak ditemukan"})
		return
	}

	n["id"] = id
	n["title"] = title
	n["content"] = content
	n["image_url"] = imageURL
	n["created_at"] = createdAt
	n["updated_at"] = updatedAt

	c.JSON(http.StatusOK, n)
}

// Hapus note
func DeleteNote(c *gin.Context) {
	id := c.Param("id")
	res, err := config.DB.Exec(context.Background(), "DELETE FROM notes WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menghapus note"})
		return
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "note tidak ditemukan"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "note berhasil dihapus"})
}

func CreateNote(c *gin.Context) {
    title := c.PostForm("title")
    content := c.PostForm("content")

    var imageURL string
    file, err := c.FormFile("image")
    if err == nil {
        // buat folder uploads jika belum ada
        os.MkdirAll("uploads", os.ModePerm)
        filePath := "uploads/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + file.Filename
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal upload gambar"})
            return
        }
        imageURL = filePath
    }

    _, err = config.DB.Exec(context.Background(),
        "INSERT INTO notes (title, content, image_url, created_at) VALUES ($1, $2, $3, $4)",
        title, content, imageURL, time.Now())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal membuat note"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "note berhasil dibuat"})
}

func UpdateNote(c *gin.Context) {
    id := c.Param("id")
    title := c.PostForm("title")
    content := c.PostForm("content")

    var imageURL string
    file, err := c.FormFile("image")
    if err == nil {
        os.MkdirAll("uploads", os.ModePerm)
        filePath := "uploads/" + strconv.FormatInt(time.Now().UnixNano(), 10) + "_" + file.Filename
        if err := c.SaveUploadedFile(file, filePath); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal upload gambar"})
            return
        }
        imageURL = filePath
    }

    if imageURL != "" {
        _, err = config.DB.Exec(context.Background(),
            "UPDATE notes SET title=$1, content=$2, image_url=$3, updated_at=$4 WHERE id=$5",
            title, content, imageURL, time.Now(), id)
    } else {
        _, err = config.DB.Exec(context.Background(),
            "UPDATE notes SET title=$1, content=$2, updated_at=$3 WHERE id=$4",
            title, content, time.Now(), id)
    }

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "gagal update note"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "note berhasil diupdate"})
}
