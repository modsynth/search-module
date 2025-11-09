package search

import (
	"context"
)

type SearchClient interface {
	Index(ctx context.Context, index, id string, doc interface{}) error
	Search(ctx context.Context, index string, query map[string]interface{}) ([]map[string]interface{}, error)
	Delete(ctx context.Context, index, id string) error
}

type ElasticsearchClient struct {
	url string
}

func NewClient(url string) *ElasticsearchClient {
	return &ElasticsearchClient{url: url}
}

func (c *ElasticsearchClient) Index(ctx context.Context, index, id string, doc interface{}) error {
	// Implementation for indexing document
	return nil
}

func (c *ElasticsearchClient) Search(ctx context.Context, index string, query map[string]interface{}) ([]map[string]interface{}, error) {
	// Implementation for searching
	return nil, nil
}

func (c *ElasticsearchClient) Delete(ctx context.Context, index, id string) error {
	// Implementation for deleting document
	return nil
}

func BuildMatchQuery(field, value string) map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				field: value,
			},
		},
	}
}

func BuildTermQuery(field, value string) map[string]interface{} {
	return map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				field: value,
			},
		},
	}
}
