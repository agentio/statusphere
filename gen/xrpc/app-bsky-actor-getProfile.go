// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.actor.getProfile

import (
	"context"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyActorGetProfile_Description = "Get detailed profile view of an actor. Does not require auth, but contains relevant metadata with auth."

type AppBskyActorGetProfile_Output = AppBskyActorDefs_ProfileViewDetailed

// Get detailed profile view of an actor. Does not require auth, but contains relevant metadata with auth.
func AppBskyActorGetProfile(ctx context.Context, c slink.Client, actor string) (*AppBskyActorGetProfile_Output, error) {
	var output AppBskyActorGetProfile_Output
	params := map[string]any{
		"actor": actor,
	}
	if err := c.Do(ctx, slink.Query, "", "app.bsky.actor.getProfile", params, nil, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
