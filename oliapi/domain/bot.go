package domain

import "github.com/google/uuid"

type Bot struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	GreetingMessage string    `json:"greetingMessage"`
	CustomPrompt    string    `json:"customPrompt"`
}

type BotWithTimeData struct {
	Bot
	TimeData
}

type BotWithCategories struct {
	Bot
	Categories []CategoryWithDocuments `json:"categories"`
}

// The category that a info text belongs to.
type Category struct {
	ID   string `json:"id"`
	Name string `json:"text"`
}

type CategoryWithDocuments struct {
	Category
	Documents []Document `json:"documents"`
}

// store a piece of block info to be embedded.
type Document struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
