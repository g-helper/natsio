package natsiostaff

// CheckPermissionReq ...
type CheckPermissionReq struct {
	Token      string   `json:"token"`
	Permission []string `json:"permission"`
}

// CheckPermissionRes ...
type CheckPermissionRes struct {
	IsAccess   bool     `json:"isAccess"`
	Permission []string `json:"permission"`
}

// GetPartnerInfoByCodeReq ...
type GetPartnerInfoByCodeReq struct {
	Code []string `json:"codes"`
}

// GetPartnerInfoRes ...
type GetPartnerInfoRes struct {
	Data []Partner `json:"data"`
}

// Partner ...
type Partner struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Code      string `json:"code"`
	AccessKey string `bson:"accessKey"`
	SecretKey string `bson:"secretKey"`
}

// GetPartnerInfoByAccessKeyReq ...
type GetPartnerInfoByAccessKeyReq struct {
	AccessKey string `json:"accessKey"`
}

// GetPartnerInfoByAccessKeyRes ...
type GetPartnerInfoByAccessKeyRes struct {
	Data Partner `json:"data"`
}
