package tests

import (
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"idm/inner/employee"
	"idm/inner/role"
	"testing"
)

func Test_Employees(t *testing.T) {
	a := assert.New(t)
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())

	var clearEmployeeDb = func() {
		employeeRepository.ExecuteQuery("DELETE FROM employee")
	}
	defer func() {
		if r := recover(); r != nil {
			clearEmployeeDb()
		}
	}()
	var fixture = NewEmployeeFixture(employeeRepository)
	fixture.CreateEmployeeTable()

	t.Run("save an employee", func(t *testing.T) {
		res, err := employeeRepository.Save(&employee.EmployeeEntity{Name: "testNameSaved"})
		if err != nil {
			t.Fatal(err)
		}
		employeeFound, err := employeeRepository.FindById(res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, employeeFound.Name, "testNameSaved")
		clearEmployeeDb()
	})

	t.Run("find an employee by id", func(t *testing.T) {
		var newEmployeeId = fixture.Employee("Test name2")
		got, err := employeeRepository.FindById(newEmployeeId)
		a.Nil(err)
		a.NotEmpty(got)
		a.NotEmpty(got.Id)
		a.NotEmpty(got.CreatedAt)
		a.NotEmpty(got.UpdatedAt)
		a.Equal("Test name2", got.Name)
		clearEmployeeDb()
	})

	t.Run("find an employees by their ids", func(t *testing.T) {
		var employeesId = fixture.Employees("Name1", "Name2", "Name3", "Name4")
		got, err := employeeRepository.FindByIds(employeesId)
		a.Nil(err)
		a.NotEmpty(got)
		a.Equal("Name1", got[0].Name)
		a.Equal(4, len(employeesId))
		clearEmployeeDb()
	})

	t.Run("find all employees", func(t *testing.T) {
		var employeesId = fixture.Employees("Name1", "Name2", "Name3", "Name4", "Name5")
		got, err := employeeRepository.FindAll()
		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(5, len(employeesId))
		clearEmployeeDb()
	})

	t.Run("delete an employee by its id", func(t *testing.T) {
		var employeesId = fixture.Employees("Name1", "Name2", "Name3", "Name4", "Name5")
		err1 := employeeRepository.DeleteById(employeesId[0])
		a.Nil(err1)
		if err1 != nil {
			return
		}
		got, err := employeeRepository.FindAll()
		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(4, len(got))
		clearEmployeeDb()
	})

	t.Run("delete employee by their ids", func(t *testing.T) {
		employeeRepository1 := employee.NewEmployeeRepository(database.ConnectDb())
		id1, err2 := employeeRepository1.Save(&employee.EmployeeEntity{Name: "testName1"})
		id2, err3 := employeeRepository1.Save(&employee.EmployeeEntity{Name: "testName2"})
		_, err4 := employeeRepository1.Save(&employee.EmployeeEntity{Name: "testName3"})
		_, err5 := employeeRepository1.Save(&employee.EmployeeEntity{Name: "testName4"})
		if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
			ids := make([]int64, 2)
			ids[0], ids[1] = id1, id2
			err := employeeRepository1.DeleteByIds(ids)
			if err != nil {
				t.Fatal(err)
			}
			res, err := employeeRepository1.FindAll()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, 2, len(res))
		}
		clearEmployeeDb()
	})

}

func Test_Roles(t *testing.T) {
	a := assert.New(t)
	roleRepository := role.NewRoleRepository(database.ConnectDb())

	var clearRoleDb = func() {
		roleRepository.ExecuteQuery("DELETE FROM role")
	}
	defer func() {
		if r := recover(); r != nil {
			clearRoleDb()
		}
	}()
	var fixture = NewRoleFixture(roleRepository)
	fixture.CreateRoleTable()

	t.Run("save a role", func(t *testing.T) {
		res, err := roleRepository.Save(&role.RoleEntity{Name: "testNameSaved"})
		if err != nil {
			t.Fatal(err)
		}
		employeeFound, err := roleRepository.FindById(res)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, employeeFound.Name, "testNameSaved")
		clearRoleDb()
	})

	t.Run("find a role by id", func(t *testing.T) {
		var newEmployeeId = fixture.Role("Test name2")
		got, err := roleRepository.FindById(newEmployeeId)
		a.Nil(err)
		a.NotEmpty(got)
		a.NotEmpty(got.Id)
		a.NotEmpty(got.CreatedAt)
		a.NotEmpty(got.UpdatedAt)
		a.Equal("Test name2", got.Name)
		clearRoleDb()
	})

	t.Run("find roles by their ids", func(t *testing.T) {
		var employeesId = fixture.Roles("Name1", "Name2", "Name3", "Name4")
		got, err := roleRepository.FindByIds(employeesId)
		a.Nil(err)
		a.NotEmpty(got)
		a.Equal("Name1", got[0].Name)
		a.Equal(4, len(employeesId))
		clearRoleDb()
	})

	t.Run("find all roles", func(t *testing.T) {
		var employeesId = fixture.Roles("Name1", "Name2", "Name3", "Name4", "Name5")
		got, err := roleRepository.FindAll()
		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(5, len(employeesId))
		clearRoleDb()
	})

	t.Run("delete a role by its id", func(t *testing.T) {
		var employeesId = fixture.Roles("Name1", "Name2", "Name3", "Name4", "Name5")
		err1 := roleRepository.DeleteById(employeesId[0])
		a.Nil(err1)
		if err1 != nil {
			return
		}
		got, err := roleRepository.FindAll()
		a.Nil(err)
		a.NotEmpty(got)
		a.Equal(4, len(got))
		clearRoleDb()
	})

	t.Run("delete roles by their ids", func(t *testing.T) {
		roleRepository := role.NewRoleRepository(database.ConnectDb())
		id1, err2 := roleRepository.Save(&role.RoleEntity{Name: "testRole1"})
		id2, err3 := roleRepository.Save(&role.RoleEntity{Name: "testRole2"})
		_, err4 := roleRepository.Save(&role.RoleEntity{Name: "testRole3"})
		_, err5 := roleRepository.Save(&role.RoleEntity{Name: "testRole4"})
		if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
			res, err := roleRepository.FindAll()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, 4, len(res))
			ids := make([]int64, 2)
			ids[0], ids[1] = id1, id2
			err = roleRepository.DeleteByIds(ids)
			if err != nil {
				t.Fatal(err)
			}
			res, err = roleRepository.FindAll()
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, 2, len(res))
		}
		roleRepository.ExecuteQuery("DELETE FROM role")
	})

}
