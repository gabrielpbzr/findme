package core_test

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/gabrielpbzr/findme/core"
	"github.com/gabrielpbzr/findme/infra"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var check *assert.Assertions
var service *core.PositionServiceDB

func setup(t *testing.T) error {
	check = assert.New(t)
	dbHandler := infra.Database{DSN: ":memory:"}
	conn, err := dbHandler.Open()
	if err != nil {
		t.Errorf("Couldn't  open database")
	}
	// Initialize database structure
	infra.InitDB(conn)
	service = core.NewPositionService(conn)
	return nil
}

func getPosition() *core.Position {
	longitude := -42.875326
	latitude := -20.757294
	return core.CreatePosition(longitude, latitude)
}

func TestInsertRecords(t *testing.T) {
	err := setup(t)
	if err != nil {
		t.Errorf("couldn't initialize database: %s", err.Error())
	}

	pos := getPosition()

	t.Run("should insert a record", func(t *testing.T) {
		service.Create(pos)
		check.NotNil(pos.Id)
		count, err := service.Count()
		if err != nil {
			t.Errorf("couldn't insert record into database: %s", err.Error())
		}
		check.Equal(1, count)
	})

}

func TestQueryRecords(t *testing.T) {
	err := setup(t)
	if err != nil {
		t.Errorf("couldn't initialize database: %s", err.Error())
	}

	err = loadData()
	if err != nil {
		t.Errorf("couldn't load records from file into database: %s", err.Error())
	}

	t.Run("should query a list of records", func(t *testing.T) {
		quantity := 5
		records, err := service.List(0, quantity)
		if err != nil {
			t.Errorf("couldn't query records from database: %s", err.Error())
		}
		check.Equal(quantity, len(records))
	})

	t.Run("should query a single record", func(t *testing.T) {
		id := uuid.MustParse("b1a0044f-9f2a-43e1-854e-d1f5a45aaa98")
		longitude := -42.8848640
		latitude := -20.7564320

		record, err := service.Get(id)
		if err != nil {
			t.Errorf("Couldn't query records from database: %s", err.Error())
		}
		check.Equal(id, record.Id)
		check.Equal(longitude, record.Longitude)
		check.Equal(latitude, record.Latitude)
	})

	t.Run("should get nil when quering a inexistent record", func(t *testing.T) {
		id := uuid.MustParse("02a9c57a-bd37-4fdc-b44d-995d210b370a")

		record, err := service.Get(id)
		if err != nil {
			t.Errorf("couldn't query records from database: %s", err.Error())
		}
		check.Nil(record)
	})
}

func loadData() error {
	content, err := ioutil.ReadFile("../data/test_data.json")
	if err != nil {
		return err
	}
	var positions []core.Position

	err = json.Unmarshal(content, &positions)
	if err != nil {
		return err
	}

	for _, p := range positions {
		err = service.Create(&p)
		if err != nil {
			return err
		}
	}

	return nil
}
