package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

var url = ""

const unicorn = "/md ![unicorn](https://media.giphy.com/media/26AHG5KGFxSkUWw1i/giphy.gif) Good morning!"

type message struct {
	Content string `json:"Content"`
}

func init() {
	url = os.Getenv("URL")
	if len(url) < 1 {
		panic("You must pass in the URL environment variable")
	}
}

func ChimeHandler() error {
	msg, err := json.Marshal(&message{Content: unicorn})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(msg)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func main() {
	lambda.Start(ChimeHandler)
}
