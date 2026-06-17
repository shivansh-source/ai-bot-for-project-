package bot

type ChatBot struct {
    APIKey  string
    History []Message
}

func NewChatBot(apiKey string) *ChatBot {
    return &ChatBot{
        APIKey: apiKey,
        History: []Message{},
    }
}
func (c *ChatBot) AddMessage(role , content string ) {
    c.History = append(c.History, Message{
		Role: role,
		Content: content,
	})
}

func (c *ChatBot) PrintHistory(){
    for _, message := range c.History {
		println(message.Role + ": " + message.Content)
	}
}
