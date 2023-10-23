package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/KulwinderSingh07/POW-Blockchain/model"
	"github.com/KulwinderSingh07/POW-Blockchain/utils"
	"github.com/davecgh/go-spew/spew"
)

var NewBlock model.Block

var Blockchain []model.Block

var Difficulty int = 1

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
		Hash:       utils.CalculateHash(genesisBlock),
		PrevHash:   "",
		Difficulty: Difficulty,
		Nonce:      "",
	}
	spew.Dump(genesisBlock)

	mutex.Lock()
	Blockchain = append(Blockchain, genesisBlock)
	mutex.Unlock()
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
	res.Header().Set("Content-Type", "application/json")
	var m Message

	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&m); err != nil {
		utils.RespondWithJSON(res, req, http.StatusBadRequest, req.Body)
		return
	}
	defer req.Body.Close()

	mutex.Lock()
	newBlock := GenerateBlock(Blockchain[len(Blockchain)-1], m.Data)
	mutex.Unlock()

	if utils.IsBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
		Blockchain = append(Blockchain, newBlock)
		spew.Dump(Blockchain)
	}

	utils.RespondWithJSON(res, req, http.StatusCreated, newBlock)

}

func GenerateBlock(oldBlock model.Block, Data int) model.Block {
	var newBlock model.Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = Difficulty

	for i := 0; ; i++ {
		hex := fmt.Sprintf("%x", i)
		newBlock.Nonce = hex
		if !utils.IsHashValid(utils.CalculateHash(newBlock), newBlock.Difficulty) {
			fmt.Println(utils.CalculateHash(newBlock), " do more work!")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(utils.CalculateHash(newBlock), " work done!")
			newBlock.Hash = utils.CalculateHash(newBlock)
			break
		}

	}
	return newBlock
}
