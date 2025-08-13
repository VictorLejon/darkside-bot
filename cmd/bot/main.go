package main

import (
	"log"
	"fmt"
	"darkside-bot/internal/config"
	"darkside-bot/internal/httpserver"
	"darkside-bot/internal/discord"
)

func main() {
	cfg := config.Load();

	r := discord.NewRouter();
	r.RegisterCommand("ping", func(i discord.Interaction) discord.InteractionResponse {
		return discord.RespMessage("pong", false);
	})


	srv := httpserver.New(cfg, r);
	fmt.Println("Starting...")
	log.Fatal(srv.ListenAndServe());
}
