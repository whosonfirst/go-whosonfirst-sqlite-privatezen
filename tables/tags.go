package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"github.com/whosonfirst/go-whosonfirst-sqlite-privatezen"
	_ "log"
)

type TagsTable struct {
	privatezen.PrivatezenTable
	name string
}

func NewTagsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewTagsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewTagsTable() (sqlite.Table, error) {

	t := TagsTable{
		name: "tags",
	}

	return &t, nil
}

func (t *TagsTable) Name() string {
	return t.name
}

func (t *TagsTable) Schema() string {

	sql := `CREATE TABLE %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		tag TEXT
	);

	CREATE UNIQUE INDEX `%s_by_tag` ON tags (`tag`);`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name())
}

func (t *TagsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *TagsTable) IndexRecord(db sqlite.Database, i interface{}) error {

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
