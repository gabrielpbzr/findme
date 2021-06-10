package infra_test

import (
	"testing"

	"github.com/gabrielpbzr/findme/infra"
	"github.com/stretchr/testify/assert"
)

func TestOpenDatabase(t *testing.T) {
	t.Run("Should open a database connection", func(t *testing.T) {
		database := &infra.Database{DSN: ":memory:"}
		conn, err := database.Open()
		if err != nil {
			t.Fail()
		}
		defer conn.Close()
		assert.NotNil(t, conn)
	})
}

func TestCloseDatabase(t *testing.T) {
	t.Run("Should close an open database connection", func(t *testing.T) {
		database := &infra.Database{DSN: ":memory:"}
		conn, err := database.Open()
		if err != nil {
			t.Fail()
		}
		database.Close(conn)
		if err != nil {
			t.Fail()
		}
		assert.Equal(t, 0, conn.Stats().OpenConnections)
	})
}
