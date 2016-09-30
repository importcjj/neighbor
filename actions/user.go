package actions

import (
	"errors"

	"github.com/importcjj/neighbor/lbs"
	"github.com/importcjj/neighbor/models"
	"sort"
)

// SetUserlongitude 设置用户定位.
func SetUserlongitude(userID int64, latitude float64, longitude float64) error {
	db := models.DB()

	// Todo: check the user ID.

	location := new(models.Location)
	geohash := lbs.GeoHash(latitude, longitude)
	if db.Where("user_id = ?", userID).First(location).RecordNotFound() {
		location.UserID = userID
	}
	location.Latitude = latitude
	location.Longitude = longitude
	location.GeoHash = geohash
	db.Save(location)
	return nil
}

// NeighborFilter 附近的人的过滤选项.
type NeighborFilter struct {
	OnlyGirl    bool
	OnlyBoy     bool
	MaxDistance float64
}

// ValidDistance 检查最大距离.
func (filter NeighborFilter) ValidDistance(distance float64) bool {
	if filter.MaxDistance > 0 && distance > filter.MaxDistance {
		return false
	}
	return true
}

var defaultFilter = &NeighborFilter{
	MaxDistance: 1000,
}

// MgetUserByID 通过用户ID批量获取用户.
func MgetUserByID(userIDs []int64, filter *NeighborFilter) []*models.User {
	if filter == nil {
		filter = defaultFilter
	}
	db := models.DB()

	var users []*models.User
	query := db.Table(models.TableUser).Where("id IN (?)", userIDs)

	if filter.OnlyBoy {
		query = db.Where("sex = ?", "male")
	} else if filter.OnlyGirl {
		query = db.Where("sex = ?", "female")
	}

	query.Find(&users)

	return users
}

// Neighbor 附近的人.
type Neighbor struct {
	UserID        int64   `json:"user_id"`
	Username      string  `json:"username"`
	Sex           string  `json:"sex"`
	Distance      float64 `json:"distance"`
	DistanceHuman string  `json:"distance_human"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"Longitude"`
}

// Neighbors 附近的人.
type Neighbors []*Neighbor

func (n Neighbors) Len() int           { return len(n) }
func (n Neighbors) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n Neighbors) Less(i, j int) bool { return n[i].Distance < n[j].Distance }

// NeighborOfUser 用户附近的人.
func NeighborOfUser(userID int64, filter *NeighborFilter) ([]*Neighbor, error) {
	if filter == nil {
		filter = defaultFilter
	}
	db := models.DB()
	// Todo: check the user.

	// location of the user.
	location := new(models.Location)
	if db.Where("user_id = ?", userID).First(location).RecordNotFound() {
		return nil, errors.New("尚未设置位置信息")
	}

	return NeighborOfLocation(location.Latitude, location.Longitude, filter)
}

// NeighborOfLocation 获取某坐标点附近的人.
func NeighborOfLocation(lat float64, lon float64, filter *NeighborFilter) ([]*Neighbor, error) {
	if filter == nil {
		filter = defaultFilter
	}
	var neighbors = make([]*Neighbor, 0)
	// 附近的位置信息.
	geohashs := lbs.GeoHashMatrix(lat, lon)
	if len(geohashs) == 0 {
		return neighbors, errors.New("无效的经纬度坐标")
	}
	db := models.DB()
	// geohash := lbs.GeoHash(lat, lon, 4)
	var locations []*models.Location
	query := db.Where("geo_hash LIKE ?", geohashs[0]+"%")
	for _, h := range geohashs[1:] {
		query = query.Or("geo_hash LIKE ?", h+"%")
	}

	query.Find(&locations)
	if len(locations) == 0 {
		return neighbors, nil
	}

	// 附近的用户ID.
	var userIDs []int64
	var userLocationMap = make(map[int64]*models.Location)
	for _, locInfo := range locations {
		if locInfo.Latitude == lat && locInfo.Longitude == lon {
			continue
		}
		userIDs = append(userIDs, locInfo.UserID)
		userLocationMap[locInfo.UserID] = locInfo
	}
	if len(userIDs) == 0 {
		return neighbors, nil
	}

	// 附近的用户.
	users := MgetUserByID(userIDs, filter)
	for _, user := range users {
		userLatitude := userLocationMap[user.ID].Latitude
		userLongitude := userLocationMap[user.ID].Longitude
		distance := lbs.Distance(
			lat,
			lon,
			userLatitude,
			userLongitude,
		)
		if !filter.ValidDistance(distance) {
			continue
		}
		neighbor := &Neighbor{
			UserID:        user.ID,
			Username:      user.Username,
			Sex:           user.Sex,
			Distance:      distance,
			DistanceHuman: lbs.DistanceHuman(distance),
			Latitude:      userLatitude,
			Longitude:     userLongitude,
		}
		neighbors = append(neighbors, neighbor)
	}

	// 附近的用户按照距离远近排序.
	sort.Sort(Neighbors(neighbors))

	return neighbors, nil
}
