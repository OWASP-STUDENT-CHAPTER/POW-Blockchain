package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
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

type Message struct {
	Data int
}

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

func HandleWriteBlock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&m); err != nil {
		respondWithJSON(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	newBlock := generateBlock(Blockchain[len(Blockchain)-1], m.Data)
	mutex.Unlock()

	if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}

	respondWithJSON(w, r, http.StatusCreated, newBlock)

}

func respondWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}

func isBlockValid(newBlock, oldBlock model.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func generateBlock(oldBlock model.Block, Data int) model.Block {
	var newBlock model.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !isHashValid(calculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(calculateHash(newBlock), " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(calculateHash(newBlock), " work done!")
			newBlock.Hash = calculateHash(newBlock)
			break
		}

	}
	return newBlock
}
