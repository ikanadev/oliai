package domain

type Bot struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	GreetingMessage string `json:"-"`
	CustomPropmt    string `json:"-"`
	TimeData
}

// The category that a info text belongs to
type BlockCategory struct {
	Id   string `json:"id"`
	Name string `json:"text"`
	TimeData
}

// store a piece of block info to be embedded
type BlockInfo struct {
	Id      string `json:"id"`
	Content string `json:"content"`
	TimeData
}