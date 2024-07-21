package types

type Notification struct {
	Id        int    `json:"id"`
	TitleId   int    `json:"titleId"`
	MessageId int    `json:"messageId"`
	IconId    int    `json:"iconId"`
	TypeId    int    `json:"typeId"`
	Trigger   string `json:"trigger"`
}
