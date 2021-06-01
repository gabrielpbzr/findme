package core

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type PositionService interface {
	Create(p *Position) error
	Delete(uuid uuid.UUID) error
	Get(uuid uuid.UUID) (*Position, error)
	//List(start int, quantity int) ([]Position, error)
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
		return fmt.Errorf("Couldn't record position: " + err.Error())
	}
	tx.Commit()
	return nil
}

func (service *PositionServiceDB) Delete(uuid uuid.UUID) error {
	//iniciamos uma transação
	tx, err := service.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM position WHERE uuid = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(uuid.String())
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("Couldn't delete position record: " + err.Error())
	}
	tx.Commit()
	return nil
}

func (service *PositionServiceDB) Get(uuid uuid.UUID) (*Position, error) {
	stmt, err := service.db.Prepare("SELECT uuid, longitude, latitude, timestamp FROM position WHERE uuid = ?")
	if err != nil {
		return nil, fmt.Errorf("Couldn't retrieve position record: " + err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(uuid.String())
	if err != nil {
		return nil, fmt.Errorf("Couldn't retrieve position record: " + err.Error())
	}
	defer rows.Close()
	var position Position
	for rows.Next() {
		err = rows.Scan(&position.Id, &position.Longitude, &position.Latitude, &position.Timestamp)
		if err != nil {
			return nil, err
		}
	}

	return &position, nil
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
