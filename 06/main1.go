package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ayush6624/go-chatgpt"
)

type Data struct {
	index   int    `json:"index"`
	message string `json:"message"`
}

type JsonRespone struct {
	id      string `json:"id"`
	object  string `json:"object"`
	choices []Data `json:"choices"`
}

func main() {
	fmt.Println("starting chat gpt simple")
	var question string

	fmt.Print("input : ")
	fmt.Scanln(&question)

	key := os.Getenv("OPENAI_KEY")

	c, err := chatgpt.NewClient(key)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	res, err := c.Send(ctx, &chatgpt.ChatCompletionRequest{
		Model: chatgpt.GPT4,
		Messages: []chatgpt.ChatMessage{
			{
				Role:    chatgpt.ChatGPTModelRoleSystem,
				Content: question,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	a, _ := json.MarshalIndent(res, "", "  ")
	log.Println(string(a))

	var jsonRes JsonRespone
	err = json.Unmarshal(a, &jsonRes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(jsonRes.choices[0].message)

}
