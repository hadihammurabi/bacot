package block

type Meta struct {
	Index     int    `json:"index"`
	Timestamp int    `json:"timestamp"`
	PrevHash  string `json:"prev_hash"`
	Hash      string `json:"hash"`
	Nonce     int    `json:"nonce"`
}
