package tests

import (
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"idm/inner/employee"
	"idm/inner/role"
	"testing"
)

func Test_Save(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	res, err := employeeRepository.Save(employee.EmployeeEntity{Name: "testNameSaved"})
	if err != nil {
		t.Fatal(err)
	}
	employeeFound, err := employeeRepository.FindById(res)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, employeeFound.Name, "testNameSaved")
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindById(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	id, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName"})
	if err2 != nil {
		t.Fatal(err2)
	}
	res, err := employeeRepository.FindById(id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "testName", res.Name)
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindByIds(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	id1, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	id2, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	id3, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	id4, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		ids := []int64{id1, id2, id3, id4}
		res, err := employeeRepository.FindByIds(ids)
		if err != nil {
			t.Fatal(err2)
		}
		assert.Equal(t, 4, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindAll(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("DELETE FROM employee")
	_, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	_, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := employeeRepository.FindAll()
		if err != nil {
			t.Fatal(err2)
		}
		assert.Equal(t, 4, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_DeleteById(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())

	id1, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	_, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := employeeRepository.FindAll()
		if err != nil {
			t.Fatal(err2)
		}
		assert.Equal(t, 4, len(res))
		err = employeeRepository.DeleteById(id1)
		if err != nil {
			t.Fatal(err)
		}
		res, err = employeeRepository.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 3, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_DeleteByIds(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	id1, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	id2, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		ids := make([]int64, 2)
		ids[0], ids[1] = id1, id2
		err := employeeRepository.DeleteByIds(ids)
		if err != nil {
			t.Fatal(err)
		}
		res, err := employeeRepository.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 2, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

// Role
func Test_SaveRole(t *testing.T) {
	roleRepository := role.NewRoleRepository(database.ConnectDb())
	res, err := roleRepository.Save(role.RoleEntity{Name: "testRoleForSaving"})
	if err != nil {
		t.Fatal(err)
	}
	foundEmployee, err := roleRepository.FindById(res)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, foundEmployee.Name, "testRoleForSaving")
	roleRepository.ExecuteQuery("DELETE FROM role")
}

func Test_FindRoleById(t *testing.T) {
	roleRepository := role.NewRoleRepository(database.ConnectDb())
	id, err2 := roleRepository.Save(role.RoleEntity{Name: "testRole"})
	if err2 != nil {
		t.Fatal(err2)
	}
	res, err := roleRepository.FindById(id)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "testRole", res.Name)
	roleRepository.ExecuteQuery("DELETE FROM role")
}

func Test_FindRolesByIds(t *testing.T) {
	roleRepository := role.NewRoleRepository(database.ConnectDb())
	firstId, err2 := roleRepository.Save(role.RoleEntity{Name: "testRole1"})
	secondId, err3 := roleRepository.Save(role.RoleEntity{Name: "testRole2"})
	thirdId, err4 := roleRepository.Save(role.RoleEntity{Name: "testRole3"})
	fourthId, err5 := roleRepository.Save(role.RoleEntity{Name: "testRole4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		ids := []int64{firstId, secondId, thirdId, fourthId}
		res, err := roleRepository.FindByIds(ids)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 4, len(res))
	}
	roleRepository.ExecuteQuery("DELETE FROM role")
}

func Test_FindAllRoles(t *testing.T) {
	roleRepository := role.NewRoleRepository(database.ConnectDb())
	_, err2 := roleRepository.Save(role.RoleEntity{Name: "testRole1"})
	_, err3 := roleRepository.Save(role.RoleEntity{Name: "testRole2"})
	_, err4 := roleRepository.Save(role.RoleEntity{Name: "testRole3"})
	_, err5 := roleRepository.Save(role.RoleEntity{Name: "testRole4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := roleRepository.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 4, len(res))
	}
	roleRepository.ExecuteQuery("DELETE FROM role")
}

func Test_DeleteRoleById(t *testing.T) {
	roleRepository := role.NewRoleRepository(database.ConnectDb())
	id1, err2 := roleRepository.Save(role.RoleEntity{Name: "testRole1"})
	_, err3 := roleRepository.Save(role.RoleEntity{Name: "testRole2"})
	_, err4 := roleRepository.Save(role.RoleEntity{Name: "testRole3"})
	_, err5 := roleRepository.Save(role.RoleEntity{Name: "testRole4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := roleRepository.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 4, len(res))
		err = roleRepository.DeleteById(id1)
		if err != nil {
			t.Fatal(err)
		}
		res, err = roleRepository.FindAll()
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 3, len(res))
	}
	roleRepository.ExecuteQuery("DELETE FROM role")
}

func Test_DeleteRolesByIds(t *testing.T) {
	roleRepository := role.NewRoleRepository(database.ConnectDb())
	id1, err2 := roleRepository.Save(role.RoleEntity{Name: "testRole1"})
	id2, err3 := roleRepository.Save(role.RoleEntity{Name: "testRole2"})
	_, err4 := roleRepository.Save(role.RoleEntity{Name: "testRole3"})
	_, err5 := roleRepository.Save(role.RoleEntity{Name: "testRole4"})
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
}
