package postgres

import (
	"book_service/genproto/book_service"
	"book_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type bookCategoryRepo struct {
	db *sqlx.DB
}

func NewBookCategoryRepo(db *sqlx.DB) repo.BookCategoryRepoI {
	return &bookCategoryRepo{
		db: db,
	}
}

func (r *bookCategoryRepo) Create(req *book_service.CreateBookCategory) (string, error) {
	var (
		id uuid.UUID
	)
	tx, err := r.db.Begin()
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	id, err = uuid.NewRandom()

	if err != nil {
		return "", err
	}

	query := `INSERT INTO book_category ( id, name) VALUES($1,$2)`

	_, err = r.db.Exec(query, id, req.Name)

	if err != nil {
		return "", err
	}

	return id.String(), nil

}

func (r *bookCategoryRepo) GetAll(req *book_service.GetAllBookCategoryRequest) (*book_service.GetAllBookCategoryResponse, error) {
	var (
		filter string
		count  int32
		arr    []*book_service.BookCategory
	)

	args := make(map[string]interface{})

	if req.Name != "" {
		filter += " AND name ILIKE '%' || :filter_name ||'%'"
		args["filter_name"] = req.Name
	}

	filter += " LIMIT :limi OFFSET :offs"
	args["limi"] = req.Limit
	args["offs"] = req.Offset
	countQuery := `SELECT count(1) FROM book_category WHERE true` + filter
	rows, err := r.db.NamedQuery(countQuery, args)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return nil, err
		}
	}

	query := `SELECT id,name,created_at,updated_at FROM book_category WHERE true` + filter
	rows, err = r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var bookCategory book_service.BookCategory

		err = rows.Scan(&bookCategory.Id, &bookCategory.Name, &bookCategory.CreatedAt, &bookCategory.UpdatedAt)

		if err != nil {

			return nil, err
		}

		arr = append(arr, &bookCategory)
	}

	return &book_service.GetAllBookCategoryResponse{
		Bookcategorylist: arr,
		Count:            count,
	}, nil
}

func (r *bookCategoryRepo) GetById(id string) (*book_service.BookCategory, error) {
	var bookCategory book_service.BookCategory

	query := `SELECT id,name,created_at,updated_at FROM book_category WHERE id = $1`

	rows, err := r.db.Query(query, id)
	if err != nil {

		return nil, err
	}
	rows.Next()
	err = rows.Scan(
		&bookCategory.Id,
		&bookCategory.Name,
		&bookCategory.CreatedAt,
		&bookCategory.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &book_service.BookCategory{
		Id:        bookCategory.Id,
		Name:      bookCategory.Name,
		CreatedAt: bookCategory.CreatedAt,
		UpdatedAt: bookCategory.UpdatedAt,
	}, nil
}

func (r *bookCategoryRepo) Update(req *book_service.BookCategory) (*book_service.MsgResponse, error) {
	_, err := r.db.Exec(
		`UPDATE book_category SET name=$1,updated_at=Now() WHERE id=$2`,
		req.Name,
		req.Id,
	)
	if err != nil {
		return &book_service.MsgResponse{
			Msg: "error",
		}, nil
	}
	return &book_service.MsgResponse{
		Msg: "updated",
	}, nil
}
func (r *bookCategoryRepo) Delete(id string) (*book_service.MsgResponse, error) {
	_, err := r.db.Exec(
		`DELETE FROM book_category WHERE id = $1;`,
		id,
	)
	if err != nil {
		return &book_service.MsgResponse{
			Msg: "error",
		}, nil
	}
	return &book_service.MsgResponse{
		Msg: "deleted",
	}, nil

}
