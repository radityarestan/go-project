package main

import (
	"fmt"
	"latihan-protobuf/model"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
)

var user1 = &model.User{
	Id:       "u001",
	Name:     "John",
	Password: "123456",
	Gender:   model.UserGender_Male,
}

var user2 = &model.User{
	Id:       "u0002",
	Name:     "Ndi",
	Password: "123567",
	Gender:   model.UserGender_FEMALE,
}

var userList = &model.UserList{
	List: []*model.User{
		user1,
		user2,
	},
}

var garage1 = &model.Garage{
	Id:   "g001",
	Name: "Kalimdor",
	Coordinate: &model.GarageCoordinate{
		Latitude:  23.2212847,
		Longitude: 53.22033123,
	},
}

var garageList = &model.GarageList{
	List: []*model.Garage{
		garage1,
	},
}

var garageListByUser = &model.GarageListByUser{
	List: map[string]*model.GarageList{
		user1.Id: garageList,
	},
}

func main() {

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", user1.String())

	// =========== as json
	bytes, err1 := protojson.Marshal(garageList)

	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As JSON\n       %s \n", bytes)

	// =========== as proto
	protoObject := new(model.GarageList)

	err2 := protojson.Unmarshal(bytes, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As Proto (string)\n       %v \n", protoObject.String())

}
