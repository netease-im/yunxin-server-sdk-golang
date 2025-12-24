package special_relation

import (
	"github.com/netease-im/yunxin-server-sdk-golang/src/core"
)

// Special RelationV1Service special_relation服务接口
type SpecialRelationV1Service interface {
	// SetSpecialRelation SetSpecialRelation
	SetSpecialRelation(req *SetSpecialRelationRequestV1) (*core.Result[*SetSpecialRelationResponseV1], error)

	// ListSpecialRelation ListSpecialRelation
	ListSpecialRelation(req *ListSpecialRelationRequestV1) (*core.Result[*ListSpecialRelationResponseV1], error)
}

// SpecialRelationV1ServiceImpl special_relation服务实现
type SpecialRelationV1ServiceImpl struct {
	httpClient core.YunxinApiHttpClient
}

// NewSpecialRelationV1Service 创建special_relation服务实例
func NewSpecialRelationV1Service(httpClient core.YunxinApiHttpClient) SpecialRelationV1Service {
	return &SpecialRelationV1ServiceImpl{
		httpClient: httpClient,
	}
}
