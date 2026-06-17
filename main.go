package main

import (
    "fmt"

    "ai-bot-for-project/bot"
)

func main() {

    chatbot := bot.NewChatbot("YOUR_API_KEY")

    fmt.Println("Bot Created")
    fmt.Println(chatbot)
}