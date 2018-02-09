package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"github.com/whosonfirst/go-whosonfirst-sqlite-privatezen"
	_ "log"
)

type TripsTable struct {
	privatezen.PrivatezenTable
	name string
}

func NewTripsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewTripsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewTripsTable() (sqlite.Table, error) {

	t := TripsTable{
		name: "trips",
	}

	return &t, nil
}

func (t *TripsTable) Name() string {
	return t.name
}

func (t *TripsTable) Schema() string {

	sql := `CREATE TABLE %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		arrival TEXT,
		departure TEXT,
		wof_id INTEGER,
		status_id INTEGER,
		notes TEXT
	);

	CREATE INDEX `%s_by_date` ON trips (`arrival`, `departure`);
	CREATE INDEX `%s_by_wofid` ON trips (`wof_id`);`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name(), t.Name())
}

func (t *TripsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *TripsTable) IndexRecord(db sqlite.Database, i interface{}) error {

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
