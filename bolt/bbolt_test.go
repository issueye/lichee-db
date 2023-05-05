package bolt

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	lichee_db "github.com/issueye/lichee-db"
)

type PatientInfo struct {
	Code     string `son:"code"`
	Name     string `json:"name"`
	Age      string `json:"age"`
	Sex      string `json:"sex"`
	DeptCode string `json:"deptCode"`
	DeptName string `json:"deptName"`
	DocCode  string `json:"docCode"`
	DocName  string `json:"docName"`
	RoomCode string `json:"roomCode"`
	RoomName string `json:"roomName"`
}

func Test_Bolt(t *testing.T) {
	// 测试列表数据
	db := NewBbolt()
	t.Run("create", func(t *testing.T) {
		err := db.Create("db", "test")
		if err != nil {
			t.Errorf("创建数据库失败，失败原因：%s", err.Error())
		}
	})

	t.Run("bucket", func(t *testing.T) {

		wg := &sync.WaitGroup{}
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func(num int, wg *sync.WaitGroup) {
				key := fmt.Sprintf("test:bucket:00%d", num)
				b := db.GetBucket("test", key)
				l, err := b.List(key)
				if err != nil {
					t.Errorf("获取列表失败，失败原因：%s", err.Error())
				}

				for j := 0; j < 50; j++ {
					tb := &PatientInfo{
						Code:     fmt.Sprintf("测试病人-%d-%d", num, j),
						Name:     fmt.Sprintf("测试病人-%d-%d", num, j),
						Age:      "10",
						Sex:      "男",
						DeptCode: "0001",
						DeptName: "0001",
						DocCode:  "0001",
						DocName:  "0001",
						RoomCode: "0001",
						RoomName: "0001",
					}

					var buf bytes.Buffer
					err := gob.NewEncoder(&buf).Encode(tb)
					if err != nil {
						t.Errorf("序列化数据失败，失败原因：%s", err.Error())
					}
					l.LPush(buf.Bytes())
				}
				print(l)
				wg.Done()
			}(i, wg)
		}

		wg.Wait()
	})
}

func print(l lichee_db.List) {
	fmt.Println("============================")
	list := l.List()
	for _, data := range list {
		tb := &PatientInfo{}
		gob.NewDecoder(bytes.NewBuffer(data)).Decode(tb)
		name, _ := json.Marshal(tb)
		fmt.Println(string(name))
	}
	fmt.Println("============================")
}
