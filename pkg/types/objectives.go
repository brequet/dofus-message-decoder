package types

type QuestObjectiveParameters struct {
	NumParams   uint `json:"numParams"`
	Parameter0  int  `json:"parameter0"`
	Parameter1  int  `json:"parameter1"`
	Parameter2  int  `json:"parameter2"`
	Parameter3  int  `json:"parameter3"`
	Parameter4  int  `json:"parameter4"`
	DungeonOnly bool `json:"dungeonOnly"`
}

type QuestObjectiveFightMonstersOnMap struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveCraftItem struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveDiscoverMap struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveFightMonster struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveMultiFightMonster struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveBringItemToNpc struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveDuelSpecificPlayer struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveDiscoverSubArea struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveFreeForm struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveBringSoulToNpc struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}

type QuestObjectiveGoToNpc struct {
	StepId   int `json:"stepId"`
	TypeId   int `json:"typeId"`
	MapId    int `json:"mapId"`
	Id       int `json:"id"`
	DialogId int `json:"dialogId"`
	// parameters custom type not implemented (2)
	// coords custom type not implemented (1)
}
