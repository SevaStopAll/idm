package employee

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

type EmployeeEntity struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (r *EmployeeRepository) FindById(id int64) (res EmployeeEntity, err error) {
	err = r.db.Get(&res, "SELECT * from employee where id = $1", id)
	return res, err
}

func (r *EmployeeRepository) FindAll() (res []EmployeeEntity, err error) {
	err = r.db.Select(&res, "SELECT * from employee")
	return res, err
}

func (r *EmployeeRepository) Save(employee EmployeeEntity) (id int64, err error) {
	result, err := r.db.Query("INSERT INTO employee (name) VALUES($1) RETURNING id", employee.Name)
	if err != nil {
		return 0, err
	}
	if result.Next() {
		result.Scan(&id)
		return id, err
	}
	return 0, err
}

func (r *EmployeeRepository) FindByIds(ids []int64) (res []EmployeeEntity, err error) {
	query, args, err := sqlx.In("SELECT * from employee where id IN(?)", ids)
	query = r.db.Rebind(query)
	rows, err := r.db.Queryx(query, args...)
	for rows.Next() {
		var employee EmployeeEntity
		err = rows.StructScan(&employee)
		res = append(res, employee)
	}
	return res, err
}

func (r *EmployeeRepository) DeleteById(id int64) (err error) {

	_, err = r.db.Query("DELETE from employee where id = $1", id)
	return err
}

func (r *EmployeeRepository) DeleteByIds(ids []int64) (err error) {
	query, args, err := sqlx.In("DELETE from employee where id IN(?)", ids)
	query = r.db.Rebind(query)
	_, err = r.db.Query(query, args...)
	return err
}

func (r *EmployeeRepository) ExecuteQuery(query string) {
	r.db.MustExec(query)
}

//3. Написать интеграционные тесты, проверяющие работу API всех созданых репозиториев.
