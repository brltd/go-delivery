package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brltd/delivery/logger"
)

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	encodedData, err := json.Marshal(data)
	if err != nil {
		logger.Error(fmt.Sprintf("Error encoding response %+v", err))
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(code)
	w.Write(encodedData)
}
