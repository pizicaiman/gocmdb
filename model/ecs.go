package model

// Ecs 结构体
type Ecs struct {
	Uid       int    `json:"uid"`
	Ecsname   string `json:"ecsname"`
	Ecsarea   string `json:"ecsarea"`
	Ecsstatus string `json:"ecsstatus"`
	Ecssize   string `json:"ecssize"`
	Ecsos     string `json:"ecsos"`
	Ecsip     string `json:"ecsip"`
	Ecsdate   string `json:"ecsdate"`
}
