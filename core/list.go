package core

import (
	"bytes"
	"encoding/gob"
	"errors"

	lichee_db "github.com/issueye/lichee-db"
)

var (
	ErrOutOfRange = errors.New("out of range")
)

type SaveDataFunc func(*List)

type List struct {
	name  string       // 名称
	queue [][]byte     // 数据
	fn    SaveDataFunc // 数据回调
}

func NewList(name string, data []byte, fn SaveDataFunc) (*List, error) {
	list, err := UnmarshalList(name, data, fn)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// UnmarshalList 从 bytes 反序列化到 List
func UnmarshalList(name string, data []byte, fn SaveDataFunc) (*List, error) {
	queue := make([][]byte, 0)

	list := new(List)
	list.name = name
	list.fn = fn

	if len(data) > 0 {
		// 反序列化到结构体
		err := gob.NewDecoder(bytes.NewBuffer(data)).Decode(&queue)
		if err != nil {
			return &List{}, err
		}
	}

	list.queue = queue

	return list, nil
}

// MarshalList 序列化 List 到 []byte
func MarshalList(list *List) ([]byte, error) {
	var buf bytes.Buffer

	// 创建 gob 编码器,对 list 进行序列化编码
	err := gob.NewEncoder(&buf).Encode(list.queue)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (l *List) callback() {
	if l.fn != nil {
		l.fn(l)
	}
}

func (l *List) List() [][]byte {
	return l.queue
}

func (l *List) LPush(datas ...[]byte) {
	l.queue = append(datas, l.queue...)
	l.callback()
}

func (l *List) RPush(datas ...[]byte) {
	l.queue = append(l.queue, datas...)
	l.callback()
}

func (l *List) LPop() []byte {
	if len(l.queue) == 0 {
		return nil
	}

	data := l.queue[0]
	l.queue = l.queue[1:]
	l.callback()
	return data
}

func (l *List) RPop() []byte {
	if len(l.queue) == 0 {
		return nil
	}

	data := l.queue[len(l.queue)-1]
	l.queue = l.queue[:len(l.queue)-1]
	l.callback()
	return data
}

func (l *List) Insert(index int, it lichee_db.DataOrder, datas ...[]byte) error {
	if index < 0 || index > len(l.queue) {
		return ErrOutOfRange
	}

	// 扩展队列容量
	l.queue = append(l.queue, make([][]byte, len(datas))...)

	// 后移元素
	copy(l.queue[index+len(datas):], l.queue[index:])

	// 插入新元素
	for i, data := range datas {
		if it == lichee_db.Before {
			l.queue[index+i] = data
		} else {
			l.queue[index+i+1] = data
		}
	}
	l.callback()
	return nil
}

func (l *List) Move(index, targetIndex int, it lichee_db.DataOrder) error {
	if index < 0 || index >= len(l.queue) ||
		targetIndex < 0 || targetIndex >= len(l.queue) {
		return ErrOutOfRange
	}

	// 取出要移动的元素
	data := l.queue[index]

	// 后移其他元素,在目标位置腾出空间
	if it == lichee_db.Before {
		copy(l.queue[targetIndex+1:], l.queue[targetIndex:index])
		l.queue[targetIndex] = data
	} else {
		copy(l.queue[targetIndex+2:], l.queue[targetIndex+1:index])
		l.queue[targetIndex+1] = data
	}

	l.callback()
	return nil
}

func (l *List) Remove(index int) error {
	if index < 0 || index >= len(l.queue) {
		return ErrOutOfRange
	}

	// 将后续元素前移,覆盖要删除的元素
	copy(l.queue[index:], l.queue[index+1:])
	l.queue = l.queue[:len(l.queue)-1]

	l.callback()
	return nil
}

func (l *List) Len() int {
	return len(l.queue)
}
