package bot

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
)

var (
	// Discord holds the live session once started.
	Discord *discordgo.Session

	once sync.Once
	mu   sync.RWMutex
)

// Start initializes and opens the Discord session once.
// Call this from your application startup (or let callers trigger it lazily).
func Start() error {
	var startErr error
	once.Do(func() {
		token := os.Getenv("DISCORD_BOT_TOKEN")
		if token == "" {
			startErr = fmt.Errorf("DISCORD_BOT_TOKEN not set")
			return
		}

		s, err := discordgo.New("Bot " + token)
		if err != nil {
			startErr = err
			return
		}

		s.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMembers | discordgo.IntentsGuildEmojis | discordgo.IntentsGuildPresences

		if err := s.Open(); err != nil {
			_ = s.Close()
			startErr = err
			return
		}

		mu.Lock()
		Discord = s
		mu.Unlock()

		// Background monitor to keep package aware of session lifecycle.
		go monitor()
	})
	return startErr
}

// Get returns the current session (may be nil if not started).
func Get() *discordgo.Session {
	mu.RLock()
	defer mu.RUnlock()
	return Discord
}

// Close shuts down the session if active.
func Close() error {
	mu.Lock()
	defer mu.Unlock()
	if Discord == nil {
		return nil
	}
	err := Discord.Close()
	Discord = nil
	return err
}

// monitor logs / watches session and exits when session is nil.
// You can expand this to reconnect logic if desired.
func monitor() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		mu.RLock()
		active := Discord != nil
		mu.RUnlock()
		if !active {
			return
		}
	}
}
