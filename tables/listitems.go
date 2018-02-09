package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"github.com/whosonfirst/go-whosonfirst-sqlite-privatezen"
	_ "log"
)

type ListItemsTable struct {
	privatezen.PrivatezenTable
	name string
}

func NewListItemsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewListItemsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewListItemsTable() (sqlite.Table, error) {

	t := ListItemsTable{
		name: "list_items",
	}

	return &t, nil
}

func (t *ListItemsTable) Name() string {
	return t.name
}

func (t *ListItemsTable) Schema() string {

	sql := `CREATE TABLE %s (
		list_id INTEGER,
		wof_id INTEGER,
		position INTEGER
	);

	CREATE UNIQUE INDEX `%s_by_list_item` ON %s (`list_id`, `wof_id`);`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name(), t.Name())
}

func (t *ListItemsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *ListItemsTable) IndexRecord(db sqlite.Database, i interface{}) error {

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
