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
	Code    uint8           `json:"code"`
	Message string          `json:"message"`
	Data    ImgViewRespData `json:"data"`
}

type ImageUploadMeta struct {
	LocalFilePath string `json:"local_file_path"`
	FromClipboard bool   `json:"from_clipboard"`
	FileName      string `json:"file_name"`
}

type ImageListMeta struct {
	Page     int    `json:"page"`
	Password string `json:"password"`
	Path     string `json:"path"`
	PerPage  int    `json:"per_page"`
	Refresh  bool   `json:"refresh"`
}

type EnvConfig struct {
	ServerPort    int    `json:"port"`
	ImageViewApi  string `json:"image_view_api"`
	AlistUrl      string `json:"alist_url"`
	AlistPassword string `json:"alist_password"`
	AlistUser     string `json:"alist_user"`
	ServerUrl     string `json:"server_url"`
	AesKey        string `json:"aes_key"`
}
