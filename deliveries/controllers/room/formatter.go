package room

import R "group-project2/entities/room"

type RequestRoom struct {
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Guest       uint   `json:"guest" form:"guest"`
	Bedroom     uint   `json:"bedroom" form:"bedroom"`
	HasWifi     bool   `json:"has_wifi" form:"has_wifi"`
	HasAc       bool   `json:"has_ac" form:"has_ac"`
	HasKitchen  bool   `json:"has_kitchen" form:"has_kitchen"`
	HasFp       bool   `json:"has_fp" form:"has_fp"`
	Longitude   string `json:"longitude" form:"longitude"`
	Latitude    string `json:"latitude" form:"latitude"`
	City        string `json:"city" form:"city"`
	Price       uint   `json:"price" form:"price"`
}

func (Req RequestRoom) ToEntityRoom(UserID uint) R.Rooms {
	return R.Rooms{
		Name:        Req.Name,
		Description: Req.Description,
		Guest:       Req.Guest,
		Bedroom:     Req.Bedroom,
		HasWifi:     Req.HasWifi,
		HasAc:       Req.HasAc,
		HasKitchen:  Req.HasKitchen,
		HasFp:       Req.HasFp,
		Longitude:   Req.Longitude,
		Latitude:    Req.Latitude,
		City:        Req.City,
		Price:       Req.Price,
		UserID:      UserID,
	}
}

type ResponseCreateRoom struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Guest       uint   `json:"guest"`
	Bedroom     uint   `json:"bedroom"`
	HasWifi     bool   `json:"has_wifi"`
	HasAc       bool   `json:"has_ac"`
	HasKitchen  bool   `json:"has_kitchen"`
	HasFp       bool   `json:"has_fp"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	City        string `json:"city"`
	Price       uint   `json:"price"`
	UserID      uint   `json:"user_id"`
}

func ToResponseCreateRoom(Room R.Rooms) ResponseCreateRoom {
	return ResponseCreateRoom{
		Name:        Room.Name,
		Description: Room.Description,
		Guest:       Room.Guest,
		Bedroom:     Room.Bedroom,
		HasWifi:     Room.HasWifi,
		HasAc:       Room.HasAc,
		HasKitchen:  Room.HasKitchen,
		HasFp:       Room.HasFp,
		Longitude:   Room.Longitude,
		Latitude:    Room.Latitude,
		City:        Room.City,
		Price:       Room.Price,
		UserID:      Room.UserID,
	}
}

type ResponseGetRoom struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Guest       uint   `json:"guest"`
	Bedroom     uint   `json:"bedroom"`
	HasWifi     bool   `json:"has_wifi"`
	HasAc       bool   `json:"has_ac"`
	HasKitchen  bool   `json:"has_kitchen"`
	HasFp       bool   `json:"has_fp"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	City        string `json:"city"`
	Price       uint   `json:"price"`
	UserID      uint   `json:"user_id"`
}

func ToResponseGetAllRooms(Rooms []R.Rooms) []ResponseGetRoom {
	Responses := make([]ResponseGetRoom, len(Rooms))

	for i := 0; i < len(Rooms); i++ {
		Responses[i].ID = Rooms[i].ID
		Responses[i].Name = Rooms[i].Name
		Responses[i].Description = Rooms[i].Description
		Responses[i].Guest = Rooms[i].Guest
		Responses[i].Bedroom = Rooms[i].Bedroom
		Responses[i].HasWifi = Rooms[i].HasWifi
		Responses[i].HasAc = Rooms[i].HasAc
		Responses[i].HasKitchen = Rooms[i].HasKitchen
		Responses[i].HasFp = Rooms[i].HasFp
		Responses[i].Longitude = Rooms[i].Longitude
		Responses[i].Latitude = Rooms[i].Latitude
		Responses[i].City = Rooms[i].City
		Responses[i].Price = Rooms[i].Price
		Responses[i].UserID = Rooms[i].UserID
	}

	return Responses
}

