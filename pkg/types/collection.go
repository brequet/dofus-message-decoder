package types

type Collectable struct {
	EntityId int    `json:"entityId"`
	Name     string `json:"name"`
	TypeId   int    `json:"typeId"`
	GfxId    int    `json:"gfxId"`
	Order    int    `json:"order"`
	Rarity   int    `json:"rarity"`
}

type Collection struct {
	TypeId    int    `json:"typeId"`
	Name      string `json:"name"`
	Criterion string `json:"criterion"`
	// collectables vector custom subtype not implemented (1)
}
