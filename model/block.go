package model

type Block struct {
	Index      int
	Timestamp  string
	Data       int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      string
}
