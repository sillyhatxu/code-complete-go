package build

import (
	"fmt"
	"time"
)

type user struct {
	id         string
	name       string
	age        *int
	status     bool
	createTime time.Time
}

func (u *user) String() string {
	if u.age == nil {
		return fmt.Sprintf(`{id:"%s",name:"%s",age:nil,status:%v,createTime:"%s"}`,
			u.id,
			u.name,
			u.status,
			u.createTime.Format(time.RFC3339),
		)
	}
	return fmt.Sprintf(`{id:"%s",name:"%s",age:%v,status:%v,createTime:"%s"}`,
		u.id,
		u.name,
		*u.age,
		u.status,
		u.createTime.Format(time.RFC3339),
	)
}

type userBuilder struct {
	id         string
	name       string
	age        *int
	status     bool
	createTime time.Time
}

func New() *userBuilder {
	return &userBuilder{}
}

func (ub *userBuilder) Id(id string) *userBuilder {
	ub.id = id
	return ub
}
func (ub *userBuilder) Name(name string) *userBuilder {
	ub.name = name
	return ub
}
func (ub *userBuilder) Age(age int) *userBuilder {
	ub.age = &age
	return ub
}
func (ub *userBuilder) Status(status bool) *userBuilder {
	ub.status = status
	return ub
}
func (ub *userBuilder) CreateTime(createTime time.Time) *userBuilder {
	ub.createTime = createTime
	return ub
}
func (ub *userBuilder) Build() *user {
	return &user{
		id:         ub.id,
		name:       ub.name,
		age:        ub.age,
		status:     ub.status,
		createTime: ub.createTime,
	}
}
