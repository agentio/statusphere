// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.actor.defs

import (
	"encoding/json"
	"fmt"

	"github.com/agentio/slink/pkg/slink"
)

const AppBskyActorDefs_KnownFollowers_Description = "The subject's followers whom you also follow"

// AppBskyActorDefs_KnownFollowers is a record with Lexicon type app.bsky.actor.defs#knownFollowers
type AppBskyActorDefs_KnownFollowers struct {
	LexiconTypeID string                               `json:"$type,omitempty"`
	Count         int64                                `json:"count"`
	Followers     []*AppBskyActorDefs_ProfileViewBasic `json:"followers,omitempty"`
}

const AppBskyActorDefs_ProfileAssociated_Description = ""

// AppBskyActorDefs_ProfileAssociated is a record with Lexicon type app.bsky.actor.defs#profileAssociated
type AppBskyActorDefs_ProfileAssociated struct {
	LexiconTypeID        string                                                  `json:"$type,omitempty"`
	ActivitySubscription *AppBskyActorDefs_ProfileAssociatedActivitySubscription `json:"activitySubscription,omitempty"`
	Chat                 *AppBskyActorDefs_ProfileAssociatedChat                 `json:"chat,omitempty"`
	Feedgens             *int64                                                  `json:"feedgens,omitempty"`
	Germ                 *AppBskyActorDefs_ProfileAssociatedGerm                 `json:"germ,omitempty"`
	Labeler              *bool                                                   `json:"labeler,omitempty"`
	Lists                *int64                                                  `json:"lists,omitempty"`
	StarterPacks         *int64                                                  `json:"starterPacks,omitempty"`
}

const AppBskyActorDefs_ProfileAssociatedActivitySubscription_Description = ""

// AppBskyActorDefs_ProfileAssociatedActivitySubscription is a record with Lexicon type app.bsky.actor.defs#profileAssociatedActivitySubscription
type AppBskyActorDefs_ProfileAssociatedActivitySubscription struct {
	LexiconTypeID      string `json:"$type,omitempty"`
	AllowSubscriptions string `json:"allowSubscriptions"`
}

const AppBskyActorDefs_ProfileAssociatedChat_Description = ""

// AppBskyActorDefs_ProfileAssociatedChat is a record with Lexicon type app.bsky.actor.defs#profileAssociatedChat
type AppBskyActorDefs_ProfileAssociatedChat struct {
	LexiconTypeID string `json:"$type,omitempty"`
	AllowIncoming string `json:"allowIncoming"`
}

const AppBskyActorDefs_ProfileAssociatedGerm_Description = ""

// AppBskyActorDefs_ProfileAssociatedGerm is a record with Lexicon type app.bsky.actor.defs#profileAssociatedGerm
type AppBskyActorDefs_ProfileAssociatedGerm struct {
	LexiconTypeID string `json:"$type,omitempty"`
	MessageMeUrl  string `json:"messageMeUrl"`
	ShowButtonTo  string `json:"showButtonTo"`
}

const AppBskyActorDefs_ProfileViewBasic_Description = ""

// AppBskyActorDefs_ProfileViewBasic is a record with Lexicon type app.bsky.actor.defs#profileViewBasic
type AppBskyActorDefs_ProfileViewBasic struct {
	LexiconTypeID string                              `json:"$type,omitempty"`
	Associated    *AppBskyActorDefs_ProfileAssociated `json:"associated,omitempty"`
	Avatar        *string                             `json:"avatar,omitempty"`
	CreatedAt     *string                             `json:"createdAt,omitempty"`
	Debug         *any                                `json:"debug,omitempty"`
	Did           string                              `json:"did"`
	DisplayName   *string                             `json:"displayName,omitempty"`
	Handle        string                              `json:"handle"`
	Labels        []*ComATProtoLabelDefs_Label        `json:"labels,omitempty"`
	Pronouns      *string                             `json:"pronouns,omitempty"`
	Status        *AppBskyActorDefs_StatusView        `json:"status,omitempty"`
	Verification  *AppBskyActorDefs_VerificationState `json:"verification,omitempty"`
	Viewer        *AppBskyActorDefs_ViewerState       `json:"viewer,omitempty"`
}

const AppBskyActorDefs_ProfileViewDetailed_Description = ""

// AppBskyActorDefs_ProfileViewDetailed is a record with Lexicon type app.bsky.actor.defs#profileViewDetailed
type AppBskyActorDefs_ProfileViewDetailed struct {
	LexiconTypeID        string                                 `json:"$type,omitempty"`
	Associated           *AppBskyActorDefs_ProfileAssociated    `json:"associated,omitempty"`
	Avatar               *string                                `json:"avatar,omitempty"`
	Banner               *string                                `json:"banner,omitempty"`
	CreatedAt            *string                                `json:"createdAt,omitempty"`
	Debug                *any                                   `json:"debug,omitempty"`
	Description          *string                                `json:"description,omitempty"`
	Did                  string                                 `json:"did"`
	DisplayName          *string                                `json:"displayName,omitempty"`
	FollowersCount       *int64                                 `json:"followersCount,omitempty"`
	FollowsCount         *int64                                 `json:"followsCount,omitempty"`
	Handle               string                                 `json:"handle"`
	IndexedAt            *string                                `json:"indexedAt,omitempty"`
	JoinedViaStarterPack *AppBskyGraphDefs_StarterPackViewBasic `json:"joinedViaStarterPack,omitempty"`
	Labels               []*ComATProtoLabelDefs_Label           `json:"labels,omitempty"`
	PinnedPost           *ComATProtoRepoStrongRef               `json:"pinnedPost,omitempty"`
	PostsCount           *int64                                 `json:"postsCount,omitempty"`
	Pronouns             *string                                `json:"pronouns,omitempty"`
	Status               *AppBskyActorDefs_StatusView           `json:"status,omitempty"`
	Verification         *AppBskyActorDefs_VerificationState    `json:"verification,omitempty"`
	Viewer               *AppBskyActorDefs_ViewerState          `json:"viewer,omitempty"`
	Website              *string                                `json:"website,omitempty"`
}

