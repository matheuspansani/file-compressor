package mapper

import (
	"file-compressor/reader"
)

func MapBytes(file string) (map[byte]int, error) {
    byteMap := make(map[byte]int)
    
    data, err := reader.ReadFile(file)
    if err != nil {
        return nil, err 
    }
    
    for _, b := range data {
        byteMap[b]++
    }
    
    return byteMap, nil
}