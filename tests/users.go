package main

import (
	"github.com/importcjj/neighbor/actions"
	"github.com/importcjj/neighbor/models"
)

func mockLocations() []*models.Location {
	locs := []*models.Location{
		{Latitude: 39.9046363143, Longitude: 116.4071136987},
		{Latitude: 39.9046363143, Longitude: 116.4071136988},
		{Latitude: 39.9046463143, Longitude: 116.4071135987},
		{Latitude: 39.9147363143, Longitude: 116.4071136987},
		{Latitude: 39.9097563143, Longitude: 116.4151236988},
		{Latitude: 39.9107363143, Longitude: 116.4071136987},
		{Latitude: 31.2519930642, Longitude: 121.3578674217},
		{Latitude: 31.2511440642, Longitude: 121.3534664217},
		{Latitude: 31.2502950642, Longitude: 121.3569154217},
		{Latitude: 31.2538910642, Longitude: 121.3560894217},
		{Latitude: 31.2444768021, Longitude: 121.3566589483},
		{Latitude: 31.2471628021, Longitude: 121.3486099483},
		{Latitude: 31.2470138780, Longitude: 121.3662039778},
		{Latitude: 31.2431488021, Longitude: 121.3505859483},
		{Latitude: 31.2421608021, Longitude: 121.3558319483},
		{Latitude: 31.2480578021, Longitude: 121.3512329483},
		{Latitude: 31.2438898021, Longitude: 121.3554729483},
		{Latitude: 31.2452798021, Longitude: 121.3545029483},
		{Latitude: 31.2520240642, Longitude: 121.3579214217},
		{Latitude: 31.2458658021, Longitude: 121.3547179483},
		{Latitude: 31.2483048021, Longitude: 121.3534969483},
		{Latitude: 31.2489378021, Longitude: 121.3584549483},
		{Latitude: 31.2441838021, Longitude: 121.3539819483},
		{Latitude: 31.2465758021, Longitude: 121.3553469483},
	}
	return locs
}

func mockUsers() []*models.User {
	users := []*models.User{
		{Username: "阿香", Sex: "Female"},
		{Username: "Jacky", Sex: "Male"},
		{Username: "老顾大业", Sex: "Male"},
	}

	return users
}

func addUsers() {
	db := models.DB()
	for _, user := range mockUsers() {
		db.Save(user)
	}
}

func addLocations() {
	for i, location := range mockLocations() {
		actions.SetUserlongitude(int64(i+1), location.Latitude, location.Longitude)
	}
}

func main() {
	addUsers()
	addLocations()
}
