package role

//
//import (
//	"github.com/stretchr/testify/assert"
//	"idm/inner/database"
//	"testing"
//)
//
//func Test_Save(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	res, err := roleRepository.Save(RoleEntity{Name: "testRole"})
//	if err != nil {
//		assert.Equal(t, int64(1), res)
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_FindById(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	roleRepository.Save(RoleEntity{Name: "testRole"})
//	res, err := roleRepository.FindById(1)
//	if err == nil {
//		assert.Equal(t, "testRole", res.Name)
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_FindByIds(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	roleRepository.Save(RoleEntity{Name: "testRole1"})
//	roleRepository.Save(RoleEntity{Name: "testRole2"})
//	roleRepository.Save(RoleEntity{Name: "testRole3"})
//	roleRepository.Save(RoleEntity{Name: "testRole4"})
//	ids := make([]int64, 4)
//	ids[0], ids[1], ids[2], ids[3] = 1, 2, 3, 4
//	res, err := roleRepository.FindByIds(ids)
//	if err != nil {
//		assert.Equal(t, 4, len(res))
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_FindAll(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	roleRepository.Save(RoleEntity{Name: "testRole1"})
//	roleRepository.Save(RoleEntity{Name: "testRole2"})
//	roleRepository.Save(RoleEntity{Name: "testRole3"})
//	roleRepository.Save(RoleEntity{Name: "testRole4"})
//	res, err := roleRepository.FindAll()
//	if err != nil {
//		assert.Equal(t, 4, len(res))
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_DeleteById(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	roleRepository.Save(RoleEntity{Name: "testRole1"})
//	roleRepository.Save(RoleEntity{Name: "testRole2"})
//	roleRepository.Save(RoleEntity{Name: "testRole3"})
//	roleRepository.Save(RoleEntity{Name: "testRole4"})
//	res, err := roleRepository.FindAll()
//	if err == nil {
//		assert.Equal(t, 4, len(res))
//	}
//	roleRepository.DeleteById(1)
//	res, err = roleRepository.FindAll()
//	if err == nil {
//		assert.Equal(t, 3, len(res))
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_DeleteByIds(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	roleRepository.Save(RoleEntity{Name: "testRole1"})
//	roleRepository.Save(RoleEntity{Name: "testRole2"})
//	roleRepository.Save(RoleEntity{Name: "testRole3"})
//	roleRepository.Save(RoleEntity{Name: "testRole4"})
//	res, err := roleRepository.FindAll()
//	if err == nil {
//		assert.Equal(t, 4, len(res))
//	}
//	ids := make([]int64, 2)
//	ids[0], ids[1] = 1, 2
//	roleRepository.DeleteByIds(ids)
//	res, err = roleRepository.FindAll()
//	if err == nil {
//		assert.Equal(t, 2, len(res))
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
