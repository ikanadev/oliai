package domain

type EmbeddingModel string

const (
	EmbeddingModelOpenAI3Small EmbeddingModel = "text-embedding-3-small"
	EmbeddingModelOpenAI3Large EmbeddingModel = "text-embedding-3-large"
)

type EmbeddingProvider struct {
	Provider   string
	Model      EmbeddingModel
	VectorSize int
	MaxTokens  int
}

var (
	EmbeddingOpenAI3Small = EmbeddingProvider{ //nolint:gochecknoglobals
		Provider:   "openai",
		Model:      EmbeddingModelOpenAI3Small,
		VectorSize: 1536, //nolint:gomnd
		MaxTokens:  8191, //nolint:gomnd
	}

	EmbeddingOpenAI3Large = EmbeddingProvider{ //nolint:gochecknoglobals
		Provider:   "openai",
		Model:      EmbeddingModelOpenAI3Large,
		VectorSize: 3072, //nolint:gomnd
		MaxTokens:  8191, //nolint:gomnd
	}
)

func EmbeddingProviderFromModel(s string) EmbeddingProvider {
	switch s {
	case "text-embedding-3-small":
		return EmbeddingOpenAI3Small
	case "text-embedding-3-large":
		return EmbeddingOpenAI3Large
	}

	return EmbeddingOpenAI3Small
}
