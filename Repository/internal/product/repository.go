package product

import (
	"context"
	"database/sql"
	"Repository/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetOne(ctx context.Context, id int) (domain.Product, error)
	
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

/*

func (r *repository) Exists(ctx context.Context, productCode string) bool {
	query := "SELECT product_code FROM products WHERE product_code=?;"
	row := r.db.QueryRow(query, productCode)
	err := row.Scan(&productCode)
	return err == nil
}

func (r *repository) Store(ctx context.Context, p domain.Product) (domain.Product, error) {

	stmt, err := r.db.Prepare("INSERT INTO products(description,expiration_rate,freezing_rate,height,lenght,netweight,product_code,recommended_freezing_temperature,width,id_product_type,id_seller) VALUES (?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()

	var res sql.Result
	res, err := stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
	if err != nil {
		return domain.Product{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product) (domain.Product, error) {

	query := "UPDATE products SET description=?, expiration_rate=?, freezing_rate=?, height=?, lenght=?, netweight=?, product_code=?, recommended_freezing_temperature=?, width=?, id_product_type=?, id_seller=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return domain.Product{}, err
	}

	defer stmt.Close()


	_, err := stmt.Exec(p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)

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

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows < 1 {
		return ErrNotFound
	}

	return nil
}

*/