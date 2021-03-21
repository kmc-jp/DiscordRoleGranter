package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http/cgi"
)

var Settings Setting

func init() {
	var err error

	b, err := ioutil.ReadFile(SettingsFilePath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(b, &Settings)
	if err != nil {
		panic(err)
	}

	if Settings.Discord.Token == "" {
		fmt.Println("No token provided. Please run: airhorn -t <bot token>")
		return
	}

}

func main() {
	cgi.Serve(&Serve{})
}
