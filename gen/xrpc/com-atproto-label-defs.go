// Code generated ... DO NOT EDIT.

// Package xrpc is generated from Lexicon source files by slink.
// Code produced by slink and slink itself are released under the AGPL.
// Get slink at https://github.com/agentio/slink.
package xrpc // com.atproto.label.defs

const ComATProtoLabelDefs_Label_Description = "Metadata tag on an atproto resource (eg, repo or record)."

// ComATProtoLabelDefs_Label is a record with Lexicon type com.atproto.label.defs#label
type ComATProtoLabelDefs_Label struct {
	LexiconTypeID string  `json:"$type,omitempty"`
	Cid           *string `json:"cid,omitempty"`
	Cts           string  `json:"cts"`
	Exp           *string `json:"exp,omitempty"`
	Neg           *bool   `json:"neg,omitempty"`
	Sig           *[]byte `json:"sig,omitempty"`
	Src           string  `json:"src"`
	Uri           string  `json:"uri"`
	Val           string  `json:"val"`
	Ver           *int64  `json:"ver,omitempty"`
}
