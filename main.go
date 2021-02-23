package main

import (
	"github.com/hegedustibor/htgo-tts"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("fileName")
	if err != nil {
		log.Fatal(err)
	}
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak(string(content))
}
