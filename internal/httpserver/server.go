package httpserver

import (
	"io"
	"encoding/json"
	"net/http"
	"darkside-bot/internal/discord"
	"darkside-bot/internal/config"
)


func New(cfg config.Config, router *discord.Router) *http.Server {
	mux := http.NewServeMux()
	
	mux.HandleFunc("/interactions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if (err != nil){
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		if !discord.VerifySignature(r, cfg.PublicKey, body) {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		var it discord.Interaction
		if err := json.Unmarshal(body, &it); err != nil {
			http.Error(w, "bad json", http.StatusBadRequest)
			return
		}

		resp := router.Dispatch(it)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	});
	
	return &http.Server{
		Addr: ":" + cfg.Port,
		Handler: mux,
	}
}

