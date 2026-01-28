// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.repo.createRecord

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const ComATProtoRepoCreateRecord_Description = "Create a single new repository record. Requires auth, implemented by PDS."

type ComATProtoRepoCreateRecord_Input struct {
	Collection string  `json:"collection"`
	Record     any     `json:"record"`
	Repo       string  `json:"repo"`
	Rkey       *string `json:"rkey,omitempty"`
	SwapCommit *string `json:"swapCommit,omitempty"`
	Validate   *bool   `json:"validate,omitempty"`
}

type ComATProtoRepoCreateRecord_Output struct {
	Cid              string                         `json:"cid"`
	Commit           *ComATProtoRepoDefs_CommitMeta `json:"commit,omitempty"`
	Uri              string                         `json:"uri"`
	ValidationStatus *string                        `json:"validationStatus,omitempty"`
}

// Create a single new repository record. Requires auth, implemented by PDS.
func ComATProtoRepoCreateRecord(ctx context.Context, c slink.Client, input *ComATProtoRepoCreateRecord_Input) (*ComATProtoRepoCreateRecord_Output, error) {
	var output ComATProtoRepoCreateRecord_Output
	if err := c.Do(ctx, slink.Procedure, "application/json", "com.atproto.repo.createRecord", nil, input, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
