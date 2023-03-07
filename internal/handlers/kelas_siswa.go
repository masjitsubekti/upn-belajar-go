package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"

	"github.com/go-chi/chi"
)

type KelasSiswaHandler struct {
	KelasSiswaService master.KelasSiswaService
}

func ProvideKelasSiswaHandler(service master.KelasSiswaService) KelasSiswaHandler {
	return KelasSiswaHandler{
		KelasSiswaService: service,
	}
}

func (h *KelasSiswaHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/kelas-siswa", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/", h.Create)
			r.Get("/cek-siswa", h.ExistByIDSiswa)
		})
	})
}

// createSiswa adalah untuk menambah data kelas siswa.
// @Summary menambahkan data kelas siswa.
// @Description Endpoint ini adalah untuk menambahkan data kelas siswa.
// @Tags kelas-siswa
// @Produce json
// @Param kelasSiswa body master.KelasSiswaRequest true "Kelas Siswa yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.KelasSiswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa [post]
func (h *KelasSiswaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.KelasSiswaRequest
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.KelasSiswaService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error create")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}

// ExistByIDSiswa adalah untuk mendapatkan satu data kelas siswa berdasarkan idSiswa, idKelasSiswa.
// @Summary Mendapatkan satu data kelas siswa berdasarkan idSiswa, idKelasSiswa.
// @Description Endpoint ini adalah untuk mendapatkan kelas siswa berdasarkan idSiswa, idKelasSiswa.
// @Tags kelas-siswa
// @Produce json
// @Param idSiswa query string true "Set idSiswa"
// @Param idKelasSiswa query string true "Set idKelasSiswa"
// @Success 200 {object} response.Base
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/kelas-siswa/cek-siswa [get]
func (h *KelasSiswaHandler) ExistByIDSiswa(w http.ResponseWriter, r *http.Request) {
	idSiswa := r.URL.Query().Get("idSiswa")
	idKelasSiswa := r.URL.Query().Get("idKelasSiswa")
	siswa, err := h.KelasSiswaService.ExistByIdSiswa(idSiswa, idKelasSiswa)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusOK, siswa)
}
