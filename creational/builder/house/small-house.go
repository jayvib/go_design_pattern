package house

type House struct {
	post   int
	window int
	door   int
	roof   int
}

type SmallHouse struct {
	house *House
}

func (s *SmallHouse) SetPost() HouseComponentStructurer {
	s.house.post = 4
	return s
}

func (s *SmallHouse) SetWindow() HouseComponentStructurer {
	s.house.window = 1
	return s
}

func (s *SmallHouse) SetDoor() HouseComponentStructurer {
	s.house.door = 1
	return s
}

func (s *SmallHouse) SetRoof() HouseComponentStructurer {
	s.house.roof = 1
	return s
}

func (s *SmallHouse) GetHouse() *House {
	return s.house
}
