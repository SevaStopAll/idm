package tests

import "idm/inner/employee"

type EmployeeFixture struct {
	employees *employee.EmployeeRepository
}

func NewEmployeeFixture(employees *employee.EmployeeRepository) *EmployeeFixture {
	return &EmployeeFixture{employees}
}

func (f *EmployeeFixture) CreateEmployeeTable() {
	f.employees.ExecuteQuery(
		"create table employee(" +
			"    id         bigint generated always as identity primary key," +
			"    name text not null," +
			"created_at timestamp with time zone default now() not null," +
			"    updated_at timestamp with time zone default now() not null)")
}

func (f *EmployeeFixture) Employee(name string) int64 {
	var entity = employee.EmployeeEntity{
		Name: name,
	}
	var newId, err = f.employees.Save(&entity)
	if err != nil {
		panic(err)
	}
	return newId
}

func (f *EmployeeFixture) Employees(names ...string) []int64 {
	var ids = make([]int64, len(names))
	for i := 0; i < len(names); i++ {
		var entity = employee.EmployeeEntity{
			Name: names[i],
		}
		var newId, err = f.employees.Save(&entity)
		if err != nil {
			panic(err)
		}
		ids[i] = newId
	}
	return ids
}
