package tables

import (
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-sqlite"
	"github.com/whosonfirst/go-whosonfirst-sqlite/utils"
	"github.com/whosonfirst/go-whosonfirst-sqlite-privatezen"
	_ "log"
)

type VisitsTable struct {
	privatezen.PrivatezenTable
	name string
}

func NewVisitsTableWithDatabase(db sqlite.Database) (sqlite.Table, error) {

	t, err := NewVisitsTable()

	if err != nil {
		return nil, err
	}

	err = t.InitializeTable(db)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func NewVisitsTable() (sqlite.Table, error) {

	t := VisitsTable{
		name: "visits",
	}

	return &t, nil
}

func (t *VisitsTable) Name() string {
	return t.name
}

func (t *VisitsTable) Schema() string {

	sql := `CREATE TABLE %s (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	latitude NUMERIC,
	longitude NUMERIC,
	wof_id INTEGER,
	neighbourhood_id INTEGER,
	macrohood_id INTEGER,
	locality_id INTEGER,
	metroarea_id INTEGER,			
	region_id INTEGER,
	country_id INTEGER,
	feelings_id INTEGER,
	date TEXT,
	name TEXT
);

CREATE INDEX `%s_by_date` ON visits (`date`);
CREATE INDEX `%s_by_neighbourhood` ON visits (`neighbourhood_id`);
CREATE INDEX `%s_by_macrohood` ON visits (`macrohood_id`);
CREATE INDEX `%s_by_locality` ON visits (`locality_id`, `neighbourhood_id`);
CREATE INDEX `%s_by_metroarea` ON visits (`metroarea_id`);
CREATE INDEX `%s_by_locality_macro` ON visits (`locality_id`, `macrohood_id`);
CREATE INDEX `%s_by_region` ON visits (`region_id`);
CREATE INDEX `%s_by_country` ON visits (`country_id`);
`

	// this is a bit stupid really... (20170901/thisisaaronland)
	return fmt.Sprintf(sql, t.Name(), t.Name(), t.Name(), t.Name(), t.Name(), t.Name(), t.Name(), t.Name(), t.Name())
}

func (t *VisitsTable) InitializeTable(db sqlite.Database) error {

	return utils.CreateTableIfNecessary(db, t)
}

func (t *VisitsTable) IndexRecord(db sqlite.Database, i interface{}) error {

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
