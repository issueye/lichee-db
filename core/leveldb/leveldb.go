package leveldb

import (
	"errors"
	"fmt"
	"path/filepath"

	lichee_db "github.com/issueye/lichee-db"
	"github.com/issueye/lichee-db/core"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	ErrNotFound = errors.New("db not found")
)

type DB struct {
	list       map[string]*leveldb.DB // 数据库
	path       string                 // 数据库存放路径
	bucketList map[string]string      // 前缀
}

func NewDB() lichee_db.DB {
	return &DB{
		list:       make(map[string]*leveldb.DB),
		bucketList: make(map[string]string),
	}
}

func (b *DB) Delete(name string) {
	delete(b.list, name)
}

func (db *DB) Create(name string) error {
	path := filepath.Join(db.path, name)
	ldb, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return err
	}
	// 保存数据库对象
	db.list[name] = ldb
	return nil
}

func (db *DB) Modify(old, new string) error {
	return nil
}

func (db *DB) SetPath(path string) {
	db.path = path
}

func (db *DB) GetBucket(dbName, name string) lichee_db.Bucket {
	ldb, ok := db.list[dbName]
	if !ok {
		db.Create(name)
		ldb = db.list[dbName]
	}

	return &Bucket{
		keys: make([]string, 0),
		Name: []byte(name),
		db:   ldb,
	}
}

// Bucket 数据桶
type Bucket struct {
	keys []string    // 所有的键
	Name []byte      // 名称
	db   *leveldb.DB //  数据库
}

// Keys 返回所有键
func (b *Bucket) Keys() []string {
	return b.keys
}

func (b *Bucket) getKey(k string) []byte {
	return []byte(fmt.Sprintf("%s:%s", b.Name, k))
}

// 返回列表对象
func (b *Bucket) List(name string) (lichee_db.List, error) {
	value, err := b.db.Get(b.getKey(name), nil)
	if err != nil {
		if err != leveldb.ErrNotFound {
			return nil, err
		}
	}

	return core.NewList(name, value, func(l *core.List) error {
		//序列化
		data, err := core.MarshalList(l)
		if err != nil {
			return err
		}

		return b.db.Put(b.getKey(name), data, nil)
	})
}

func (b *Bucket) Key(name string) lichee_db.String {
	value, err := b.db.Get(b.getKey(name), nil)
	if err != nil {
		return nil
	}
	return core.Str(value)
}

func (b *Bucket) Hash(name string) lichee_db.Hash {
	//TODO implement me
	panic("implement me")
}

func (b *Bucket) Set(name string) lichee_db.Set {
	//TODO implement me
	panic("implement me")
}
