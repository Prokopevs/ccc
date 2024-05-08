package pg

import (
	"context"
	"database/sql"

	"github.com/Prokopevs/ccc/users/internal/model"
)

func (d *db) AddUser(ctx context.Context, u *model.UserReq) error {
	const (
		adduser     = "insert into users(id, firstname, username, createdat) values(:id, :firstname, :username, :createdat)"
		addGameQ    = "insert into game(ownerId) values($1)"
		addReferral = "insert into userReferral(inviterId, referralId) values($1, $2)"
		UpdateScoreQ = "UPDATE game SET score = score + $1 WHERE ownerId = $2;"
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

	if u.InviterId != 0 {
		_, err = tx.ExecContext(ctx, addReferral, u.InviterId, u.Id)
		if err != nil {
			tx.Rollback()
			return err
		}

		_, err := d.db.ExecContext(ctx, UpdateScoreQ, 100, u.InviterId)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	return err
}

func (d *db) GetUser(ctx context.Context, id int) (*model.UserRes, error) {
	const q = "select * from users where id=$1"

	u := &model.UserRes{}

	err := d.db.GetContext(ctx, u, q, id)

	return u, err
}

func (d *db) IsUserWithIdExists(ctx context.Context, id int) (bool, error) {
	const q = "select exists(select from users where id=$1)"

	exists := false
	err := d.db.GetContext(ctx, &exists, q, id)

	return exists, err
}

func (d *db) GetUserReferrals(ctx context.Context, id int) ([]*model.UserReferrals, error) {
	const q = "SELECT u.firstname, u.username, ur.referralId FROM users u JOIN userReferral ur ON u.id = ur.referralId WHERE ur.inviterId = $1"

	referrals := []*model.UserReferrals{}
	err := d.db.SelectContext(ctx, &referrals, q, id)

	return referrals, err
}
