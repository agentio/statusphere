// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.repo.listRecords

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const ComATProtoRepoListRecords_Description = "List a range of records in a repository, matching a specific collection. Does not require auth."

type ComATProtoRepoListRecords_Output struct {
	Cursor  *string                             `json:"cursor,omitempty"`
	Records []*ComATProtoRepoListRecords_Record `json:"records,omitempty"`
}

// List a range of records in a repository, matching a specific collection. Does not require auth.
func ComATProtoRepoListRecords(ctx context.Context, c slink.Client, collection string, cursor string, limit int64, repo string, reverse bool) (*ComATProtoRepoListRecords_Output, error) {
	var output ComATProtoRepoListRecords_Output
	params := map[string]any{
		"reverse":    reverse,
		"repo":       repo,
		"collection": collection,
		"limit":      limit,
		"cursor":     cursor,
	}
	if err := c.Do(ctx, slink.Query, "", "com.atproto.repo.listRecords", params, nil, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

const ComATProtoRepoListRecords_Record_Description = ""

// ComATProtoRepoListRecords_Record is a record with Lexicon type com.atproto.repo.listRecords#record
type ComATProtoRepoListRecords_Record struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Cid           string `json:"cid"`
	Uri           string `json:"uri"`
	Value         any    `json:"value"`
}
