// data contains functions for reading the csv file of karmoy stickers and converts it
// to an array of KarmoySticker structs
package data

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)


// KarmoySticker contains the ID, location (Latitude and Longitude) and Description
// for one karmøy sticker.
type KarmoySticker struct {
  ID          uint16    `json:"id"`
  Latitude    float32   `json:"latitude"`
  Longitude   float32   `json:"longitude"`
  Description string    `json:"description"`
}


// GetData gets the data of karmøy stickers from a given csv file and returns
// a list of KarmoySticker structs.
func GetData(path string) ([]KarmoySticker, error) {
  f, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer f.Close()

  csvReader := csv.NewReader(f)
  csvReader.Comma = ';'

  data, err := csvReader.ReadAll()
  if err != nil {
    return nil, err
  }

  stickers := make([]KarmoySticker, len(data)-1)
  for i, line := range data[1:] {
    var sticker KarmoySticker

    id, err := strconv.ParseUint(line[1], 10, 16) 
    if err != nil {
      return nil, err
    }
    sticker.ID = uint16(id)
    
    latLon := strings.Split(line[0], " ")

    lon, err := strconv.ParseFloat(latLon[1][1:], 32)
    if err != nil {
      return nil, err
    }
    sticker.Longitude = float32(lon)

    lat, err := strconv.ParseFloat(latLon[2][:len(latLon[1])-2], 32)
    if err != nil {
      return nil, err
    }
    sticker.Latitude = float32(lat)
    
    sticker.Description = line[2]

    stickers[i] = sticker
  }


  return stickers, nil
}
