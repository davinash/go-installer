package pkg

type InstallConfig struct {
	Components []Components `json:"Components"`
	Services   []Services   `json:"Services"`
}
type Action struct {
	Command string `json:"command"`
}
type Files struct {
	Name               string `json:"Name"`
	Platform           string `json:"Platform"`
	Action             Action `json:"Action"`
	SourceRelativePath string `json:"SourceRelativePath"`
}
type Components struct {
	Name      string  `json:"Name"`
	TargetDir string  `json:"TargetDir"`
	Files     []Files `json:"Files"`
}
type Linux struct {
	InitScriptName string `json:"InitScriptName"`
}
type Windows struct {
	ServiceName string `json:"ServiceName"`
}
type Services struct {
	Name             string  `json:"Name"`
	Description      string  `json:"Description"`
	ShortDescription string  `json:"ShortDescription"`
	Command          string  `json:"Command"`
	Options          string  `json:"Options"`
	Component        string  `json:"Component"`
	Linux            Linux   `json:"linux"`
	Windows          Windows `json:"windows"`
}
