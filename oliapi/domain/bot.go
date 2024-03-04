package domain

type Bot struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	GreetingMessage string `json:"-"`
	CustomPropmt    string `json:"-"`
}

type BotWithCategories struct {
	Bot
	Categories []Category `json:"categories"`
}

// The category that a info text belongs to.
type Category struct {
	ID   string `json:"id"`
	Name string `json:"text"`
}

// store a piece of block info to be embedded.
type Document struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}
