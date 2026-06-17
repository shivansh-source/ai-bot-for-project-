package main

	import (
    "ai-bot-for-project/bot"
)

func main() {

    chatbot := bot.NewChatBot("API_KEY")

    chatbot.AddMessage("user", "Hello")

    chatbot.AddMessage(
        "assistant",
        "Hi, how can I help you?",
    )

    chatbot.PrintHistory()
}