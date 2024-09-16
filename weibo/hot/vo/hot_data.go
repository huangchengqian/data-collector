package vo

import "time"

type Resp struct {
	HotGov   SovItem   `json:"hotgov"`
	RealTime []HotItem `json:"realtime"`
}

type HotItem struct {
	Word       string `json:"word"`
	Num        int64  `json:"num"`
	Rank       int64  `json:"rank"`
	Note       string `json:"note"`
	Period     int
	CreateTime time.Time
}

type SovItem struct {
	Note string `json:"note"`
	Word string `json:"word"`
	Num  int64  `json:"num"`
	Url  string `json:"url"`
}
