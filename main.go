package main

import ("fmt"
		"file-compressor/mapper")

func main() {
	maps, _ := mapper.MapBytes("")
	fmt.Println(maps)
}