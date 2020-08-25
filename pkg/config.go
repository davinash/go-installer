package pkg

type InstallConfig struct {
	Components []Components           `json:"Components"`
	Services   []ServiceConfiguration `json:"Services"`
}

type Files struct {
	Name             string `json:"Name"`
	Platform         string `json:"Platform"`
	Action           string `json:"Action"`
	DecompressOutDir string `json:"DecompressOutDir"`
	Command          string `json:"Command"`
}
type Components struct {
	Name  string  `json:"Name"`
	Files []Files `json:"Files"`
}

type ServiceConfiguration struct {
	Name             string `json:"Name"`
	Description      string `json:"Description"`
	ShortDescription string `json:"ShortDescription"`
	Command          string `json:"Command"`
	Options          string `json:"Options"`
	Component        string `json:"Component"`
	InitScriptName   string `json:"InitScriptName"`
	LdLibraryPath    string `json:"LdLibraryPath"`
	DecompressOutDir string `json:"DecompressOutDir"`
}
