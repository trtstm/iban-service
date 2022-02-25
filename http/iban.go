package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type ibanHandler struct {
	logger *log.Logger
}

func newIBANHandler(logger *log.Logger) *ibanHandler {
	return &ibanHandler{
		logger: logger,
	}
}

func (h *ibanHandler) Routes(router chi.Router) {
	// Route for fetching a specific user on id.
	router.Get("/{iban}", h.validateIBAN)
}

// validateIBAN handles the request to validate an IBAN number.
func (h *ibanHandler) validateIBAN(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(validateIBANResponse{Valid: true})
	if err != nil {
		h.logger.Printf("failed to marshal response: %v", err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)
}
