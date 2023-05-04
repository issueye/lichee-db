package bolt

import (
	"errors"
	"fmt"
	bolt "go.etcd.io/bbolt"
	"issueye/lichee-db"
	"path/filepath"
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
func (b *DB) Create(name string) error {
	db, err := bolt.Open(filepath.Join("db", fmt.Sprintf("%s.db", name)), 0666, nil)
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
func (b *DB) GetBucket(name string) lichee_db.Bucket {
	b.list[name].Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		return err
	})

	bucket := &Bucket{
		Name: []byte(name),
		keys: make([]string, 0),
		db:   b.list[name],
	}

	return bucket
}

// Bucket 数据桶
type Bucket struct {
	keys []string // 所有的键
	Name []byte   // 名称
	db   *bolt.DB //  数据库
}

func (b *Bucket) View(r lichee_db.ReadFn) error {
	err := r(b)
	return err
}

func (b *Bucket) Update(w lichee_db.WriteFn) error {

	err := w(b)
	return err
}

// Keys 返回所有键
func (b *Bucket) Keys() []string {
	return b.keys
}

func (b *Bucket) List(name string) lichee_db.List {
	return nil
}

func (b *Bucket) Key(name string) (data []byte) {
	_ = b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(b.Name)
		data = bucket.Get([]byte(name))
		return nil
	})

	return
}

func (b *Bucket) Hash(name string) lichee_db.Hash {
	//TODO implement me
	panic("implement me")
}

func (b *Bucket) Set(name string) lichee_db.Set {
	//TODO implement me
	panic("implement me")
}
