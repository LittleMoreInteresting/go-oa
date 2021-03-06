//+build wireinject

//The build tag makes sure the stub is not built in the final build.
package member

import (
	"github.com/google/wire"
	"go-oa/internal/biz"
	"go-oa/internal/service"
)

var BizNewMember = wire.NewSet(biz.NewMember, wire.Bind(new(service.MemberModel), new(*biz.Member)))
var NewMemberServer = wire.NewSet(service.NewMemberServer, wire.Bind(new(MemberServer), new(*service.MemberServer)))

// InitializeEvent 声明injector的函数签名
func InitializeMember() *Member {
	wire.Build(NewMember, NewMemberServer, BizNewMember)
	return &Member{}
}
