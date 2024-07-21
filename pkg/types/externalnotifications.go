package types

type ExternalNotification struct {
	Id                  int    `json:"id"`
	CategoryId          int    `json:"categoryId"`
	IconId              int    `json:"iconId"`
	ColorId             int    `json:"colorId"`
	DescriptionId       int    `json:"descriptionId"`
	DefaultEnable       bool   `json:"defaultEnable"`
	DefaultSound        bool   `json:"defaultSound"`
	DefaultMultiAccount bool   `json:"defaultMultiAccount"`
	DefaultNotify       bool   `json:"defaultNotify"`
	Name                string `json:"name"`
	MessageId           int    `json:"messageId"`
}
