package handler

import (
	ef "EffectiveMobile"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

// @Summary Create song
// @Tags Songs
// @Description create song
// @ID create-song
// @Accept  json
// @Produce  json
// @Param input body EffectiveMobile.Song true "song info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /music [post]
func (h *Handler) createSong(c *gin.Context) {
	var input ef.Song
	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.service.CreateSong(input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

	slog.Info("Created wallet", id)
}

type createVerse struct {
	Text string `json:"text"`
}

// @Summary Create verse
// @Tags Songs
// @Description create verse
// @ID create-verse
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /music/:id [post]
func (h *Handler) createVerse(c *gin.Context) {
	songID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input createVerse
	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())
	}

	err = h.service.CreateVerse(songID, input.Text)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"add verse, song id": songID,
	})
}

func (h *Handler) createInfo(c *gin.Context) {
	var input ef.Info
	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateInfo(input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"add info, song id": id,
	})
}

type getLibrary struct {
	Library []ef.Song `json:"library"`
}

// @Summary Get library
// @Tags Songs
// @Description get library
// @ID get-library
// @Accept  json
// @Produce  json
// @Success 200 {integer} getLibrary
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /music [get]
func (h *Handler) getLibrary(c *gin.Context) {
	songs, err := h.service.GetLibrary()
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getLibrary{
		Library: songs,
	})

	slog.Info("Get library", len(songs))
}

type getSong struct {
	SongName    string `json:"song_name"`
	SongGroup   string `json:"song_group"`
	ReleaseData string `json:"release_data"`
	Link        string `json:"link"`
	Text        string `json:"text"`
	Message     string `json:"message"`
}

// @Summary Get song
// @Tags Songs
// @Description get song
// @ID get-song
// @Accept  json
// @Produce  json
// @Success 200 {integer} getSong
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /music/:id [get]
func (h *Handler) getSong(c *gin.Context) {
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	song, info, text, err := h.service.GetSong(input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	page := c.Query("page")
	switch page {
	case "0":
		c.JSON(http.StatusOK, getSong{
			Text: info.Chorus,
		})
	case "1":
		c.JSON(http.StatusOK, getSong{
			Text: text.Verse[0],
		})
	case "2":
		c.JSON(http.StatusOK, getSong{
			Text: text.Verse[1],
		})
	case "3":
		c.JSON(http.StatusOK, getSong{
			Text: text.Verse[2],
		})
	case "info":
		c.JSON(http.StatusOK, getSong{
			SongName:    song.Name,
			SongGroup:   song.Group,
			ReleaseData: info.ReleaseDate,
			Link:        info.Link,
		})
	default:
		c.JSON(http.StatusOK, getSong{
			Message: `
на 0 страницу припев песни
на 1-3 страницах куплеты песни
на странице info название песни, дата релиза и ссылка для прослушивания`,
		})
	}

	slog.Info("Get info", song.Name)
}

// @Summary Update song
// @Tags Songs
// @Description update song
// @ID update-song
// @Accept  json
// @Produce  json
// @Param input body EffectiveMobile.Info true "song info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /music [put]
func (h *Handler) updateSong(c *gin.Context) {
	var input ef.Info

	if err := c.BindJSON(&input); err != nil {
		NewRespError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.UpdateSong(input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"update_song": input.SongID,
	})

	slog.Info("Update song", input.SongID)
}

// @Summary Delete song
// @Tags Songs
// @Description delete song
// @ID delete-song
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} respError
// @Failure 500 {object} respError
// @Failure default {object} respError
// @Router /music/:id [get]
func (h *Handler) deleteSong(c *gin.Context) {
	input, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.DeleteSong(input)
	if err != nil {
		NewRespError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"delete_song": input,
	})
}
