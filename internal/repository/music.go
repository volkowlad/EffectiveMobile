package repository

import (
	ef "EffectiveMobile"
	"database/sql"
	"fmt"
	"log/slog"
)

type MusicPostgres struct {
	db *sql.DB
}

func NewMusicDB(db *sql.DB) *MusicPostgres {
	return &MusicPostgres{db: db}
}

func (m *MusicPostgres) CreateSong(song ef.Song) (int, error) {
	var id int

	createWalletQuery := fmt.Sprintf(`
INSERT INTO %s (song_name, song_group)
VALUES ($1, $2)
RETURNING id`, songsTable)
	err := m.db.QueryRow(createWalletQuery, song.Name, song.Group).Scan(&id)
	if err != nil {
		slog.Error("failed to insert song into wallet", err.Error())
		return -1, err
	}

	return id, nil
}

func (m *MusicPostgres) CreateVerse(songID int, verse string) error {
	createVerseQuery := fmt.Sprintf(`INSERT INTO %s (song_id, verse) VALUES ($1, $2)`, textTable)
	_, err := m.db.Exec(createVerseQuery, songID, verse)
	if err != nil {
		slog.Error("failed to create verse into wallet", err.Error())
		return err
	}

	return nil
}

func (m *MusicPostgres) GetLibrary() ([]ef.Song, error) {
	var songs []ef.Song

	getQuery := fmt.Sprintf(`SELECT * FROM %s`, songsTable)

	rows, err := m.db.Query(getQuery)
	if err != nil {
		slog.Error("failed to get library", err.Error())
		return nil, err
	}

	for rows.Next() {
		var song ef.Song
		if err := rows.Scan(&song.ID, &song.Name, &song.Group); err != nil {
			slog.Error("failed to scan song from wallet to get library", err.Error())
			return songs, err
		}

		songs = append(songs, song)
	}
	if err := rows.Err(); err != nil {
		slog.Error("failed to get library", err.Error())
		return songs, err
	}

	return songs, nil
}

func (m *MusicPostgres) GetSong(songID int) (ef.Song, ef.Info, ef.Text, error) {
	var song ef.Song
	var info ef.Info
	var text ef.Text

	tx, err := m.db.Begin()
	if err != nil {
		slog.Error("failed to begin transaction", err.Error())
		return song, info, text, err
	}

	getSong := fmt.Sprintf(`SELECT * FROM %s WHERE id=$1`, songsTable)
	err = tx.QueryRow(getSong, songID).Scan(&song.ID, &song.Name, &song.Group)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to get song from wallet", err.Error())
		return song, info, text, err
	}

	getInfo := fmt.Sprintf(`SELECT * FROM %s WHERE song_id=$1`, infoTable)
	err = tx.QueryRow(getInfo, songID).Scan(&info.ID, &info.SongID, &info.ReleaseDate, &info.Link, &info.Chorus)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to get song", err.Error())
		return song, info, text, err
	}

	getText := fmt.Sprintf(`SELECT verse FROM %s WHERE song_id=$1`, textTable)
	rows, err := tx.Query(getText, songID)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to query verses", err.Error())
		return song, info, text, err
	}

	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			tx.Rollback()
			slog.Error("failed to get one verse", err.Error())
			return song, info, text, err
		}

		text.Verse = append(text.Verse, verse)
	}
	if err := rows.Err(); err != nil {
		tx.Rollback()
		slog.Error("failed to get verses", err.Error())
		return song, info, text, err
	}

	return song, info, text, tx.Commit()
}

func (m *MusicPostgres) UpdateSong(info ef.Info) error {
	updateQuery := fmt.Sprintf(`UPDATE %s SET release_date=$1, link=$2, chorus=$3 WHERE song_id=$4`, infoTable)

	_, err := m.db.Exec(updateQuery, info.ReleaseDate, info.Link, info.Chorus, info.SongID)
	if err != nil {
		slog.Error("failed to update song", err.Error())
		return err
	}

	return nil
}

func (m *MusicPostgres) CreateInfo(info ef.Info) (int, error) {
	var id int

	tx, err := m.db.Begin()
	if err != nil {
		slog.Error("failed to begin transaction", err.Error())
		return -1, err
	}

	createInfoQuery := fmt.Sprintf(`INSERT INTO %s (song_id, release_date, link, chorus) VALUES ($1, $2, $3, $4) RETURNING id`, infoTable)

	err = tx.QueryRow(createInfoQuery, info.SongID, info.ReleaseDate, info.Link, info.Chorus).Scan(&id)
	if err != nil {
		tx.Rollback()
		slog.Error("failed to create info into song", err.Error())
		return -1, err
	}

	return id, tx.Commit()
}

func (m *MusicPostgres) DeleteSong(songID int) error {
	deleteQuery := fmt.Sprintf(`DELETE FROM %s WHERE id=$1`, songsTable)
	_, err := m.db.Exec(deleteQuery, songID)
	if err != nil {
		slog.Error("failed to delete song", err.Error())
		return err
	}

	return nil
}
