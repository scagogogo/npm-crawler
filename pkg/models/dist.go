package models

type Dist struct {
	Shasum     string       `json:"shasum"`
	Tarball    string       `json:"tarball"`
	Integrity  string       `json:"integrity"`
	Signatures []*Signature `json:"signatures"`
}

type Signature struct {
	Keyid string `json:"keyid"`
	Sig   string `json:"sig"`
}
