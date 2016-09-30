package lbs

import (
	"fmt"
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"github.com/gansidui/geohash"
)

const (
	defaultGeoHashPrecision = 6
	defaultPrecision        = 6
)

// GeoHashMatrix 返回九宫格中每个小矩形中点的geohash所组成的数组.
func GeoHashMatrix(latitude float64, longitude float64, precision ...int) []string {
	var _precision = defaultPrecision
	if len(precision) > 0 {
		_precision = precision[0]
	}
	neighbors := geohash.GetNeighbors(latitude, longitude, _precision)
	return neighbors
}

// GeoHash 计算给定经纬度所对应的geohash值.
func GeoHash(latitude float64, longitude float64, precision ...int) string {
	var _precision = defaultGeoHashPrecision
	if len(precision) > 0 {
		_precision = precision[0]
	}
	hash, _ := geohash.Encode(latitude, longitude, _precision)
	return hash
}

var geo = ellipsoid.Init(
	"WGS84",
	ellipsoid.Degrees,
	ellipsoid.Meter,
	ellipsoid.LongitudeIsSymmetric,
	ellipsoid.BearingIsSymmetric,
)

// Distance 计算两点之间的距离.
func Distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	distance, _ := geo.To(lat1, lon1, lat2, lon2)
	return distance
}

var distanceStandards = []float64{
	10, 50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000,
}

// DistanceHuman 提高距离的可读性.
func DistanceHuman(distance float64) string {
	var standard float64
	for _, standard = range distanceStandards {
		if distance < standard {
			return fmt.Sprintf("%v米以内", standard)
		}
	}
	return fmt.Sprintf("超过%v米", standard)
}
