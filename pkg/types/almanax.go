package types

type AlmanaxCalendar struct {
	Id         int   `json:"id"`
	NameId     int   `json:"nameId"`
	DescId     int   `json:"descId"`
	NpcId      int   `json:"npcId"`
	BonusesIds []int `json:"bonusesIds"`
}
