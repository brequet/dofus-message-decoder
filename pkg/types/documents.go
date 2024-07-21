package types

type Document struct {
	Id                  int    `json:"id"`
	TypeId              int    `json:"typeId"`
	ShowTitle           bool   `json:"showTitle"`
	ShowBackgroundImage bool   `json:"showBackgroundImage"`
	TitleId             int    `json:"titleId"`
	AuthorId            int    `json:"authorId"`
	SubTitleId          int    `json:"subTitleId"`
	ContentId           int    `json:"contentId"`
	ContentCSS          string `json:"contentCSS"`
	ClientProperties    string `json:"clientProperties"`
}
