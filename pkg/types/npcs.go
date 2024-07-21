package types

type TaxCollectorName struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
}

type NpcAction struct {
	Id     int `json:"id"`
	RealId int `json:"realId"`
	NameId int `json:"nameId"`
}

type NpcMessage struct {
	Id            int      `json:"id"`
	MessageId     int      `json:"messageId"`
	MessageParams []string `json:"messageParams"`
}

type Npc struct {
	Id     int `json:"id"`
	NameId int `json:"nameId"`
	// dialogMessages vector subtype not implemented
	// dialogReplies vector subtype not implemented
	Actions []uint `json:"actions"`
	Gender  int    `json:"gender"`
	Look    string `json:"look"`
	// animFunList vector custom subtype not implemented (1)
	FastAnimsFun   bool `json:"fastAnimsFun"`
	TooltipVisible bool `json:"tooltipVisible"`
}

type AnimFunNpcData struct {
	AnimId     int    `json:"animId"`
	EntityId   int    `json:"entityId"`
	AnimName   string `json:"animName"`
	AnimWeight int    `json:"animWeight"`
	// subAnimFunData vector custom subtype not implemented (1)
}

type TaxCollectorFirstname struct {
	Id          int `json:"id"`
	FirstnameId int `json:"firstnameId"`
}
