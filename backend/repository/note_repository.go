package repository

import (
	"context"
	"notes_project/config"
	"notes_project/models"
)

// CreateNote insert note baru
func CreateNote(note models.Note) error {
	query := `
		INSERT INTO notes (title, content, image_url, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err := config.DB.Exec(context.Background(), query,
		note.Title, note.Content, note.ImageURL, note.UserID)
	return err
}

// GetNotes ambil semua note
func GetNotes() ([]models.Note, error) {
	query := `
		SELECT id, title, content, image_url, user_id, created_at, updated_at
		FROM notes
	`

	rows, err := config.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.ImageURL, &n.UserID, &n.CreatedAt, &n.UpdatedAt)
		if err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}

	return notes, nil
}

// GetNoteByID ambil note berdasarkan id
func GetNoteByID(id int) (models.Note, error) {
	query := `
		SELECT id, title, content, image_url, user_id, created_at, updated_at
		FROM notes
		WHERE id = $1
	`

	row := config.DB.QueryRow(context.Background(), query, id)
	var n models.Note
	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.ImageURL, &n.UserID, &n.CreatedAt, &n.UpdatedAt)
	return n, err
}

// DeleteNote hapus note sesuai user
func DeleteNote(id int, userID int) error {
	query := `
		DELETE FROM notes
		WHERE id = $1 AND user_id = $2
	`
	_, err := config.DB.Exec(context.Background(), query, id, userID)
	return err
}

func GetNotesByUserID(userID int) ([]models.Note, error) {
	query := `SELECT id, title, content, image_url, user_id, created_at, updated_at
	          FROM notes WHERE user_id=$1`

	rows, err := config.DB.Query(context.Background(), query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note
	for rows.Next() {
		var n models.Note
		if err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.ImageURL, &n.UserID, &n.CreatedAt, &n.UpdatedAt); err != nil {
			return nil, err
		}
		notes = append(notes, n)
	}
	return notes, nil
}

func GetNoteByIDAndUserID(id, userID int) (models.Note, error) {
	query := `SELECT id, title, content, image_url, user_id, created_at, updated_at
	          FROM notes WHERE id=$1 AND user_id=$2`

	row := config.DB.QueryRow(context.Background(), query, id, userID)

	var n models.Note
	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.ImageURL, &n.UserID, &n.CreatedAt, &n.UpdatedAt)
	return n, err
}

func UpdateNote(id int, userID int, updatedNote models.Note) error {
	query := `
		UPDATE notes
		SET title = $1,
		    content = $2,
		    image_url = $3,
		    updated_at = NOW()
		WHERE id = $4 AND user_id = $5
	`
	_, err := config.DB.Exec(context.Background(), query,
		updatedNote.Title, updatedNote.Content, updatedNote.ImageURL, id, userID)
	return err
}