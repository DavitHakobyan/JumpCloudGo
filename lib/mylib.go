package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sync"
)

type Library struct {
	counter map[string][]int
	mutex sync.Mutex
}

type ActionTime struct {
	Action string
	Time   int
}

type ActionAvg struct {
	Action string
	Avg    int
}

func NewLibrary() Library {
	return Library{
		counter: make(map[string][]int),
		mutex: sync.Mutex{},
	}
}

func (lib Library) addAction(jsonString string) error {
	var actionTime ActionTime
	if err := json.Unmarshal([]byte(jsonString), &actionTime); err != nil {
		return err
	}

	lib.mutex.Lock()
	if array, ok := lib.counter[actionTime.Action]; ok == true {
		lib.counter[actionTime.Action] = append(array, actionTime.Time)
	} else {
		lib.counter[actionTime.Action] = []int{actionTime.Time}
	}
	lib.mutex.Unlock()
	return nil
}

func (lib Library) getStats() string {
	if len(lib.counter)==0 {
		return ""
	}

	lib.mutex.Lock()
	slice := make([]ActionAvg, 0)
	for action, list := range lib.counter {
		count := 0
		for _, time := range list {
			count += time
		}

		avg := ActionAvg{Action: action, Avg: count / len(list)}
		slice = append(slice, avg)
	}
	lib.mutex.Unlock()

	reqBodyBytes := new(bytes.Buffer)
	if err := json.NewEncoder(reqBodyBytes).Encode(slice); err != nil {
		fmt.Println(err)
	}
	resp := string(reqBodyBytes.Bytes())
	return resp
}

