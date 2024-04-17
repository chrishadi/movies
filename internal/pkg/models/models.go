package models

import (
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model            // contains ID, CreatedAt, UpdatedAt, and DeletedAt columns
	ID          uint      `jsonapi:"primary,movies"`
	Title       string    `gorm:"unique" jsonapi:"attr,title"`
	Genre       string    `jsonapi:"attr,genre"`
	ReleaseYear uint16    `jsonapi:"attr,releaseYear"`
	DirectorID  uint      `jsonapi:"attr,directorId"`
	Director    *Director `jsonapi:"relation,director"`
}

type Movie2 struct {
	gorm.Model           // contains ID, CreatedAt, UpdatedAt, and DeletedAt columns
	ID          uint     `jsonapi:"primary,movies"`
	Title       string   `gorm:"unique" jsonapi:"attr,title"`
	Genre       string   `jsonapi:"attr,genre"`
	ReleaseYear uint16   `jsonapi:"attr,releaseYear"`
	DirectorID  uint     `jsonapi:"attr,directorId"`
	Director    Director `jsonapi:"relation,director"`
}

type Director struct {
	gorm.Model
	ID        uint       `jsonapi:"primary,directors"`
	Name      string     `gorm:"unique" jsonapi:"attr,name"`
	Birthdate *time.Time `jsonapi:"attr,birthdate,omitempty"`
	Movies    []Movie    `jsonapi:"relation,movies"`
}