const AppBskyActorDefs_StatusView_Description = ""

// AppBskyActorDefs_StatusView is a record with Lexicon type app.bsky.actor.defs#statusView
type AppBskyActorDefs_StatusView struct {
	LexiconTypeID string                             `json:"$type,omitempty"`
	Cid           *string                            `json:"cid,omitempty"`
	Embed         *AppBskyActorDefs_StatusView_Embed `json:"embed,omitempty"`
	ExpiresAt     *string                            `json:"expiresAt,omitempty"`
	IsActive      *bool                              `json:"isActive,omitempty"`
	IsDisabled    *bool                              `json:"isDisabled,omitempty"`
	Record        any                                `json:"record"`
	Status        string                             `json:"status"`
	Uri           *string                            `json:"uri,omitempty"`
}

// AppBskyActorDefs_StatusView_Embed is a union with these possible values:
// - AppBskyEmbedExternal_View (app.bsky.embed.external#view)
type AppBskyActorDefs_StatusView_Embed struct {
	Wrapper AppBskyActorDefs_StatusView_Embed_Wrapper
}

// Value wrappers must conform to AppBskyActorDefs_StatusView_Embed_Wrapper
type AppBskyActorDefs_StatusView_Embed_Wrapper interface {
	isAppBskyActorDefs_StatusView_Embed()
}

// AppBskyActorDefs_StatusView_Embed__AppBskyEmbedExternal_View wraps values of type *AppBskyEmbedExternal_View
type AppBskyActorDefs_StatusView_Embed__AppBskyEmbedExternal_View struct {
	Value *AppBskyEmbedExternal_View
}

func (t AppBskyActorDefs_StatusView_Embed__AppBskyEmbedExternal_View) isAppBskyActorDefs_StatusView_Embed() {
}

func (u AppBskyActorDefs_StatusView_Embed) MarshalJSON() ([]byte, error) {
	switch v := u.Wrapper.(type) {
	case AppBskyActorDefs_StatusView_Embed__AppBskyEmbedExternal_View:
		return slink.MarshalWithLexiconType("app.bsky.embed.external#view", v.Value)
	default:
		return nil, fmt.Errorf("unsupported type %T", v)
	}
}
func (u *AppBskyActorDefs_StatusView_Embed) UnmarshalJSON(data []byte) error {
	switch slink.LexiconTypeFromJSONBytes(data) {
	case "app.bsky.embed.external#view":
		var v AppBskyEmbedExternal_View
		if err := json.Unmarshal(data, &v); err != nil {
			return err
		}
		u.Wrapper = AppBskyActorDefs_StatusView_Embed__AppBskyEmbedExternal_View{Value: &v}
		return nil
	default:
		return nil
	}
}

const AppBskyActorDefs_VerificationState_Description = "Represents the verification information about the user this object is attached to."

// AppBskyActorDefs_VerificationState is a record with Lexicon type app.bsky.actor.defs#verificationState
type AppBskyActorDefs_VerificationState struct {
	LexiconTypeID         string                               `json:"$type,omitempty"`
	TrustedVerifierStatus string                               `json:"trustedVerifierStatus"`
	Verifications         []*AppBskyActorDefs_VerificationView `json:"verifications,omitempty"`
	VerifiedStatus        string                               `json:"verifiedStatus"`
}

const AppBskyActorDefs_VerificationView_Description = "An individual verification for an associated subject."

// AppBskyActorDefs_VerificationView is a record with Lexicon type app.bsky.actor.defs#verificationView
type AppBskyActorDefs_VerificationView struct {
	LexiconTypeID string `json:"$type,omitempty"`
	CreatedAt     string `json:"createdAt"`
	IsValid       bool   `json:"isValid"`
	Issuer        string `json:"issuer"`
	Uri           string `json:"uri"`
}

const AppBskyActorDefs_ViewerState_Description = "Metadata about the requesting account's relationship with the subject account. Only has meaningful content for authed requests."

// AppBskyActorDefs_ViewerState is a record with Lexicon type app.bsky.actor.defs#viewerState
type AppBskyActorDefs_ViewerState struct {
	LexiconTypeID        string                                        `json:"$type,omitempty"`
	ActivitySubscription *AppBskyNotificationDefs_ActivitySubscription `json:"activitySubscription,omitempty"`
	BlockedBy            *bool                                         `json:"blockedBy,omitempty"`
	Blocking             *string                                       `json:"blocking,omitempty"`
	BlockingByList       *AppBskyGraphDefs_ListViewBasic               `json:"blockingByList,omitempty"`
	FollowedBy           *string                                       `json:"followedBy,omitempty"`
	Following            *string                                       `json:"following,omitempty"`
	KnownFollowers       *AppBskyActorDefs_KnownFollowers              `json:"knownFollowers,omitempty"`
	Muted                *bool                                         `json:"muted,omitempty"`
	MutedByList          *AppBskyGraphDefs_ListViewBasic               `json:"mutedByList,omitempty"`
}
