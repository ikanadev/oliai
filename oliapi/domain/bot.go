package domain

import (
	"time"

	"github.com/google/uuid"
)

type Bot struct {
	ID              uuid.UUID      `json:"id"`
	Name            string         `json:"name"`
	GreetingMessage string         `json:"greetingMessage"`
	CustomPrompt    string         `json:"customPrompt"`
	EmbeddingModel  EmbeddingModel `json:"embeddingModel"`
}

type BotWithTimeData struct {
	Bot
	TimeData
}

// The category that a info text belongs to.
type Category struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"text"`
}

type CategoryWithTimeData struct {
	Category
	TimeData
}

// store a piece of block info to be embedded.
type Document struct {
	ID            uuid.UUID `json:"id"`
	Content       string    `json:"content"`
	EmbeddingDate *time.Time `json:"embeddingDate"`
}

type DocumentWithTimeData struct {
	Document
	TimeData
}
