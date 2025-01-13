package natsiostaff

// GetStaffInfoReq ...
type GetStaffInfoReq struct {
	Ids     []string `json:"ids"`
	Partner string   `json:"partner"`
}

// GetStaffInfoRes ...
type GetStaffInfoRes struct {
	Data []StaffInfoResponse `json:"data"`
}

// StaffInfoResponse ...
type StaffInfoResponse struct {
	ID         string   `json:"_id"`
	Name       string   `json:"name"`
	Avatar     string   `json:"avatar"`
	Permission []string `json:"permission"`
	IsRoot     bool     `json:"isRoot"`
	Partner    string   `json:"partner"`
}

// CheckPermissionReq ...
type CheckPermissionReq struct {
	Token      string   `json:"token"`
	Permission []string `json:"permission"`
}

// CheckPermissionRes ...
type CheckPermissionRes struct {
	IsAccess bool              `json:"isAccess"`
	Data     StaffInfoResponse `json:"data"`
}

// GetPartnerInfoByCodeReq ...
type GetPartnerInfoByCodeReq struct {
	Code []string `json:"codes"`
	Ids  []string `json:"ids"`
}

// GetPartnerInfoRes ...
type GetPartnerInfoRes struct {
	Data []Partner `json:"data"`
}

// Partner ...
type Partner struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	AccessKey   string `json:"accessKey"`
	SecretKey   string `json:"secretKey"`
	CallbackURL string `json:"callbackURL"`
}

// GetPartnerInfoByAccessKeyReq ...
type GetPartnerInfoByAccessKeyReq struct {
	AccessKey string `json:"accessKey"`
}

// GetPartnerInfoByAccessKeyRes ...
type GetPartnerInfoByAccessKeyRes struct {
	Data Partner `json:"data"`
}
