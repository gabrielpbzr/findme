package core_test

import (
	"database/sql"
	"testing"

	"github.com/gabrielpbzr/findme/core"
	"github.com/gabrielpbzr/findme/infra"
	"github.com/stretchr/testify/assert"
)

var check *assert.Assertions

func setup(t *testing.T) (*sql.DB, error) {
	check = assert.New(t)
	dbHandler := infra.Database{DSN: ":memory:"}
	conn, err := dbHandler.Open()
	if err != nil {
		t.Errorf("Couldn't  open database")
	}

	// Initialize database structure
	infra.InitDB(conn)
	return conn, nil
}

func getPosition() *core.Position {
	longitude := -42.875326
	latitude := -20.757294
	return core.CreatePosition(longitude, latitude)
}

func TestInsertRecords(t *testing.T) {
	conn, err := setup(t)
	if err != nil {
		t.Errorf("Couldn't initialize database")
	}

	pos := getPosition()
	service := core.NewPositionService(conn)

	t.Run("should insert a record", func(t *testing.T) {
		service.Create(pos)
		check.NotNil(pos.Id)
		count, err := service.Count()
		if err != nil {
			t.FailNow()
		}
		check.Equal(1, count)
		service.Truncate()
	})
}

func TestDeleteRecords(t *testing.T) {
	conn, err := setup(t)
	if err != nil {
		t.Errorf("Couldn't initialize database")
	}

	// Initialize database structure
	infra.InitDB(conn)
	pos := getPosition()
	service := core.NewPositionService(conn)

	t.Run("should delete a record", func(t *testing.T) {
		service.Create(pos)
		service.Delete(pos.Id)
		count, err := service.Count()
		if err != nil {
			t.FailNow()
		}
		check.Equal(0, count)
		service.Truncate()
	})
}
