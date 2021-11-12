package service

import "go-oa/internal/biz"

type MemberServer struct {
	mm MemberModel
}
type MemberInfo struct {
	Id     int
	Mobile string
}

func NewMemberServer(mm MemberModel) *MemberServer {
	return &MemberServer{mm: mm}
}

func (ms *MemberServer) GetMemberByID(id int) (*MemberInfo, error) {
	mm := ms.mm.GetMember(id)
	return &MemberInfo{Id: mm.Id, Mobile: mm.Mobile}, nil
}

type MemberModel interface {
	GetMember(id int) biz.Member
}
