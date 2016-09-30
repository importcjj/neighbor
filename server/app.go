package main

import (
	"github.com/importcjj/neighbor/actions"
	"github.com/kataras/iris"
	"net/http"
	"strconv"
)

func init() {
	iris.Get("/neighbors", MyNeighbors)
	iris.Get("/neighborhood", Neighborhood)
}

// MyNeighbors 接口: 我附近的人.
func MyNeighbors(ctx *iris.Context) {
	var userID int64 = 9
	neighbors, err := actions.NeighborOfUser(userID, nil)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(http.StatusOK, neighbors)
}

// Neighborhood 接口: 某地点附近的人.
func Neighborhood(ctx *iris.Context) {
	latitude := ctx.URLParam("lat")
	longitude := ctx.URLParam("lon")

	lat, err := strconv.ParseFloat(latitude, 64)
	lon, err := strconv.ParseFloat(longitude, 64)

	neighbors, err := actions.NeighborOfLocation(lat, lon, nil)
	if err != nil {
		ctx.WriteString(err.Error())
		return
	}
	ctx.JSON(http.StatusOK, neighbors)
}

func main() {
	iris.Listen(":8080")
}
