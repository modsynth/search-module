package search

import (
	"context"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	url := "http://localhost:9200"
	client := NewClient(url)

	if client == nil {
		t.Fatal("Expected client to be created")
	}
	if client.url != url {
		t.Errorf("Expected URL %s, got %s", url, client.url)
	}
}

func TestElasticsearchClient_Index(t *testing.T) {
	client := NewClient("http://localhost:9200")
	ctx := context.Background()

	t.Run("indexes document successfully", func(t *testing.T) {
		doc := map[string]interface{}{
			"title":   "Test Document",
			"content": "This is a test",
			"author":  "Test Author",
		}

		err := client.Index(ctx, "test-index", "doc-1", doc)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("indexes struct document", func(t *testing.T) {
		type Article struct {
			Title   string
			Content string
			Tags    []string
		}

		doc := Article{
			Title:   "Go Testing",
			Content: "Testing in Go is awesome",
			Tags:    []string{"go", "testing"},
		}

		err := client.Index(ctx, "articles", "article-1", doc)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})
}

func TestElasticsearchClient_Search(t *testing.T) {
	client := NewClient("http://localhost:9200")
	ctx := context.Background()

	t.Run("searches with match query", func(t *testing.T) {
		query := BuildMatchQuery("title", "test")
		results, err := client.Search(ctx, "test-index", query)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		// In stub implementation, returns nil
		if results != nil {
			t.Error("Expected nil results in stub implementation")
		}
	})

	t.Run("searches with term query", func(t *testing.T) {
		query := BuildTermQuery("status", "published")
		results, err := client.Search(ctx, "articles", query)

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if results != nil {
			t.Error("Expected nil results in stub implementation")
		}
	})

	t.Run("searches with custom query", func(t *testing.T) {
		query := map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": []map[string]interface{}{
						{"match": map[string]interface{}{"title": "test"}},
						{"term": map[string]interface{}{"status": "published"}},
					},
				},
			},
		}

		results, err := client.Search(ctx, "test-index", query)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if results != nil {
			t.Error("Expected nil results in stub implementation")
		}
	})
}

func TestElasticsearchClient_Delete(t *testing.T) {
	client := NewClient("http://localhost:9200")
	ctx := context.Background()

	t.Run("deletes document successfully", func(t *testing.T) {
		err := client.Delete(ctx, "test-index", "doc-1")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})

	t.Run("deletes document from different index", func(t *testing.T) {
		err := client.Delete(ctx, "articles", "article-1")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
	})
}

func TestBuildMatchQuery(t *testing.T) {
	t.Run("builds match query correctly", func(t *testing.T) {
		query := BuildMatchQuery("title", "test document")

		expected := map[string]interface{}{
			"query": map[string]interface{}{
				"match": map[string]interface{}{
					"title": "test document",
				},
			},
		}

		if !reflect.DeepEqual(query, expected) {
			t.Errorf("Expected query %v, got %v", expected, query)
		}
	})

	t.Run("builds match query for different field", func(t *testing.T) {
		query := BuildMatchQuery("content", "elasticsearch tutorial")

		if query["query"] == nil {
			t.Error("Expected query to have 'query' field")
		}

		queryMap := query["query"].(map[string]interface{})
		if queryMap["match"] == nil {
			t.Error("Expected query to have 'match' field")
		}

		matchMap := queryMap["match"].(map[string]interface{})
		if matchMap["content"] != "elasticsearch tutorial" {
			t.Errorf("Expected content value 'elasticsearch tutorial', got %v", matchMap["content"])
		}
	})
}

func TestBuildTermQuery(t *testing.T) {
	t.Run("builds term query correctly", func(t *testing.T) {
		query := BuildTermQuery("status", "published")

		expected := map[string]interface{}{
			"query": map[string]interface{}{
				"term": map[string]interface{}{
					"status": "published",
				},
			},
		}

		if !reflect.DeepEqual(query, expected) {
			t.Errorf("Expected query %v, got %v", expected, query)
		}
	})

	t.Run("builds term query for different field", func(t *testing.T) {
		query := BuildTermQuery("category", "tech")

		if query["query"] == nil {
			t.Error("Expected query to have 'query' field")
		}

		queryMap := query["query"].(map[string]interface{})
		if queryMap["term"] == nil {
			t.Error("Expected query to have 'term' field")
		}

		termMap := queryMap["term"].(map[string]interface{})
		if termMap["category"] != "tech" {
			t.Errorf("Expected category value 'tech', got %v", termMap["category"])
		}
	})
}

func TestSearchClientInterface(t *testing.T) {
	ctx := context.Background()

	t.Run("ElasticsearchClient implements SearchClient", func(t *testing.T) {
		var client SearchClient = NewClient("http://localhost:9200")

		// Test Index method
		err := client.Index(ctx, "test", "1", map[string]interface{}{"field": "value"})
		if err != nil {
			t.Fatalf("Expected no error from Index, got %v", err)
		}

		// Test Search method
		_, err = client.Search(ctx, "test", map[string]interface{}{})
		if err != nil {
			t.Fatalf("Expected no error from Search, got %v", err)
		}

		// Test Delete method
		err = client.Delete(ctx, "test", "1")
		if err != nil {
			t.Fatalf("Expected no error from Delete, got %v", err)
		}
	})
}
