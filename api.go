package lichee_db

type DataOrder int

const (
	Before DataOrder = iota
	After
)

// DB interface
type DB interface {
	Create(path, name string) error       // 创建数据库
	GetBucket(dbName, name string) Bucket //  获取BUCKET
	Modify(old, new string) error         // 修改数据库名称
	Delete(name string)                   // 删除数据库
}

// Bucket interface
type Bucket interface {
	Keys() []string                 // 获取所有的键名
	List(name string) (List, error) // 列表
	Key(name string) String         // 普通键值数据
	Hash(name string) Hash          // K-V
	Set(name string) Set            // Set 集合
}

// List 列表
type List interface {
	List() [][]byte                                        // 列表数据
	LPush(datas ...[]byte)                                 // 左边压入元素
	RPush(datas ...[]byte)                                 // 右边压入元素
	LPop() []byte                                          // 左边弹出元素
	RPop() []byte                                          // 右边弹出元素
	Insert(index int, it DataOrder, datas ...[]byte) error // 插入一个元素
	Move(nowIndex, targetIndex int, t DataOrder) error     // 移动一个元素
	Remove(index int) error                                // 移除 指定位置的元素
	Len() int                                              // 长度
}

// Hash 键值对数据
type Hash interface {
	Get(key string) []byte       // 获取一条数据
	Set(key string, data []byte) // 写入一个数据
	Len() int64                  // 长度
	GetAll() map[string][]byte   // 获取所有数据
}

// Set 集合
type Set interface {
	Add(data []byte)      // 添加一个元素
	GetAll() [][]byte     // 获取所有元素
	Diff(s Set) [][]byte  // 比较元素差别
	Inner(s Set) [][]byte // 比较共同
	Pop() []byte          // 随机弹出一个元素
	Len() int64           // 长度
}

// 字符串
type String interface {
	Get(key string) string        // 获取数据
	Put(key string, value string) // 写入数据
}
