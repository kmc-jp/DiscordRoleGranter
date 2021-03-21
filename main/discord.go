package main

import (
	"github.com/bwmarrin/discordgo"
)

func NewDiscord() (*discordgo.Session, error) {
	return discordgo.New("Bot " + Settings.Discord.Token)
}
