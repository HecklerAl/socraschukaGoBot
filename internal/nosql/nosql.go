package nosql

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"sync"
	"time"
)

type Data struct {
	ActualLink string `json:"actual_link"`
	LinkMode   string `json:"link_mode"`
	LastUse    string `json:"last_use`
	CountUses  int    `json:"count_uses"`
}

var (
	MapData map[string]Data
	mu      sync.RWMutex
)

func LoadFromDB() error {
	content, err := ioutil.ReadFile("db.json")
	if err != nil {
		return err
	}

	if len(content) == 0 {
		MapData = make(map[string]Data)
		return nil
	}

	err = json.Unmarshal(content, &MapData)
	if err != nil {
		return err
	}

	return nil
}

func Add(modified, actual, link_mode string) {
	mu.Lock()
	defer mu.Unlock()
	MapData[modified] = Data{
		ActualLink: actual,
		LinkMode:   link_mode,
		LastUse:    time.Now().String(),
		CountUses:  1}
}

func SaveData() error {
	mu.Lock()
	bytes, err := json.Marshal(MapData)
	if err != nil {
		mu.Unlock()
		return err
	}
	mu.Unlock()

	err = ioutil.WriteFile("db.json", bytes, fs.ModeAppend)
	if err != nil {
		return err
	}

	return nil
}

func GetInfo(modified string) (*Data, error) {
	mu.RLock()
	defer mu.RUnlock()
	if val, exist := MapData[modified]; exist {
		return &val, nil
	}
	return nil, nil
}

func GetLen() (int, error) {
	mu.RLock()
	defer mu.RUnlock()
	size := len(MapData)
	return size, nil
}

func GetActual(short, mode string) (*string, error) {
	mu.Lock()
	defer mu.Unlock()

	str := MapData[short].ActualLink
	if mode != MapData[short].LinkMode {
		return nil, nil
	}

	inf := Data{
		ActualLink: str,
		LinkMode:   MapData[short].LinkMode,
		LastUse:    time.Now().String(),
		CountUses:  MapData[short].CountUses + 1,
	}
	MapData[short] = inf

	return &str, nil
}

func IsBooked(modified string) (bool, error) {
	mu.RLock()
	defer mu.RUnlock()
	_, exist := MapData[modified]

	fmt.Println(modified, exist)
	return exist, nil
}
