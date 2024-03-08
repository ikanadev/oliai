package embedding

import (
	"context"
	"oliapi/rest/config"

	openai "github.com/sashabaranov/go-openai"
)

func embedWithOpenAIText3Small(ctx context.Context, content string) ([]float32, error) {
	config := config.GetConfig()
	client := openai.NewClient(config.OpenAIKey)
	queryReq := openai.EmbeddingRequest{ //nolint:exhaustruct
		Input: []string{content},
		Model: openai.SmallEmbedding3,
	}

	queryResp, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		return nil, err
	}

	return queryResp.Data[0].Embedding, nil
}

func embedWithOpenAIText3Large(ctx context.Context, content string) ([]float32, error) {
	config := config.GetConfig()
	client := openai.NewClient(config.OpenAIKey)
	queryReq := openai.EmbeddingRequest{ //nolint:exhaustruct
		Input: []string{content},
		Model: openai.LargeEmbedding3,
	}

	queryResp, err := client.CreateEmbeddings(context.Background(), queryReq)
	if err != nil {
		return nil, err
	}

	return queryResp.Data[0].Embedding, nil
}
