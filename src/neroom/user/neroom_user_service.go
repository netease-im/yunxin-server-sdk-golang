package user

import (
	. "github.com/netease-im/yunxin-server-sdk-golang/src/neroom/result"
)

// NeroomUserService Neroom用户服务接口
type NeroomUserService interface {
	// CreateAccount 创建Neroom账号
	CreateAccount(req *CreateNeroomAccountRequest) (*NeroomResult[*CreateNeroomAccountResponse], error)
}
