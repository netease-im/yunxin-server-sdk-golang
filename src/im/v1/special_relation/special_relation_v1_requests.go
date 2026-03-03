package special_relation

// Special Relation V1 Request Structures

// ListSpecialRelationRequestV1 ListSpecialRelation请求
type ListSpecialRelationRequestV1 struct {
	Accid string `json:"accid,omitempty"`
}

// SetSpecialRelationRequestV1 SetSpecialRelation请求
type SetSpecialRelationRequestV1 struct {
	Accid        string `json:"accid,omitempty"`
	TargetAcc    string `json:"targetAcc,omitempty"`
	RelationType *int   `json:"relationType,omitempty"` // 使用指针，0是有效值（静音）
	Value        *int   `json:"value,omitempty"`        // 使用指针，0是有效值（移除）
}