func ToResponseGetRoomByID(Room R.Rooms) ResponseGetRoom {
	return ResponseGetRoom{
		ID:          Room.ID,
		Name:        Room.Name,
		Description: Room.Description,
		Guest:       Room.Guest,
		Bedroom:     Room.Bedroom,
		HasWifi:     Room.HasWifi,
		HasAc:       Room.HasAc,
		HasKitchen:  Room.HasKitchen,
		HasFp:       Room.HasFp,
		Longitude:   Room.Longitude,
		Latitude:    Room.Latitude,
		City:        Room.City,
		Price:       Room.Price,
		UserID:      Room.UserID,
	}
}

func ToGetRoomsByUserID(Rooms []R.Rooms) []ResponseGetRoom {
	Responses := make([]ResponseGetRoom, len(Rooms))

	for i := 0; i < len(Rooms); i++ {
		Responses[i].ID = Rooms[i].ID
		Responses[i].Name = Rooms[i].Name
		Responses[i].Description = Rooms[i].Description
		Responses[i].Guest = Rooms[i].Guest
		Responses[i].Bedroom = Rooms[i].Bedroom
		Responses[i].HasWifi = Rooms[i].HasWifi
		Responses[i].HasAc = Rooms[i].HasAc
		Responses[i].HasKitchen = Rooms[i].HasKitchen
		Responses[i].HasFp = Rooms[i].HasFp
		Responses[i].Longitude = Rooms[i].Longitude
		Responses[i].Latitude = Rooms[i].Latitude
		Responses[i].City = Rooms[i].City
		Responses[i].Price = Rooms[i].Price
		Responses[i].UserID = Rooms[i].UserID
	}

	return Responses
}

func ToResponseGetRoomsByCity(Rooms []R.Rooms) []ResponseGetRoom {
	Responses := make([]ResponseGetRoom, len(Rooms))

	for i := 0; i < len(Rooms); i++ {
		Responses[i].ID = Rooms[i].ID
		Responses[i].Name = Rooms[i].Name
		Responses[i].Description = Rooms[i].Description
		Responses[i].Guest = Rooms[i].Guest
		Responses[i].Bedroom = Rooms[i].Bedroom
		Responses[i].HasWifi = Rooms[i].HasWifi
		Responses[i].HasAc = Rooms[i].HasAc
		Responses[i].HasKitchen = Rooms[i].HasKitchen
		Responses[i].HasFp = Rooms[i].HasFp
		Responses[i].Longitude = Rooms[i].Longitude
		Responses[i].Latitude = Rooms[i].Latitude
		Responses[i].City = Rooms[i].City
		Responses[i].Price = Rooms[i].Price
		Responses[i].UserID = Rooms[i].UserID
	}

	return Responses
}

type ResponseUpdate struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Guest       uint   `json:"guest"`
	Bedroom     uint   `json:"bedroom"`
	HasWifi     bool   `json:"has_wifi"`
	HasAc       bool   `json:"has_ac"`
	HasKitchen  bool   `json:"has_kitchen"`
	HasFp       bool   `json:"has_fp"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	City        string `json:"city"`
	Price       uint   `json:"price"`
	UserID      uint   `json:"user_id"`
}

func ToResponseUpdate(Room R.Rooms) ResponseUpdate {
	return ResponseUpdate{
		ID:          Room.ID,
		Name:        Room.Name,
		Description: Room.Description,
		Guest:       Room.Guest,
		Bedroom:     Room.Bedroom,
		HasWifi:     Room.HasWifi,
		HasAc:       Room.HasAc,
		HasKitchen:  Room.HasKitchen,
		HasFp:       Room.HasFp,
		Longitude:   Room.Longitude,
		Latitude:    Room.Latitude,
		City:        Room.City,
		Price:       Room.Price,
		UserID:      Room.UserID,
	}
}
