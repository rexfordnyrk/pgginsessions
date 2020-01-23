package pgginsessions

import (
	"database/sql"
	"github.com/rexfordnyrk/pgstore"
	"github.com/gin-contrib/sessions"
)

type Store interface {
	sessions.Store
}

func NewStore(dbURL string, keyPairs ...[]byte) (Store, error) {
	npgstore, err := pgstore.NewPGStore(dbURL, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{npgstore}, nil
}

func NewPGStoreWithSchema(dbURL string, schema string, keyPairs ...[]byte) (Store, error) {
	npgstore, err := pgstore.NewPGStoreWithSchema(dbURL, schema, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{npgstore}, nil
}

// NewPGStoreFromPool creates a new PGStore instance from an existing
// database/sql pool.
func NewPGStoreFromPool(db *sql.DB, keyPairs ...[]byte) (Store, error) {

	dbStore, err := pgstore.NewPGStoreFromPool(db, keyPairs...)
	if err != nil {
		return nil, err
	}

	return &store{dbStore}, nil
}

// NewPGStoreFromPool creates a new PGStore instance from an existing
// database/sql pool.
func NewPGStoreFromPoolWithSchema(db *sql.DB, schema string, keyPairs ...[]byte) (Store, error) {

	dbStore, err := pgstore.NewPGStoreFromPoolWithSchema(db, schema, keyPairs...)
	if err != nil {
		return nil, err
	}

	return &store{dbStore}, nil
}

type store struct {
	*pgstore.PGStore
}

func (c *store) Options(options sessions.Options) {
	c.PGStore.Options = options.ToGorillaOptions()
}
