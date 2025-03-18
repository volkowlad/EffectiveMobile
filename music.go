package EffectiveMobile

type Song struct {
	ID    int    `json:"id"`
	Name  string `json:"song" required:"true"`
	Group string `json:"group" required:"true"`
}

type Info struct {
	ID          int    `json:"id"`
	SongID      int    `json:"song_id"`
	ReleaseDate string `json:"release_date" required:"true"`
	Link        string `json:"link" required:"true"`
	Chorus      string `json:"chorus" required:"true"`
}

type Text struct {
	ID     int      `json:"id"`
	SongID string   `json:"song"`
	Verse  []string `json:"verse"`
}
