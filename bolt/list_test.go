package bolt

import (
	"fmt"
	lichee_db "issueye/lichee-db"
	"testing"
)

func NewTestList(name string) *List {
	q := make([][]byte, 0)
	q = append(q, []byte("测试数据1"))
	q = append(q, []byte("测试数据2"))
	q = append(q, []byte("测试数据3"))
	q = append(q, []byte("测试数据4"))
	q = append(q, []byte("测试数据5"))

	list := &List{
		name:  fmt.Sprintf("test:%s", name),
		queue: q,
	}
	return list
}

func printData(l *List) {
	fmt.Println("============================")
	for _, data := range l.queue {
		fmt.Println(string(data))
	}
	fmt.Println("============================")
}

func TestList_Insert(t *testing.T) {
	list := NewTestList("insert")
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
	list := NewTestList("lpop")
	printData(list)
	list.LPop()
	printData(list)
}

func TestList_LPush(t *testing.T) {
	list := NewTestList("lpush")
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
}

func TestList_Len(t *testing.T) {
	list := NewTestList("len")
	fmt.Println("len", list.Len())
}

func TestList_List(t *testing.T) {
	list := NewTestList("lpush")
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
}

func TestList_Move(t *testing.T) {
	list := NewTestList("move")
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
	list.Move(3, 1, lichee_db.Before)
	printData(list)
	list.Move(7, 5, lichee_db.After)
	printData(list)
}

func TestList_RPop(t *testing.T) {
	list := NewTestList("rpop")
	printData(list)
	list.LPush([]byte("lpush 测试数据1"), []byte("lpush 测试数据2"), []byte("lpush 测试数据3"))
	printData(list)
	list.RPop()
	printData(list)
}

func TestList_RPush(t *testing.T) {
	list := NewTestList("rpop")
	printData(list)
	list.RPush([]byte("Rpush 测试数据1"), []byte("Rpush 测试数据2"), []byte("Rpush 测试数据3"))
	printData(list)
}

func TestList_Remove(t *testing.T) {
	list := NewTestList("remove")
	printData(list)
	list.RPush([]byte("Rpush 测试数据1"), []byte("Rpush 测试数据2"), []byte("Rpush 测试数据3"))
	printData(list)
	list.Remove(2)
	printData(list)
}
