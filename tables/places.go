package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"github.com/whosonfirst/go-whosonfirst-sqlite-privatezen"
	_ "log"
)

type PlacesTable struct {
	privatezen.PrivatezenTable
	name string
}

func NewPlacesTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewPlacesTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewPlacesTable() (sqlite.Table, error) {

	t := PlacesTable{
		name: "places",
	}

	return &t, nil
}

func (t *PlacesTable) Name() string {
	return t.name
}

func (t *PlacesTable) Schema() string {

	sql := `CREATE TABLE %s (
		wof_id INTEGER PRIMARY KEY,
		body TEXT,
		created TEXT
	);`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name())
}

func (t *PlacesTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *PlacesTable) IndexRecord(db sqlite.Database, i interface{}) error {

     return errors.New("Please implement me")

	conn, err := db.Conn()

	if err != nil {
		return err
	}

	tx, err := conn.Begin()

	if err != nil {
		return err
	}

	return tx.Commit()
}
