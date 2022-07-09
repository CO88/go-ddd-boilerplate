package repository

import (
	"context"
	"database/sql"

	"github.com/CO88/go-ddd-boilerplate/user/domain"
	log "github.com/sirupsen/logrus"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewUserRepository(conn *sql.DB) domain.UserRespository {
	return &mysqlUserRepository{conn}
}

func (m *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.User, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Error(errRow)
		}
	}()

	result = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Password,
			&t.PhoneNumber,
			&t.Email,
			&t.CreateAt,
			&t.UpdateAt,
		)

		if err != nil {
			log.Error(err)
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil
}

func (u *mysqlUserRepository) FindOneById(ctx context.Context, id int64) (res domain.User, err error) {
	query := `SELECT * FROM user WHERE ID = ?`

	list, err := u.fetch(ctx, query, id)
	if err != nil {
		return domain.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, nil
	}
	return
}
