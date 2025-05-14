package role

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
//	_, err2 := roleRepository.Save(RoleEntity{Name: "testRole"})
//	if err2 == nil {
//		res, err := roleRepository.FindById(1)
//		if err == nil {
//			assert.Equal(t, "testRole", res.Name)
//		}
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_FindByIds(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	firstId, err2 := roleRepository.Save(RoleEntity{Name: "testRole1"})
//	secondId, err3 := roleRepository.Save(RoleEntity{Name: "testRole2"})
//	thirdId, err4 := roleRepository.Save(RoleEntity{Name: "testRole3"})
//	fourthId, err5 := roleRepository.Save(RoleEntity{Name: "testRole4"})
//	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
//		ids := make([]int64, 4)
//		ids[0], ids[1], ids[2], ids[3] = firstId, secondId, thirdId, fourthId
//		res, err := roleRepository.FindByIds(ids)
//		if err != nil {
//			assert.Equal(t, 4, len(res))
//		}
//	}
//
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_FindAll(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	_, err2 := roleRepository.Save(RoleEntity{Name: "testRole1"})
//	_, err3 := roleRepository.Save(RoleEntity{Name: "testRole2"})
//	_, err4 := roleRepository.Save(RoleEntity{Name: "testRole3"})
//	_, err5 := roleRepository.Save(RoleEntity{Name: "testRole4"})
//	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
//		res, err := roleRepository.FindAll()
//		if err != nil {
//			assert.Equal(t, 4, len(res))
//		}
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
//
//func Test_DeleteById(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	_, err2 := roleRepository.Save(RoleEntity{Name: "testRole1"})
//	_, err3 := roleRepository.Save(RoleEntity{Name: "testRole2"})
//	_, err4 := roleRepository.Save(RoleEntity{Name: "testRole3"})
//	_, err5 := roleRepository.Save(RoleEntity{Name: "testRole4"})
//	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
//		res, err := roleRepository.FindAll()
//		if err == nil {
//			assert.Equal(t, 4, len(res))
//		}
//		err = roleRepository.DeleteById(1)
//		if err != nil {
//			res, err = roleRepository.FindAll()
//			if err == nil {
//				assert.Equal(t, 3, len(res))
//			}
//		}
//		roleRepository.ExecuteQuery("DELETE FROM role")
//	}
//}
//
//func Test_DeleteByIds(t *testing.T) {
//	roleRepository := NewRoleRepository(database.ConnectDb())
//	roleRepository.ExecuteQuery("ALTER SEQUENCE role_id_seq RESTART WITH 1")
//	_, err2 := roleRepository.Save(RoleEntity{Name: "testRole1"})
//	_, err3 := roleRepository.Save(RoleEntity{Name: "testRole2"})
//	_, err4 := roleRepository.Save(RoleEntity{Name: "testRole3"})
//	_, err5 := roleRepository.Save(RoleEntity{Name: "testRole4"})
//	if err2 == nil && err3 == nil && err4 == nil && err5 == nil {
//		res, err := roleRepository.FindAll()
//		if err == nil {
//			assert.Equal(t, 4, len(res))
//		}
//		ids := make([]int64, 2)
//		ids[0], ids[1] = 1, 2
//		err = roleRepository.DeleteByIds(ids)
//		if err == nil {
//			res, err = roleRepository.FindAll()
//			if err == nil {
//				assert.Equal(t, 2, len(res))
//			}
//		}
//	}
//	roleRepository.ExecuteQuery("DELETE FROM role")
//}
