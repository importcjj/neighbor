package main

import (
	"fmt"
	"github.com/gansidui/geohash"
	_ "github.com/importcjj/neighbor/models"
)

func GeoHashDemo() {
	latitude := 39.92324
	longitude := 116.3906
	hashPrecision := 12
	precision := 9

	hash, box := geohash.Encode(latitude, longitude, hashPrecision)

	fmt.Println(hash)
	fmt.Println(box.MinLat, box.MaxLat, box.MinLng, box.MaxLng)

	neighbors := geohash.GetNeighbors(latitude, longitude, precision)
	for _, hash = range neighbors {
		fmt.Print(hash, " ")
	}
}

func main() {
	GeoHashDemo()
}
