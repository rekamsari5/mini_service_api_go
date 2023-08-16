package customers

import (
	"context"
	"database/sql"
	"service-mini-restapi/helper"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type repository struct {
	db *sql.DB
}

const (
	tbl_customer = "tbl_customer"
)

type Repository interface {
	Create(param RequestCustomer) (ResultRequest, error)
	GetAll(param SearchCustomer) ([]Customers, error)
	Delete(param RequestDelete) error
	Update(param RequestUpdate) (ResultUpdate, error)
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Create(param RequestCustomer) (ResultRequest, error) {
	var result ResultRequest

	location, _ := time.LoadLocation("Asia/Jakarta")

	now := time.Now().In(location)
	createdAt := now.Format("2006-01-02 15:04:05")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer tx.Rollback()

	query := sq.Insert(tbl_customer).
		Columns("name", "address", "created_at").
		Values(param.Name, param.Address, createdAt)
	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	res, err := tx.ExecContext(ctx, sql, args...)
	helper.PanicIfError(err)

	err = tx.Commit()
	helper.PanicIfError(err)

	result.Name = param.Name
	result.Address = param.Address
	result.CreatedAt = createdAt
	result.ID, _ = res.LastInsertId()

	return result, nil

}

func (r *repository) GetAll(param SearchCustomer) ([]Customers, error) {
	var results []Customers

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := sq.Select("id", "name", "address", "created_at", "updated_at").
		From(tbl_customer).OrderBy("id DESC")

	if len(param.Name) > 0 {
		query = query.Where(sq.Like{"name": "%" + param.Name + "%"})
	}
	if len(param.Address) > 0 {
		query = query.Where(sq.Like{"address": "%" + param.Address + "%"})
	}

	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	rows, err := r.db.QueryContext(ctx, sql, args...)
	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		result := Customers{}
		err := rows.Scan(&result.ID, &result.Name, &result.Address, &result.CreatedAt, &result.UpdateAt)
		helper.PanicIfError(err)
		results = append(results, result)
	}

	return results, nil

}

func (r *repository) Delete(param RequestDelete) error {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	query := sq.Delete(tbl_customer).Where(sq.Eq{"id": param.ID})
	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	_, err = r.db.ExecContext(ctx, sql, args...)
	helper.PanicIfError(err)
	return nil
}

func (r *repository) Update(param RequestUpdate) (ResultUpdate, error) {
	var result ResultUpdate

	location, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(location)
	updateAt := now.Format("2006-01-02 15:04:05")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	helper.PanicIfError(err)

	defer tx.Rollback()

	query := sq.Update(tbl_customer).
		Set("name", param.Name).
		Set("address", param.Address).
		Set("updated_at", updateAt).
		Where(sq.Eq{"id": param.ID})

	sql, args, err := query.ToSql()
	helper.PanicIfError(err)

	_, err = tx.ExecContext(ctx, sql, args...)
	helper.PanicIfError(err)

	err = tx.Commit()
	helper.PanicIfError(err)

	result.Name = param.Name
	result.Address = param.Address
	result.UpdateAt = updateAt

	return result, nil

}
