package tests

import (
	"github.com/stretchr/testify/assert"
	"idm/inner/database"
	"idm/inner/employee"
	"testing"
)

func Test_Save(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	res, err := employeeRepository.Save(employee.EmployeeEntity{Name: "testName"})
	if err == nil {
		assert.Equal(t, int64(1), res)
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindById(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName"})
	res, err := employeeRepository.FindById(1)
	if err == nil {
		assert.Equal(t, "testName", res.Name)
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindByIds(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	ids := make([]int64, 4)
	ids[0], ids[1], ids[2], ids[3] = 1, 2, 3, 4
	res, err := employeeRepository.FindByIds(ids)
	if err == nil {
		assert.Equal(t, 4, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindAll(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	res, err := employeeRepository.FindAll()
	if err == nil {
		assert.Equal(t, 4, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_DeleteById(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	res, err := employeeRepository.FindAll()
	if err == nil {
		assert.Equal(t, 4, len(res))
	}
	employeeRepository.DeleteById(1)
	res, err = employeeRepository.FindAll()
	if err == nil {
		assert.Equal(t, 3, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_DeleteByIds(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	res, err := employeeRepository.FindAll()
	if err == nil {
		assert.Equal(t, 4, len(res))
	}
	ids := make([]int64, 2)
	ids[0], ids[1] = 1, 2
	employeeRepository.DeleteByIds(ids)
	res, err = employeeRepository.FindAll()
	if err == nil {
		assert.Equal(t, 2, len(res))
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}
