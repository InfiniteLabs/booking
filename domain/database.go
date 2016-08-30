package domain


type Database []Room


func (db Database) GetRoomById(id int) Room {
	for _, v := range db {
		if v.id == id {
			return v
		}
	}

	return Room{}
} 

func GetDatabase() Database {
		return Database{
		{
			0,
			"room 0345",
			"some descripiton",
			[]TimeSlot{},
		},
		{
			1,
			"room 0455",
			"some other descripiton",
			[]TimeSlot{},
		},
	}
}