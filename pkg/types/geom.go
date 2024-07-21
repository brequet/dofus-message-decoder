package types

type Rectangle struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
