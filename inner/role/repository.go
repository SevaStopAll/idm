package role

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type RoleRepository struct {
	db *sqlx.DB
}

func NewRoleRepository(db *sqlx.DB) *RoleRepository {
	return &RoleRepository{db}
}

type RoleEntity struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (r *RoleRepository) FindById(id int64) (res RoleEntity, err error) {
	err = r.db.Get(&res, "SELECT * from role where id = $1", id)
	return res, err
}

func (r *RoleRepository) FindAll() (res []RoleEntity, err error) {
	err = r.db.Select(&res, "SELECT * from role")
	return res, err
}

func (r *RoleRepository) Save(employee RoleEntity) (id int64, err error) {
	result, err := r.db.Query("INSERT INTO role (name) VALUES($1) RETURNING id", employee.Name)
	if err != nil {
		return 0, err
	}
	if result.Next() {
		result.Scan(&id)
		return id, err
	}
	return 0, err
}

func (r *RoleRepository) FindByIds(ids []int64) (res []RoleEntity, err error) {
	query, args, err := sqlx.In("SELECT * from role where id IN(?)", ids)
	query = r.db.Rebind(query)
	rows, err := r.db.Queryx(query, args...)
	for rows.Next() {
		var employee RoleEntity
		err = rows.StructScan(&employee)
		res = append(res, employee)
	}
	return res, err
}

func (r *RoleRepository) DeleteById(id int64) (err error) {

	_, err = r.db.Query("DELETE from role where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleRepository) DeleteByIds(ids []int64) (err error) {
	query, args, err := sqlx.In("DELETE from role where id IN(?)", ids)
	if err != nil {
		return err
	}
	query = r.db.Rebind(query)
	_, err = r.db.Query(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleRepository) ExecuteQuery(query string) {
	r.db.MustExec(query)
}
