package httpserver

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type ibanHandler struct {
	logger      *log.Logger
	ibanService ibanService
}

func newIBANHandler(ibanService ibanService, logger *log.Logger) *ibanHandler {
	return &ibanHandler{
		ibanService: ibanService,
		logger:      logger,
	}
}

func (h *ibanHandler) Routes(router chi.Router) {
	// Route for validating an iban number.
	router.Get("/{iban}", h.validateIBAN)
}

// validateIBAN handles the request to validate an IBAN number.
func (h *ibanHandler) validateIBAN(w http.ResponseWriter, r *http.Request) {
	iban := chi.URLParam(r, "iban")
	valid, err := h.ibanService.ValidateIBAN(r.Context(), iban)

	response := validateIBANResponse{Valid: valid}
	if err != nil {
		response.Reason = err.Error()
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		h.logger.Printf("failed to marshal response: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(jsonResponse)
}
