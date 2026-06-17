package bot 
type Chatbot struct {
	APIKey string
}
func NewChatbot(apiKey string) *Chatbot {
	return &Chatbot{
		APIKey: apiKey,
	}
}