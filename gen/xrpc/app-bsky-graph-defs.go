// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Get slink at https://github.com/agentio/slink.
package xrpc // app.bsky.graph.defs

type AppBskyGraphDefs_ListPurpose string

const AppBskyGraphDefs_ListViewBasic_Description = ""

// AppBskyGraphDefs_ListViewBasic is a record with Lexicon type app.bsky.graph.defs#listViewBasic
type AppBskyGraphDefs_ListViewBasic struct {
	LexiconTypeID string                            `json:"$type,omitempty"`
	Avatar        *string                           `json:"avatar,omitempty"`
	Cid           string                            `json:"cid"`
	IndexedAt     *string                           `json:"indexedAt,omitempty"`
	Labels        []*ComATProtoLabelDefs_Label      `json:"labels,omitempty"`
	ListItemCount *int64                            `json:"listItemCount,omitempty"`
	Name          string                            `json:"name"`
	Purpose       *AppBskyGraphDefs_ListPurpose     `json:"purpose,omitempty"`
	Uri           string                            `json:"uri"`
	Viewer        *AppBskyGraphDefs_ListViewerState `json:"viewer,omitempty"`
}

const AppBskyGraphDefs_ListViewerState_Description = ""

// AppBskyGraphDefs_ListViewerState is a record with Lexicon type app.bsky.graph.defs#listViewerState
type AppBskyGraphDefs_ListViewerState struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Blocked       *string `json:"blocked,omitempty"`
	Muted         *bool   `json:"muted,omitempty"`
}

const AppBskyGraphDefs_StarterPackViewBasic_Description = ""

// AppBskyGraphDefs_StarterPackViewBasic is a record with Lexicon type app.bsky.graph.defs#starterPackViewBasic
type AppBskyGraphDefs_StarterPackViewBasic struct {
	LexiconTypeID      string                             `json:"$type,omitempty"`
	Cid                string                             `json:"cid"`
	Creator            *AppBskyActorDefs_ProfileViewBasic `json:"creator,omitempty"`
	IndexedAt          string                             `json:"indexedAt"`
	JoinedAllTimeCount *int64                             `json:"joinedAllTimeCount,omitempty"`
	JoinedWeekCount    *int64                             `json:"joinedWeekCount,omitempty"`
	Labels             []*ComATProtoLabelDefs_Label       `json:"labels,omitempty"`
	ListItemCount      *int64                             `json:"listItemCount,omitempty"`
	Record             any                                `json:"record"`
	Uri                string                             `json:"uri"`
}
