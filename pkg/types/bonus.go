package types

type QuestKamasBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          uint  `json:"type"`
}

type MountBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          uint  `json:"type"`
}

type Bonus struct {
	Id            int   `json:"id"`
	Type          uint  `json:"type"`
	Amount        int   `json:"amount"`
	CriterionsIds []int `json:"criterionsIds"`
}

type QuestXPBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          uint  `json:"type"`
}

type MonsterStarRateBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          uint  `json:"type"`
}

type MonsterXPBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          uint  `json:"type"`
}

type MonsterDropChanceBonus struct {
	Amount        int   `json:"amount"`
	Id            int   `json:"id"`
	CriterionsIds []int `json:"criterionsIds"`
	Type          uint  `json:"type"`
}
