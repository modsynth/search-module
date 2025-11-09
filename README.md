# Search Module

> Elasticsearch search integration

Part of the [Modsynth](https://github.com/modsynth) ecosystem.

## Features

- Full-text search with Elasticsearch
- Document indexing
- Search query builders
- Aggregations support

## Installation

```bash
go get github.com/modsynth/search-module
```

## Quick Start

```go
package main

import (
    "context"
    "github.com/modsynth/search-module"
)

func main() {
    client := search.NewClient("http://localhost:9200")

    // Index a document
    doc := map[string]interface{}{
        "title": "Hello World",
        "content": "This is a test document",
    }
    client.Index(context.Background(), "posts", "1", doc)

    // Search
    query := search.BuildMatchQuery("content", "test")
    results, _ := client.Search(context.Background(), "posts", query)

    // Delete
    client.Delete(context.Background(), "posts", "1")
}
```

## Query Builders

- `BuildMatchQuery` - Full-text match query
- `BuildTermQuery` - Exact term query

## Version

Current version: `v0.1.0`

## License

MIT
