package memdb

import (
	"log"

	"github.com/hashicorp/go-memdb"
)

type MemDB struct {
	db *memdb.MemDB
}

func NewMemDB() *MemDB {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			USERS_MALE_MEM_DB_TABLE_NAME:   getUsersMaleTableSchema(),
			USERS_FEMALE_MEM_DB_TABLE_NAME: getUsersFemaleTableSchema(),
		},
	}

	db, err := memdb.NewMemDB(schema)
	if err != nil {
		log.Fatalf("NewMemDB error: %v\n", err)
	}

	return &MemDB{
		db: db,
	}
}

func (m *MemDB) ExecTrx(write bool, fn func(*memdb.Txn) error) error {
	trx := m.db.Txn(true)

	if err := fn(trx); err != nil {
		trx.Abort()
		return err
	}

	trx.Commit()

	return nil
}
