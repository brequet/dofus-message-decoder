package types

type SoundAnimation struct {
	Id                 int    `json:"id"`
	Label              string `json:"label"`
	Name               string `json:"name"`
	Filename           string `json:"filename"`
	Volume             int    `json:"volume"`
	Rolloff            int    `json:"rolloff"`
	AutomationDuration int    `json:"automationDuration"`
	AutomationVolume   int    `json:"automationVolume"`
	AutomationFadeIn   int    `json:"automationFadeIn"`
	AutomationFadeOut  int    `json:"automationFadeOut"`
	NoCutSilence       bool   `json:"noCutSilence"`
	StartFrame         int    `json:"startFrame"`
}

type SoundUi struct {
	Id        int    `json:"id"`
	UiName    string `json:"uiName"`
	OpenFile  string `json:"openFile"`
	CloseFile string `json:"closeFile"`
	// subElements vector subtype not implemented
}

type SoundUiElement struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	HookId int    `json:"hookId"`
	File   string `json:"file"`
}

type SoundUiHook struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type SoundBones struct {
	Id   int      `json:"id"`
	Keys []string `json:"keys"`
	// values vector subtype not implemented
}
