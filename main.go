package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"ai-bot-for-project/bot"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	apiKey := os.Getenv("GROQ_API_KEY")

	chatbot := bot.NewChatBot(apiKey)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("AI Bot Started")
	fmt.Println("Type exit to quit")

	for {

		scanner.Scan()

		prompt := strings.TrimSpace(scanner.Text())

		if prompt == "exit" {
			break
		}

		response, err := chatbot.GenerateResponse(prompt)

		if err != nil {
			fmt.Println("error:", err)
			continue
		}

		fmt.Println()
		fmt.Println(response)
		fmt.Println()
	}
}