package helper

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func SendJSON(w http.ResponseWriter, statusCode int, payload any) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    data, err := json.Marshal(payload)
    if err != nil {
        log.Println("ERROR:", err)
    }
    w.Write(data)
}

func ParseBody(r *http.Request, data any) {
    body, err := io.ReadAll(r.Body)
    if err != nil {
        log.Println("ERROR:", err)
    }
    json.Unmarshal(body, data)
}