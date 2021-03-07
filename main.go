package main

import (
	"google.golang.org/api/texttospeech/v1"
	"io/ioutil"
	"log"
	"os"
)

var (
	textFilePath = os.Args[1]
)

func main() {
	content, err := ioutil.ReadFile(textFilePath)
	if err != nil {
		log.Fatal(err)
	}
	speech := htgotts.Speech{Folder: "audio", Language: "en"}
	speech.Speak(string(content))
}
