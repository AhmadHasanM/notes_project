package services

import (
	"notes_project/models"
	"notes_project/repository"
)

// Buat note baru
func CreateNoteService(note models.Note) error {
	return repository.CreateNote(note)
}

/// services/note_service.go

func GetNotesService(userID int) ([]models.Note, error) {
	return repository.GetNotesByUserID(userID)
}

func GetNoteDetailService(id int, userID int) (models.NoteResponse, error) {
	note, err := repository.GetNoteByIDAndUserID(id, userID)
	if err != nil {
		return models.NoteResponse{}, err
	}

	noteResp := models.NoteResponse{
		ID:        note.ID,
		UserID:    note.UserID,
		Title:     note.Title,
		Content:   note.Content,
		ImageURL:  note.ImageURL,
		CreatedAt: note.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return noteResp, nil
}


// Delete note
func DeleteNoteService(id int, userID int) error {
	return repository.DeleteNote(id, userID)
}

func UpdateNoteService(id int, userID int, updatedNote models.Note) error {
	return repository.UpdateNote(id, userID, updatedNote)
}
