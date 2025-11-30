package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	pg "task_2/PostgreSQL"
)

type RequestData struct {
	Data map[string]string `json:"data"`
}

// TODO: Добавить syncpool и описание
func NewRequestData() *RequestData {
	return &RequestData{
		Data: make(map[string]string),
	}
}

func HandleCreateData(postgres *pg.Postgres) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST-method allowed", http.StatusMethodNotAllowed)
			return
		}

		req := NewRequestData()

		err := json.NewDecoder(r.Body).Decode(&req.Data)
		if err != nil {
			http.Error(w, "invalid JSON "+err.Error(), http.StatusBadRequest)
			return
		}
		conn := postgres.GetConn()
		defer postgres.Release(conn)
		fmt.Printf("Processing request on connection %d...\n", conn.ID)
		for key, value := range req.Data {
			postgres.Create(key, value)
		}

	}
}

func HandleGetData(postgres *pg.Postgres) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "only GET-methods allowed", http.StatusMethodNotAllowed)
			return
		}

		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing \"key\" query parameter", http.StatusBadRequest)
			return
		}

		conn := postgres.GetConn()
		defer postgres.Release(conn)

		value, err := postgres.GetData(key)

		if err != nil {
			http.Error(w, "not found: "+err.Error(), http.StatusNotFound)
		}

		response := map[string]string{
			key: value,
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}
}
