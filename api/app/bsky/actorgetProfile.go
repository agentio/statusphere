package bsky

// schema: app.bsky.actor.getProfile

import (
	"context"

	"github.com/agentio/statusphere-io/pkg/xrpc"
)

// ActorDefs_ProfileViewDetailed is a "profileViewDetailed" in the app.bsky.actor.defs schema.
type ActorDefs_ProfileViewDetailed struct {
	Avatar         *string `json:"avatar,omitempty" cborgen:"avatar,omitempty"`
	Banner         *string `json:"banner,omitempty" cborgen:"banner,omitempty"`
	CreatedAt      *string `json:"createdAt,omitempty" cborgen:"createdAt,omitempty"`
	Description    *string `json:"description,omitempty" cborgen:"description,omitempty"`
	Did            string  `json:"did" cborgen:"did"`
	DisplayName    *string `json:"displayName,omitempty" cborgen:"displayName,omitempty"`
	FollowersCount *int64  `json:"followersCount,omitempty" cborgen:"followersCount,omitempty"`
	FollowsCount   *int64  `json:"followsCount,omitempty" cborgen:"followsCount,omitempty"`
	Handle         string  `json:"handle" cborgen:"handle"`
	IndexedAt      *string `json:"indexedAt,omitempty" cborgen:"indexedAt,omitempty"`
	PostsCount     *int64  `json:"postsCount,omitempty" cborgen:"postsCount,omitempty"`
}

// ActorGetProfile calls the XRPC method "app.bsky.actor.getProfile".
//
// actor: Handle or DID of account to fetch profile of.
func ActorGetProfile(ctx context.Context, c xrpc.Client, actor string) (*ActorDefs_ProfileViewDetailed, error) {
	var out ActorDefs_ProfileViewDetailed

	params := map[string]interface{}{
		"actor": actor,
	}
	if err := c.Do(ctx, xrpc.Query, "", "app.bsky.actor.getProfile", params, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
