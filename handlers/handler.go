package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brltd/delivery/logger"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error(fmt.Sprintf("Error encoding response %+v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
