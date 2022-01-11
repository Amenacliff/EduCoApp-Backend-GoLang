package dto

type Project struct {
	ProjectName        string   `json:"projectName"`
	ProjectDescription string   `json:"projectDescription"`
	ProjectAssets      []string `json:"projectAssets"`
	TypeOfProject      string   `json:"typeOfProject"`
}
