package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FinalWork writes and flush response as JSON
func FinalWork(w http.ResponseWriter, res interface{}, status ...int) {
	retStatus := http.StatusOK
	if len(status) > 0 {
		retStatus = status[0]
	}
	js, err := json.Marshal(res)
	if err != nil {
		fmt.Errorf("Error in unmarshalling response: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(retStatus)
	_, err = w.Write(js)
	if err != nil {
		fmt.Errorf("Error in writing response: %v", err)
		return
	}
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
