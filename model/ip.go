package model

// ip 结构体
type Ip struct {
	Uid        int    `json:"uid"`
	Ipname     string `json:"ipname"`
	Iptype     string `json:"iptype"`
	Ipbandwith string `json:"ipbandwith"`
	Ipbound    string `json:"ipbound"`
	Ipdate     string `json:"ipdate"`
}
