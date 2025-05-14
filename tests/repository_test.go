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
	_, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName"})
	if err2 == nil {
		res, err := employeeRepository.FindById(1)
		if err == nil {
			assert.Equal(t, "testName", res.Name)
		}
	}

	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_FindByIds(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	_, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	_, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		ids := make([]int64, 4)
		ids[0], ids[1], ids[2], ids[3] = 1, 2, 3, 4
		res, err := employeeRepository.FindByIds(ids)
		if err == nil {
			assert.Equal(t, 4, len(res))
		}
		employeeRepository.ExecuteQuery("DELETE FROM employee")
	}
}

func Test_FindAll(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	_, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	_, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := employeeRepository.FindAll()
		if err == nil {
			assert.Equal(t, 4, len(res))
		}
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_DeleteById(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	_, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	_, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := employeeRepository.FindAll()
		if err == nil {
			assert.Equal(t, 4, len(res))
		}
		err = employeeRepository.DeleteById(1)
		if err == nil {
			res, err = employeeRepository.FindAll()
			if err == nil {
				assert.Equal(t, 3, len(res))
			}
		}
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}

func Test_DeleteByIds(t *testing.T) {
	employeeRepository := employee.NewEmployeeRepository(database.ConnectDb())
	employeeRepository.ExecuteQuery("ALTER SEQUENCE employee_id_seq RESTART WITH 1")
	_, err2 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName1"})
	_, err3 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName2"})
	_, err4 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName3"})
	_, err5 := employeeRepository.Save(employee.EmployeeEntity{Name: "testName4"})
	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
		res, err := employeeRepository.FindAll()
		if err == nil {
			assert.Equal(t, 4, len(res))
		}
		ids := make([]int64, 2)
		ids[0], ids[1] = 1, 2
		err = employeeRepository.DeleteByIds(ids)
		if err == nil {
			res, err = employeeRepository.FindAll()
		}

		if err == nil {
			assert.Equal(t, 2, len(res))
		}
	}
	employeeRepository.ExecuteQuery("DELETE FROM employee")
}
