package helper

type ResponseMessage struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseToken struct {
	Message     string   `json:"message"`
	AccessToken string   `json:"access_token"`
	Permission  []string `json:"permission"`
	ScopeID     string   `json:"scope_id"`
}

type ResponseTokenLogin struct {
	Status         bool   `json:"status"`
	Message        string `json:"message"`
	AccessToken    string `json:"access_token"`
	Organisasi     string `json:"organisasi"`
	UnitOrganisasi string `json:"unit_organisasi"`
	Uuid           string `json:"uuid"`
	Username       string `json:"username"`
	NamaPengguna   string `json:"nama_pengguna"`
}

type ResponseDataInfo struct {
	Status      bool   `json:"status"`
	Message     string `json:"message"`
	Data        any    `json:"data"`
	Information any    `json:"information"`
}
