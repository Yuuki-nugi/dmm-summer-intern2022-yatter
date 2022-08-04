package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	status struct {
		db *sqlx.DB
	}
)

// Create accout repository
func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

func (r *status) CreateStatus(ctx context.Context, status *object.Status) (*object.Status, error) {
	result, err := r.db.ExecContext(ctx, "insert into status (account_id, content) value (?, ?)", status.AccountID, status.Content)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	status_id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}
	if err := r.db.QueryRowxContext(ctx, "select * from status where id = ?", status_id).StructScan(status); err != nil {
		log.Println(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("%w", err)
	}
	return status, nil
}

// FindByUsername : ユーザ名からユーザを取得
// func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
// 	entity := new(object.Account)
// 	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, nil
// 		}

// 		return nil, fmt.Errorf("%w", err)
// 	}

// 	return entity, nil
// }
