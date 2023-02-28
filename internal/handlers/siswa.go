package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/upn-belajar-go/internal/domain/master"
	"gitlab.com/upn-belajar-go/shared"
	"gitlab.com/upn-belajar-go/shared/failure"
	"gitlab.com/upn-belajar-go/transport/http/middleware"
	"gitlab.com/upn-belajar-go/transport/http/response"

	"github.com/go-chi/chi"
)

type SiswaHandler struct {
	SiswaService master.SiswaService
}

func ProvideSiswaHandler(service master.SiswaService) SiswaHandler {
	return SiswaHandler{
		SiswaService: service,
	}
}

func (h *SiswaHandler) Router(r chi.Router, middleware *middleware.JWT) {
	r.Route("/master/siswa", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			// r.Use(middleware.VerifyToken)
			r.Post("/", h.Create)
		})
	})
}

// createSiswa adalah untuk menambah data siswa.
// @Summary menambahkan data siswa.
// @Description Endpoint ini adalah untuk menambahkan data siswa.
// @Tags siswa
// @Produce json
// @Param siswa body master.RequestSiswaFormat true "Siswa yang akan ditambahkan"
// @Success 200 {object} response.Base{data=master.Siswa}
// @Failure 400 {object} response.Base
// @Failure 500 {object} response.Base
// @Router /v1/master/siswa [post]
func (h *SiswaHandler) Create(w http.ResponseWriter, r *http.Request) {
	var reqFormat master.RequestSiswaFormat
	err := json.NewDecoder(r.Body).Decode(&reqFormat)
	if err != nil {
		fmt.Print("error jsondecoder")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	// Validasi required
	err = shared.GetValidator().Struct(reqFormat)
	if err != nil {
		response.WithStatusMessage(w, http.StatusCreated, false, "Nama, Kelas Wajib diisi")
		return
	}

	// userID := middleware.GetClaimsValue(r.Context(), "userId").(string)
	userID := ""
	data, err := h.SiswaService.Create(reqFormat, userID)
	if err != nil {
		fmt.Print("error create")
		response.WithError(w, failure.BadRequest(err))
		return
	}

	response.WithJSON(w, http.StatusCreated, data)
}
