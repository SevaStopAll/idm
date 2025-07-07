package tests

import (
	"idm/inner/role"
)

type RoleFixture struct {
	roles *role.RoleRepository
}

func NewRoleFixture(roles *role.RoleRepository) *RoleFixture {
	return &RoleFixture{roles}
}

func (f *RoleFixture) CreateRoleTable() {
	f.roles.ExecuteQuery(
		"create table role(" +
			"    id         bigint generated always as identity primary key," +
			"    name text not null," +
			"created_at timestamp with time zone default now() not null," +
			"    updated_at timestamp with time zone default now() not null)")
}

func (f *RoleFixture) Role(name string) int64 {
	var entity = role.RoleEntity{
		Name: name,
	}
	var newId, err = f.roles.Save(&entity)
	if err != nil {
		panic(err)
	}
	return newId
}

func (f *RoleFixture) Roles(names ...string) []int64 {
	var ids = make([]int64, len(names))
	for i := 0; i < len(names); i++ {
		var entity = role.RoleEntity{
			Name: names[i],
		}
		var newId, err = f.roles.Save(&entity)
		if err != nil {
			panic(err)
		}
		ids[i] = newId
	}
	return ids
}
