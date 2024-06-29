// data contains functions for reading the csv file of karmoy stickers and converts it
// to an array of KarmoySticker structs
package data

import (
	"encoding/json"
	"os"
	"sync"
)


var fileMutex sync.RWMutex


// KarmoySticker contains the ID, location (Latitude and Longitude) and Description
// for one karmøy sticker.
type KarmoySticker struct {
  ID          uint16    `json:"id"`
  Latitude    float64   `json:"latitude"`
  Longitude   float64   `json:"longitude"`
  Description string    `json:"description"`
}


// GetData gets the data of karmøy stickers from a given csv file and returns
// a list of KarmoySticker structs.
func GetData(path string) ([]KarmoySticker, error) {
    fileMutex.RLock()
    defer fileMutex.RUnlock()

    data, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }

    var karmoyStickers []KarmoySticker
    if err := json.Unmarshal(data, &karmoyStickers); err != nil {
        return nil, err
    }

    return karmoyStickers, nil
}


func SaveData(path string, data []KarmoySticker) error {
    fileMutex.Lock()
    defer fileMutex.Unlock()

    dJson, err := json.Marshal(data)
    if err != nil {
        return err
    }

    return os.WriteFile(path, dJson, 0644)
}
