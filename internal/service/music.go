package service

import (
	ef "EffectiveMobile"
	"EffectiveMobile/internal/repository"
)

type MusicService struct {
	repos repository.Music
}

func NewMusicService(repos repository.Music) *MusicService {
	return &MusicService{
		repos: repos,
	}
}

func (m *MusicService) CreateSong(song ef.Song) (int, error) {
	return m.repos.CreateSong(song)
}

func (m *MusicService) CreateVerse(songID int, verse string) error {
	return m.repos.CreateVerse(songID, verse)
}

func (m *MusicService) CreateInfo(info ef.Info) (int, error) {
	return m.repos.CreateInfo(info)
}

func (m *MusicService) GetLibrary() ([]ef.Song, error) {
	return m.repos.GetLibrary()
}

func (m *MusicService) GetSong(songID int) (ef.Song, ef.Info, ef.Text, error) {
	return m.repos.GetSong(songID)
}

func (m *MusicService) UpdateSong(info ef.Info) error {
	return m.repos.UpdateSong(info)
}

func (m *MusicService) DeleteSong(songID int) error {
	return m.repos.DeleteSong(songID)
}
