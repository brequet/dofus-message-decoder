package types

type Playlist struct {
	Id   int `json:"id"`
	Type int `json:"type"`
	// sounds vector custom subtype not implemented (1)
	StartRandom       bool `json:"startRandom"`
	StartRandomOnce   bool `json:"startRandomOnce"`
	CrossfadeDuration int  `json:"crossfadeDuration"`
	Random            bool `json:"random"`
}
