package pg_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/Prokopevs/ccc/game/internal/pg"
	"github.com/jmoiron/sqlx"
)

func TestGame_GetGame(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d := pg.Newdb(sqlxDB)

	ctx := context.Background()
	userId := 1

	game := &model.Game{
		OwnerId:    1,
		Score:      20,
		GasStorage: 1,
		GasMining:  1,
		Protection: 1,
	}

	row := sqlmock.NewRows([]string{"ownerid", "score", "gasstorage", "gasmining", "protection"}).
		AddRow(game.OwnerId, game.Score, game.GasStorage, game.GasMining, game.Protection)

	mock.ExpectQuery("select (.+) from game where (.+)").WithArgs(userId).
		WillReturnRows(row)

	result, err := d.GetGame(ctx, userId)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if result.OwnerId != game.OwnerId || result.Score != game.Score || result.GasStorage != game.GasStorage {
		t.Errorf("expected %v, got %v", game, result)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGame_UpdateScore(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d := pg.Newdb(sqlxDB)

	ctx := context.Background()

	score := &model.Score{
		Id:    1,
		Score: 20,
	}

	mock.ExpectExec("UPDATE game SET (.+) WHERE (.+);").
		WithArgs(score.Score, score.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = d.UpdateScore(ctx, score)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGame_UpdateMultiplicator(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	d := pg.Newdb(sqlxDB)

	ctx := context.Background()

	multipUpdate := &model.MultipUpdate{
		Id:       1,
		NameType: "gasStorage",
	}

	prices := map[int]int{
		1: 10,
		2: 20,
		3: 30,
		4: 40,
		5: 50,
		6: 60,
	}

	mock.ExpectBegin()

	// Request 1
	row := sqlmock.NewRows([]string{"gasstorage"}).AddRow(1)

	mock.ExpectQuery("SELECT (.+) FROM game WHERE (.+)").
		WithArgs(multipUpdate.Id).
		WillReturnRows(row)

	// Request 2
	row = sqlmock.NewRows([]string{"score"}).AddRow(20)

	mock.ExpectQuery("select score from game where (.+)").
		WithArgs(multipUpdate.Id).
		WillReturnRows(row)

	// Request 3
	mock.ExpectExec("UPDATE game SET score = score - (.+) WHERE (.+);").
        WithArgs(prices[1], multipUpdate.Id).
        WillReturnResult(sqlmock.NewResult(0, 1))

	// Request 4
	mock.ExpectExec("UPDATE game SET (.+) WHERE (.+);").
		WithArgs(multipUpdate.Id).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectCommit()

	err = d.UpdateMultiplicator(ctx, multipUpdate)
    if err != nil {
        t.Errorf("unexpected error: %s", err)
    }

    err = mock.ExpectationsWereMet()
    if err != nil {
        t.Errorf("there were unfulfilled expectations: %s", err)
    }
}
