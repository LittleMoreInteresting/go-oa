package server

type MemberServer struct {
}
type MemberInfo struct {
	Id     int
	Mobile string
}

func (ms *MemberServer) GetMemberByID(id int) (*MemberInfo, error) {
	return &MemberInfo{Id: 2, Mobile: "121212"}, nil
}
