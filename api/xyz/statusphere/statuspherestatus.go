package statusphere

// schema: xyz.statusphere.status

// RECORDTYPE: Status
type Status struct {
	LexiconTypeID string `json:"$type,const=xyz.statusphere.status" cborgen:"$type,const=xyz.statusphere.status"`
	CreatedAt     string `json:"createdAt" cborgen:"createdAt"`
	Status        string `json:"status" cborgen:"status"`
}
