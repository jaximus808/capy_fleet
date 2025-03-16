package game

type Zone struct {
	id      int
	players map[int](*Player)
}

func CreateZone(id int) *Zone {
	return &Zone{
		id:      id,
		players: make(map[int]*Player),
	}
}
