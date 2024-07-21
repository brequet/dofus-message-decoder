package types

type PlaylistSound struct {
	Id         string `json:"id"`
	Volume     int    `json:"volume"`
	SoundOrder int    `json:"soundOrder"`
}
