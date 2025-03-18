package repository

import (
	ef "EffectiveMobile"
	"database/sql"
)

type Music interface {
	CreateSong(song ef.Song) (int, error)
	CreateVerse(songID int, verse string) error
	CreateInfo(info ef.Info) (int, error)
	GetLibrary() ([]ef.Song, error)
	GetSong(songID int) (ef.Song, ef.Info, ef.Text, error)
	UpdateSong(info ef.Info) error
	DeleteSong(songID int) error
}

type Repository struct {
	Music
}

func NewRepos(db *sql.DB) *Repository {
	return &Repository{
		Music: NewMusicDB(db),
	}
}
