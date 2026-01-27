// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.notification.defs

const AppBskyNotificationDefs_ActivitySubscription_Description = ""

// AppBskyNotificationDefs_ActivitySubscription is a record with Lexicon type app.bsky.notification.defs#activitySubscription
type AppBskyNotificationDefs_ActivitySubscription struct {
	LexiconTypeID string `json:"$type,omitempty"`
	Post          bool   `json:"post"`
	Reply         bool   `json:"reply"`
}
