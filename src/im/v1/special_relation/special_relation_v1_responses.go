package special_relation

// Special Relation V1 Response Structures

// ListSpecialRelationResponseV1 ListSpecialRelation响应
type ListSpecialRelationResponseV1 struct {
	Mutelist  []string `json:"mutelist,omitempty"`
	Blacklist []string `json:"blacklist,omitempty"`
}

// SetSpecialRelationResponseV1 SetSpecialRelation响应
type SetSpecialRelationResponseV1 struct {
}
