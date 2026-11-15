package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type errorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type errorEnvelope struct {
	Error errorBody `json:"error"`
}

type successEnvelope struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// JSON mengirim response JSON dengan status yang diberikan.
func JSON(w http.ResponseWriter, status int, data any) {
	// Encode dulu sebelum menulis header. Setelah WriteHeader dipanggil,
	// status tidak bisa diubah lagi kalau ternyata encode gagal.
	body, err := json.Marshal(data)
	if err != nil {
		log.Printf("[ERROR] gagal encode response: %v", err)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"Terjadi kesalahan pada server"}}`))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if _, err := w.Write(body); err != nil {
		log.Printf("[ERROR] gagal menulis response: %v", err)
	}
}

// Success mengirim response sukses dengan pesan dan data.
func Success(w http.ResponseWriter, status int, message string, data any) {
	JSON(w, status, successEnvelope{
		Message: message,
		Data:    data,
	})
}

// Error mengirim response gagal dengan format yang konsisten.
func Error(w http.ResponseWriter, status int, code, message string) {
	JSON(w, status, errorEnvelope{Error: errorBody{Code: code, Message: message}})
}

// NoContent untuk response tanpa body, misalnya DELETE yang berhasil (204).
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
