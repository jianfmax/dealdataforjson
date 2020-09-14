package dealdataforjson

import (
	"fmt"
	"testing"
)

func TestChangeData(t *testing.T) {
	data := Read("tmp.json")
	fmt.Println(string(data))
	indexMap := make(map[string]string)
	indexMap["value.go_value.times"] = "2006-01-02 15:04:05"
	indexMap["new_time"] = "2006-01-02 15:04:05"
	newData, err := ChangeDataFromIndex(data, indexMap, TimeToDefault)
	fmt.Println(string(newData))
	if err != nil {
		t.Error(err.Error())
	}
}
