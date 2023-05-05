package leveldb

import (
	"errors"

	lichee_db "github.com/issueye/lichee-db"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ErrNotFound = errors.New("db not found")
)

type DB struct {
	list map[string]*leveldb.DB // 数据库
}

func NewDB() lichee_db.DB {
	return &DB{
		list: make(map[string]*leveldb.DB),
	}
}

func (b *DB) Delete(name string) {
	delete(b.list, name)
}

func (db *DB) Create(name string) error {
	return nil
}

func (db *DB) Modify(old, new string) error {
	return nil
}

func (db *DB) SetPath(path string) {

}

func (db *DB) GetBucket(dbName, name string) lichee_db.Bucket {
	return nil
}
