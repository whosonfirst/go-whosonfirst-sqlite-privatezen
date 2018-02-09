package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"github.com/whosonfirst/go-whosonfirst-sqlite-privatezen"
	_ "log"
)

type ListsTable struct {
	privatezen.PrivatezenTable
	name string
}

func NewListsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewListsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewListsTable() (sqlite.Table, error) {

	t := ListsTable{
		name: "lists",
	}

	return &t, nil
}

func (t *ListsTable) Name() string {
	return t.name
}

func (t *ListsTable) Schema() string {

	sql := `CREATE TABLE %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		created TEXT
	);

	CREATE UNIQUE INDEX `%s_by_list` ON `lists` (`name`);`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name())
}

func (t *ListsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *ListsTable) IndexRecord(db sqlite.Database, i interface{}) error {

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
