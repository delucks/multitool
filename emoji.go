package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

// based on kyokomi/generateEmojiCodeMap... a little copying is better than a little dependency

const (
	gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
	localCacheFile  = ".cache/emoji.json"
)

func relativeToHome(path string) (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(user.HomeDir, path)
	return fullPath, nil
}

type gemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

func downloadLocally() error {
	absPath, err := relativeToHome(localCacheFile)
	if err != nil {
		return err
	}
	res, err := http.Get(gemojiDBJsonURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	out, err := os.Create(absPath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

func generateJson() (map[string]string, error) {
	expiration, err := time.ParseDuration("720h")
	if err != nil {
		return nil, err
	}
	absPath, err := relativeToHome(localCacheFile)
	if err != nil {
		return nil, err
	}
	stat, err := os.Stat(absPath)
	// If the file doesn't exist or is >30d old, redownload it
	if os.IsNotExist(err) || time.Since(stat.ModTime()) > expiration {
		err = downloadLocally()
		if err != nil {
			return nil, err
		}
	}
	emojiFile, err := os.Open(absPath)
	defer emojiFile.Close()
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(emojiFile)
	if err != nil {
		return nil, err
	}
	var gs []gemojiEmoji
	if err := json.Unmarshal(contents, &gs); err != nil {
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
