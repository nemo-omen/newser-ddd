package entity

import (
	"newser/internal/domain/value"
)

type NoteId = value.ID

type Note struct {
	ID      NoteId
	Title   string
	Content string
}
