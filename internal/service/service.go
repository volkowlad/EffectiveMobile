package service

import (
	ef "EffectiveMobile"
	"EffectiveMobile/internal/repository"
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

type Service struct {
	Music
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Music: NewMusicService(repos.Music),
	}
}
