//    Copyright 2018 Yoshi Yamaguchi
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/bwmarrin/discordgo"
)

type ChannelID string

const (
	InfoChannel = "503169629493919744"
)

type Secrets struct {
	ID      string
	Secrets string
	Token   string
}

func readSecrets() (*Secrets, error) {
	b, err := ioutil.ReadFile("secrets.json")
	if err != nil {
		return nil, err
	}
	var s Secrets
	if err = json.Unmarshal(b, &s); err != nil {
		return nil, err
	}
	return &s, nil
}

func main() {
	done := make(chan bool)
	discord, err := discordgo.New()
	s, err := readSecrets()
	if err != nil {
		log.Fatalf("Error logging in: %v", err)
	}
	discord.Token = s.Token
	discord.AddHandler(infoHandler)
	if err = discord.Open(); err != nil {
		log.Fatalf("Error opening discord: %v", err)
	}
	<-done
}

func infoHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot || m.ChannelID != InfoChannel {
		return
	}
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Printf("Error in handler: %v", err)
	}

	if _, err := s.ChannelMessageSend(c.ID, "yes"); err != nil {
		log.Printf("Error sending message: %v", err)
	}
}
