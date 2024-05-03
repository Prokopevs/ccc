package pg

import (
	"context"
	"database/sql"
	"time"
)

type UserReq struct {
	Id         int        `db:"id,omitempty"`
	Firstname  string     `db:"firstname,omitempty"`
	Username   string     `db:"username,omitempty"`
	ReferralId int        `db:"referralId,omitempty"`
	Createdat  *time.Time `db:"createdat,omitempty"`
}

type UserRes struct {
	Id        int        `db:"id,omitempty"`
	Firstname string     `db:"firstname,omitempty"`
	Username  string     `db:"username,omitempty"`
	Referrals []byte     `db:"referrals,omitempty"`
	Createdat *time.Time `db:"createdat,omitempty"`
}

func (d *db) AddUser(ctx context.Context, u *UserReq) error {
	const (
		adduser     = "insert into users(id, firstname, username, createdat) values(:id, :firstname, :username, :createdat)"
		addGameQ    = "insert into game(ownerId) values($1)"
		addReferral = "UPDATE users SET referrals = array_append(referrals, $1) WHERE id = $2;"
	)
	tx, err := d.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.NamedExecContext(ctx, adduser, u)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, addGameQ, u.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	if u.ReferralId != 0 {
		_, err = tx.ExecContext(ctx, addReferral, u.Id, u.ReferralId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	return err
}

func (d *db) GetUser(ctx context.Context, id int) (*UserRes, error) {
	const q = "select * from users where id=$1"

	u := &UserRes{}
	err := d.db.GetContext(ctx, u, q, id)

	return u, err
}

func (d *db) IsUserWithIdExists(ctx context.Context, id int) (bool, error) {
	const q = "select exists(select from users where id=$1)"

	exists := false
	err := d.db.GetContext(ctx, &exists, q, id)

	return exists, err
}
