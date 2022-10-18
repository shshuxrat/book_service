package postgres

import (
	"book_service/genproto/book_service"
	"book_service/storage/repo"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type bookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) repo.BookRepoI {
	return &bookRepo{
		db: db,
	}
}

func (r *bookRepo) Create(req *book_service.CreateBook) (string, error) {
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

	query := `INSERT INTO book ( id, name,category_id) VALUES($1,$2,$3)`

	_, err = r.db.Exec(query, id, req.Name, req.CategoryId)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *bookRepo) GetAll(req *book_service.GetAllBookRequest) (*book_service.GetAllBookResponse, error) {
	var (
		filter string
		count  int32
		arr    []*book_service.Book
	)
	args := make(map[string]interface{})

	if req.Name != "" {
		filter += " AND name ILIKE '%' || :filter_name ||'%'"
		args["filter_name"] = req.Name
	}
	filter += " LIMIT :limi OFFSET :offs"
	args["limi"] = req.Limit
	args["offs"] = req.Offset

	countQuery := `SELECT count(1) FROM book WHERE true` + filter
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
	query := `SELECT id,name,category_id,created_at,updated_at FROM book WHERE true` + filter
	rows, err = r.db.NamedQuery(query, args)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var book book_service.Book
		err = rows.Scan(&book.Id, &book.Name, &book.CategoryId, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		arr = append(arr, &book)
	}
	return &book_service.GetAllBookResponse{
		BookList: arr,
		Count:    count,
	}, nil

}

func (r *bookRepo) GetById(id string) (*book_service.GetBookByIdResponse, error) {
	var bookData book_service.GetBookByIdResponse
	query := `SELECT b.id, b.name,c.name
			FROM book AS b, book_category AS c
			WHERE c.id = b.category_id AND b.id = $1`
	row, err := r.db.Query(query, id)
	if err != nil {
		return nil, err

	}
	row.Next()
	row.Scan(
		&bookData.Id,
		&bookData.Name,
		&bookData.Category,
	)

	return &book_service.GetBookByIdResponse{
		Id:       bookData.Id,
		Name:     bookData.Name,
		Category: bookData.Category,
	}, nil
}

func (r *bookRepo) Update(req *book_service.UpdateBook) (*book_service.MsgRespons, error) {
	_, err := r.db.Exec(
		`UPDATE book SET name=$1,category_id=$2,updated_at=Now() WHERE id=$3`,
		req.Name,
		req.CategoryId,
		req.Id,
	)
	if err != nil {
		return &book_service.MsgRespons{
			Msg: "error",
		}, err
	}

	return &book_service.MsgRespons{
		Msg: "updated",
	}, nil
}

func (r *bookRepo) Delete(id string) (*book_service.MsgRespons, error) {
	_, err := r.db.Exec(
		`DELETE FROM book WHERE id = $1;`,
		id,
	)
	if err != nil {
		return &book_service.MsgRespons{
			Msg: "error",
		}, nil
	}
	return &book_service.MsgRespons{
		Msg: "deleted",
	}, nil
}
