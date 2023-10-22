package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/KulwinderSingh07/POW-Blockchain/model"
	"github.com/davecgh/go-spew/spew"
)

var NewBlock model.Block

var Blockchain []model.Block

var difficulty = 1

var mutex = &sync.Mutex{}

func Blockinitalizer() {
	t := time.Now()
	genesisBlock := model.Block{}
	genesisBlock = model.Block{
		Index:      0,
		Timestamp:  t.String(),
		Data:       0,
		Hash:       calculateHash(genesisBlock),
		PrevHash:   "",
		Difficulty: difficulty,
		Nonce:      "",
	}
	spew.Dump(genesisBlock)

	mutex.Lock()
	Blockchain = append(Blockchain, genesisBlock)
	mutex.Unlock()
}

func calculateHash(block model.Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.Data) + block.PrevHash + block.Nonce
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func HandleGetBlockchain(res http.ResponseWriter, req *http.Request) {
	bytes, err := json.MarshalIndent(Blockchain, "", " ")
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	io.WriteString(res, string(bytes))
}

func HandleWriteBlock(res http.ResponseWriter, req *http.Request) {

}
