package bolt

import (
	"errors"
	"fmt"
	"path/filepath"

	lichee_db "github.com/issueye/lichee-db"
	bolt "go.etcd.io/bbolt"
)

var (
	ErrNotFound = errors.New("db not found")
)

func NewBbolt() lichee_db.DB {
	return &DB{
		list: make(map[string]*bolt.DB),
	}
}

type DB struct {
	list map[string]*bolt.DB // 数据库
}

// Delete 删除数据库
func (b *DB) Delete(name string) {
	delete(b.list, name)
}

// Create 创建一个数据库对象
func (b *DB) Create(path, name string) error {
	fullPath := filepath.Join(path, fmt.Sprintf("%s.db", name))
	db, err := bolt.Open(fullPath, 0666, nil)
	if err != nil {
		return err
	}

	// 保存数据库对象
	b.list[name] = db
	return nil
}

// Modify 修改数据库的名称
func (b *DB) Modify(old, new string) error {
	db, ok := b.list[old]

	// 如果数据库不存在，则返回错误信息
	if !ok {
		return ErrNotFound
	}

	// 写到新的KEY
	b.list[new] = db
	// 移除原来的数据
	delete(b.list, old)
	return nil
}

// GetBucket 获取一个BUCKET
func (b *DB) GetBucket(dbName, name string) lichee_db.Bucket {
	b.list[dbName].Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})

	bucket := &Bucket{
		Name: []byte(name),
		keys: make([]string, 0),
		db:   b.list[dbName],
	}

	return bucket
}

// Bucket 数据桶
type Bucket struct {
	keys []string // 所有的键
	Name []byte   // 名称
	db   *bolt.DB //  数据库
}

// Keys 返回所有键
func (b *Bucket) Keys() []string {
	return b.keys
}

// 返回列表对象
func (b *Bucket) List(name string) (lichee_db.List, error) {
	var value []byte
	_ = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.Name)
		value = bucket.Get([]byte(name))
		return nil
	})

	return NewList(name, value, func(l *List) {
		b.db.Update(func(tx *bolt.Tx) error {
			bucket := tx.Bucket(b.Name)

			//序列化
			data, err := MarshalList(l)
			if err != nil {
				return err
			}

			// 写入数据
			return bucket.Put([]byte(name), data)
		})
	})
}

func (b *Bucket) Key(name string) lichee_db.String {
	var str []byte
	_ = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.Name)
		str = bucket.Get([]byte(name))
		return nil
	})
	return Str(str)
}

func (b *Bucket) Hash(name string) lichee_db.Hash {
	//TODO implement me
	panic("implement me")
}

func (b *Bucket) Set(name string) lichee_db.Set {
	//TODO implement me
	panic("implement me")
}
