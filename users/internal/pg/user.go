package pg

import (
	"context"
	"time"
)

type User struct {
	Id        int        `db:"id,omitempty"`
	Firstname string     `db:"firstname,omitempty"`
	Username  string     `db:"username,omitempty"`
	Createdat *time.Time `db:"createdat,omitempty"`
}

func (d *db) AddUser(ctx context.Context, u *User) error {
	const q = "insert into users(id, firstname, username, createdat) values(:id, :firstname, :username, :createdat)"

	_, err := d.db.NamedExecContext(ctx, q, u)

	return err
}

func (d *db) GetUser(ctx context.Context, id int) (*User, error) {
	const q = "select * from users where id=$1"

	u := &User{}
	err := d.db.GetContext(ctx, u, q, id)

	return u, err
}

func (d *db) IsUserWithIdExists(ctx context.Context, id int) (bool, error) {
	const q = "select exists(select from users where id=$1)"

	exists := false
	err := d.db.GetContext(ctx, &exists, q, id)

	return exists, err
}
