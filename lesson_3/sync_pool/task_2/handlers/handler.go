package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	pg "task_2/PostgreSQL"
	sp "task_2/syncPool"
)

func HandleCreateDataWithPool(postgres *pg.Postgres, pool *sp.SyncPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST-method allowed", http.StatusMethodNotAllowed)
			return
		}

		req := pool.Get()
		defer pool.Put(req)

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

		fmt.Println("objects created:", atomic.LoadInt64(&sp.Created))
	}
}

func HandleCreateDataNoPool(postgres *pg.Postgres) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST-method allowed", http.StatusMethodNotAllowed)
			return
		}

		req := sp.NewRequestData()

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
		fmt.Println("objects created:", atomic.LoadInt64(&sp.Created))
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
