package types

type BreedRole struct {
	Id            int `json:"id"`
	NameId        int `json:"nameId"`
	DescriptionId int `json:"descriptionId"`
	AssetId       int `json:"assetId"`
	Color         int `json:"color"`
}

type BreedRoleByBreed struct {
	BreedId       int `json:"breedId"`
	RoleId        int `json:"roleId"`
	DescriptionId int `json:"descriptionId"`
	Value         int `json:"value"`
	Order         int `json:"order"`
}

type Breed struct {
	Id                         int    `json:"id"`
	ShortNameId                int    `json:"shortNameId"`
	LongNameId                 int    `json:"longNameId"`
	DescriptionId              int    `json:"descriptionId"`
	GameplayDescriptionId      int    `json:"gameplayDescriptionId"`
	GameplayClassDescriptionId int    `json:"gameplayClassDescriptionId"`
	MaleLook                   string `json:"maleLook"`
	FemaleLook                 string `json:"femaleLook"`
	CreatureBonesId            int    `json:"creatureBonesId"`
	MaleArtwork                int    `json:"maleArtwork"`
	FemaleArtwork              int    `json:"femaleArtwork"`
	// statsPointsForStrength vector subtype not implemented
	// statsPointsForIntelligence vector subtype not implemented
	// statsPointsForChance vector subtype not implemented
	// statsPointsForAgility vector subtype not implemented
	// statsPointsForVitality vector subtype not implemented
	// statsPointsForWisdom vector subtype not implemented
	BreedSpellsId []uint `json:"breedSpellsId"`
	// breedRoles vector custom subtype not implemented (1)
	MaleColors   []uint `json:"maleColors"`
	FemaleColors []uint `json:"femaleColors"`
	SpawnMap     int    `json:"spawnMap"`
	Complexity   int    `json:"complexity"`
	SortIndex    int    `json:"sortIndex"`
}

type Head struct {
	Id      int    `json:"id"`
	Skins   string `json:"skins"`
	AssetId string `json:"assetId"`
	Breed   int    `json:"breed"`
	Gender  int    `json:"gender"`
	Label   string `json:"label"`
	Order   int    `json:"order"`
}
