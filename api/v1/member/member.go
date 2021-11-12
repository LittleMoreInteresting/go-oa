package member

import (
	"net/http"

	"go-oa/internal/biz"
	"go-oa/internal/server"

	"github.com/gin-gonic/gin"
)

func GetMemberByID(g *gin.Context) {
	mm := biz.NewMember()
	ms := server.NewMemberServer(mm)
	m := NewMember(ms)
	mi, _ := m.ms.GetMemberByID(1)
	g.JSON(http.StatusOK, gin.H{"id": mi.Id, "mobile": mi.Mobile})
}

type Member struct {
	ms MemberServer
}

func NewMember(ms MemberServer) *Member {
	return &Member{ms: ms}
}

type MemberServer interface {
	GetMemberByID(id int) (*server.MemberInfo, error)
}
