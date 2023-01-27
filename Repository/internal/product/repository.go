package product

import (
	"context"
	"database/sql"
	"Repository/internal/domain"
	"fmt"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetOne(ctx context.Context, id int) (domain.Product, error)
	Exists(ctx context.Context, codeValue string) bool
	CodeValueUsed(ctx context.Context, codeValue string, id int) bool
	Store(ctx context.Context, p domain.Product) (domain.Product, error)
	Update(ctx context.Context, p domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	query := "SELECT * FROM products;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	var products []domain.Product

	for rows.Next() {
		p := domain.Product{}
		_ = rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
		products = append(products, p)
	}

	return products, nil
}



func (r *repository) GetOne(ctx context.Context, id int) (domain.Product, error) {

	p := domain.Product{}

	query := "SELECT * FROM products WHERE id=?;"
	row := r.db.QueryRow(query, id)


	err := row.Scan(&p.ID, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)
	if err != nil {
		return p, err
	}

	return p, nil
}




func (r *repository) Exists(ctx context.Context, codeValue string) bool {
	query := "SELECT code_value FROM products WHERE code_Value=?;"
	row := r.db.QueryRow(query, codeValue)
	err := row.Scan(&codeValue)
	return err == nil
}


func (r *repository) CodeValueUsed(ctx context.Context, codeValue string, id int) bool {
	query := "SELECT code_value FROM products WHERE code_Value=? AND id<>?;"
	row := r.db.QueryRow(query, codeValue, id)
	err := row.Scan(&codeValue)
	return err == nil
}



func (r *repository) Store(ctx context.Context, p domain.Product) (domain.Product, error) {
/*
	if(Exists(ctx, p.CodeValue)){
		return domain.Product{}, ErrProductCodeAlreadyExists
	}
*/
	fmt.Println("entre al repository")


	query := "INSERT INTO products(name,quantity,code_value,is_published,expiration,price,id_warehouse) VALUES (?,?,?,?,?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println("rompio en el prepare")

		return domain.Product{}, err
	}

	res, err := stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.ID)
	if err != nil {
		fmt.Println("rompio en el exec")

		return domain.Product{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return domain.Product{}, err
	}

	p.ID=int(id)
		return p, nil
}


func (r *repository) Update(ctx context.Context, p domain.Product) (domain.Product, error) {

	if(r.CodeValueUsed(ctx, p.CodeValue, p.ID)){
		return domain.Product{}, ErrProductCodeAlreadyExists
	}


	query := "UPDATE products SET name=?, quantity=?, code_value=?, is_published=?, expiration=?, price=? WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		fmt.Println("rompio en el prepare")
		return domain.Product{}, err
	}

	res, err := stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.ID)
	if err != nil {
		fmt.Println("rompio en el execute")

		return domain.Product{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}



func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare("DELETE FROM products WHERE id=?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rows, errRowsAffected:= res.RowsAffected()

	if errRowsAffected != nil {
		return errRowsAffected
	}

	if rows < 1 {
		return ErrNotFound
	}

	return nil
}