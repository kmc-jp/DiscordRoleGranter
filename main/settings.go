package main

const SettingsFilePath = "settings.json"

type Setting struct {
	Discord struct {
		Token   string
		GuildID string
		RoleID  string
	}
}
