package biz

type Member struct {
	Id     int
	Name   string
	Email  string
	Mobile string
}

func NewMember() *Member {
	return &Member{}
}

func (m *Member) GetMember(id int) Member {
	return Member{
		Id:     id,
		Name:   "MockName",
		Email:  "12@12.com",
		Mobile: "1111111111",
	}
}
