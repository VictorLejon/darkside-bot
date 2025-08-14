package main

import (
	"darkside-bot/internal/config"
	"flag"
	"log"
	"fmt"
	"encoding/json"
	"net/http"
	"bytes"
	"io"
	"os"
)

func main() {

	cfg := config.Load();

	var (
		guildID string
		name	string
		desc	string
		global	bool
	)

	flag.StringVar(&guildID, "guild", "", "Guild ID (omit for global)")
	flag.StringVar(&name, "name", "", "command name (lowercase, 1-32)")
	flag.StringVar(&desc, "desc", "", "command description")
	flag.BoolVar(&global, "global", false, "force global even if -guild set")
	flag.Parse()

	if cfg.AppID == "" || cfg.BotToken == "" || name == "" || (!global && guildID == "" && os.Getenv("GUILD_ID") == "") {
		log.Fatal("need DISCORD_APP_ID, DISCORD_BOT_TOKEN, -name, and either -guild or -global (GUILD_ID env also accepted)")
	}
	
	cmd := map[string]any{
		"name":			name,
		"description":	desc,
		"type":			1,
		// "options": []any{}
	}

	var discordUrl string
	if global {
		discordUrl = fmt.Sprintf("https://discord.com/api/v10/applications/%s/commands", cfg.AppID)
	} else {
		discordUrl = fmt.Sprintf("https://discord.com/api/v10/applications/%s/guilds/%s/commands", cfg.AppID, guildID)
	}

	fmt.Println(discordUrl)
	if err := sendReq(discordUrl, cfg.BotToken, cmd); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Registered new command!")
}

func sendReq(url, botToken string, body any) error {
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", url, bytes.NewReader(b))
	req.Header.Set("Authorization", "Bot "+botToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	slurp, _ := io.ReadAll(resp.Body)
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("discord %s: %s", resp.Status, string(slurp))
	}
	fmt.Printlnf("Command: %s", string(slurp)
	return nil
}






