package bot

import "testing"

func TestAddMessage(t *testing.T) {

	bot := NewChatBot("test-key")

	bot.AddMessage("user", "hello")

	if len(bot.History) != 1 {
		t.Fatalf("expected 1 message got %d", len(bot.History))
	}

	if bot.History[0].Role != "user" {
		t.Fatalf("expected role user")
	}

	if bot.History[0].Content != "hello" {
		t.Fatalf("expected content hello")
	}
}
func TestMultipleMessages(t *testing.T) {

	bot := NewChatBot("test-key")

	bot.AddMessage("user", "hello")
	bot.AddMessage("assistant", "hi")

	if len(bot.History) != 2 {
		t.Fatalf("expected 2 messages")
	}
}
