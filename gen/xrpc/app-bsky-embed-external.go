// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.embed.external

const AppBskyEmbedExternal_View_Description = ""

// AppBskyEmbedExternal_View is a record with Lexicon type app.bsky.embed.external#view
type AppBskyEmbedExternal_View struct {
	LexiconTypeID string                             `json:"$type,omitempty"`
	External      *AppBskyEmbedExternal_ViewExternal `json:"external,omitempty"`
}

const AppBskyEmbedExternal_ViewExternal_Description = ""

// AppBskyEmbedExternal_ViewExternal is a record with Lexicon type app.bsky.embed.external#viewExternal
type AppBskyEmbedExternal_ViewExternal struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Description   string  `json:"description"`
	Thumb         *string `json:"thumb,omitempty"`
	Title         string  `json:"title"`
	Uri           string  `json:"uri"`
}
