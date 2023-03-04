package common

type ImgViewParam struct {
	Path     string `json:"path"`
	Password string `json:"password"`
}

type ImgViewRespData struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Sign     string `json:"sign"`
	RawUrl   string `json:"raw_url"`
	Provider string `json:"provider"`
	Type     int8   `json:"type"`
}

type ImgViewResp struct {
	Code    int8            `json:"code"`
	Message string          `json:"message"`
	Data    ImgViewRespData `json:"data"`
}
