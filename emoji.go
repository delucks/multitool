package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// based on kyokomi/generateEmojiCodeMap... a little copying is better than a little dependency

const gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"

type gemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

func generateJson() (map[string]string, error) {
	res, err := http.Get(gemojiDBJsonURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	emojiFile, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var gs []gemojiEmoji
	if err := json.Unmarshal(emojiFile, &gs); err != nil {
		return nil, err
	}
	emojiCodeMap := make(map[string]string)
	for _, gemoji := range gs {
		for _, a := range gemoji.Aliases {
			emojiCodeMap[a] = gemoji.Emoji
		}
	}
	return emojiCodeMap, nil
}

func GetAllEmojis(_ []string, _ io.Reader) error {
	codeMap, err := generateJson()
	if err != nil {
		return err
	}
	for emojiname, character := range codeMap {
		fmt.Printf("%s\t%s\n", emojiname, character)
	}
	return nil
}
