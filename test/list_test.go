package test

import (
	"fmt"
	"testing"

	lichee_db "github.com/issueye/lichee-db"
	"github.com/issueye/lichee-db/core"
)

func NewTestList(name string, t *testing.T) *core.List {
	list, _ := core.NewList(fmt.Sprintf("test:%s", name), nil, nil)

	list.LPush(
		[]byte("测试数据-0001"),
		[]byte("测试数据-0002"),
		[]byte("测试数据-0003"),
		[]byte("测试数据-0004"),
		[]byte("测试数据-0005"),
	)

	return list
}

func printData(l lichee_db.List) {
	fmt.Println("============================")
	list := l.List()
	for _, data := range list {
		fmt.Println(string(data))
	}
	fmt.Println("============================")
}

func TestList_Insert(t *testing.T) {
	list := NewTestList("insert", t)
	printData(list)

	t.Run("T=1", func(t *testing.T) {
		list.Insert(1, lichee_db.After, []byte("数据插入到第二条数据之后"))
	})

	printData(list)

	t.Run("T=0", func(t *testing.T) {
		list.Insert(2, lichee_db.Before, []byte("数据插入到第三条数据之前"))
	})

	printData(list)

}

func TestList_LPop(t *testing.T) {
	list := NewTestList("lpop", t)
	printData(list)
	list.LPop()
	printData(list)
}

func TestList_LPush(t *testing.T) {
	list := NewTestList("lpush", t)
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
}

func TestList_Len(t *testing.T) {
	list := NewTestList("len", t)
	fmt.Println("len", list.Len())
}

func TestList_List(t *testing.T) {
	list := NewTestList("lpush", t)
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
}

func TestList_Move(t *testing.T) {
	list := NewTestList("move", t)
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
	list.Move(3, 1, lichee_db.Before)
	printData(list)
	list.Move(7, 5, lichee_db.After)
	printData(list)
}

func TestList_RPop(t *testing.T) {
	list := NewTestList("rpop", t)
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
	list.RPop()
	printData(list)
}

func TestList_RPush(t *testing.T) {
	list := NewTestList("rpop", t)
	printData(list)
	list.RPush([]byte("Rpush 测试数据1"), []byte("Rpush 测试数据2"), []byte("Rpush 测试数据3"))
	printData(list)
}

func TestList_Remove(t *testing.T) {
	list := NewTestList("remove", t)
	printData(list)
	list.RPush([]byte("Rpush 测试数据1"), []byte("Rpush 测试数据2"), []byte("Rpush 测试数据3"))
	printData(list)
	list.Remove(2)
	printData(list)
}
