package privatezen

import (
       "github.com/whosonfirst/go-whosonfirst-sqlite"
)

type PrivatezenTable interface {
     sqlite.Table
     // IndexBrand(sqlite.Database, wof_brands.Brand) error
}

