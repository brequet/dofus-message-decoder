package types

type ArenaLeagueSeason struct {
	Uid          int     `json:"uid"`
	NameId       int     `json:"nameId"`
	Beginning    float64 `json:"beginning"`
	Closure      float64 `json:"closure"`
	ResetDate    float64 `json:"resetDate"`
	FlagObjectId int     `json:"flagObjectId"`
}

type ExpeditionSeason struct {
	Uid          int     `json:"uid"`
	NameId       int     `json:"nameId"`
	Beginning    float64 `json:"beginning"`
	Closure      float64 `json:"closure"`
	ResetDate    float64 `json:"resetDate"`
	FlagObjectId int     `json:"flagObjectId"`
}

type ServerSeason struct {
	Uid          int     `json:"uid"`
	NameId       int     `json:"nameId"`
	Beginning    float64 `json:"beginning"`
	Closure      float64 `json:"closure"`
	ResetDate    float64 `json:"resetDate"`
	FlagObjectId int     `json:"flagObjectId"`
}
