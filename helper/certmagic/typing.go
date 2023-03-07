package certmagic

type ReqeustParam struct {
	Email     string
	Domain    string
	CaType    string
	Provider  string
	SecretId  string
	SecretKey string
	EabKeyId  string
	EabMacKey string
}

type Certificate struct {
	Names       []string
	NotAfter    int64
	NotBefore   int64
	Certificate [][]byte
	PrivateKey  []byte
	Issuer      map[string]any
}
