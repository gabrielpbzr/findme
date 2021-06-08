package core

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type PositionService interface {
	Create(p *Position) error
	Get(uuid uuid.UUID) (*Position, error)
	List(start int, quantity int) ([]Position, error)
	Count() (int, error)
	Truncate() error
}

type PositionServiceDB struct {
	db *sql.DB
}

func NewPositionService(dbConn *sql.DB) *PositionServiceDB {
	return &PositionServiceDB{db: dbConn}
}

func (service *PositionServiceDB) Create(p *Position) error {
	//iniciamos uma transação
	tx, err := service.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO position (longitude, latitude, time_stamp, uuid) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Longitude, p.Latitude, p.Timestamp, p.Id.String())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("couldn't record position: %s", err.Error())
	}
	tx.Commit()
	return nil
}

func (service *PositionServiceDB) List(start int, quantity int) ([]*Position, error) {
	rows, err := service.db.Query("SELECT uuid, longitude, latitude, time_stamp FROM position LIMIT ? OFFSET ?", quantity, start)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	positions := make([]*Position, 0, quantity)
	for rows.Next() {
		position := new(Position)
		err = rows.Scan(&position.Id, &position.Longitude, &position.Latitude, &position.Timestamp)
		if err != nil {
			return nil, err
		}
		positions = append(positions, position)
	}
	return positions, nil
}

func (service *PositionServiceDB) Get(uuid uuid.UUID) (*Position, error) {
	query := "SELECT uuid, longitude, latitude, time_stamp FROM position WHERE uuid = ?"

	row := service.db.QueryRow(query, uuid.String())
	position := &Position{}
	err := row.Scan(&position.Id, &position.Longitude, &position.Latitude, &position.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return position, nil
}

func (service *PositionServiceDB) Count() (int, error) {
	stmt, err := service.db.Begin()
	if err != nil {
		return 0, err
	}
	rows, err := stmt.Query("SELECT COUNT(*) FROM position")
	if err != nil {
		return 0, err
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, nil
		}
	}
	defer rows.Close()

	return count, nil
}

func (service *PositionServiceDB) Truncate() error {
	//iniciamos uma transação
	tx, err := service.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM position")
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
