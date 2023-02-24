package certmagic

type Params struct {
	Email     string
	Domain    string
	CaType    string
	Provider  string
	SecretId  string
	SecretKey string
}

type Certificate struct {
	Names       []string
	OCSPStaple  []byte
	Certificate [][]byte
	PrivateKey  []byte
}
