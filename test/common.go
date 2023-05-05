package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"

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
