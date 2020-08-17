package block

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"go-gin-starter/models"
	"go-gin-starter/pkg/app"
	"go-gin-starter/pkg/e"
	"time"
)

var blockList []*models.Block

// @Summary block
// @Description
// @Tags block
// @accept json
// @Produce  json
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/block/list  [get]
func GetBlockList(c *gin.Context) {
	// 没有区块先初始化
	if len(blockList) <= 0 {
		newBlock := makeBlock(0, 0, "", "")
		// 日志打印newBlock 属性及值
		spew.Dump(newBlock)
		blockList = append(blockList, newBlock)
	}
	app.SuccessResp(c, blockList)
}

type ReqGenerateBlock struct {
	Bpm int `json:"bpm"`
}

// @Summary block
// @Description
// @Tags block
// @accept json
// @Produce  json
// @Param form body ReqGenerateBlock true "reqBody"
// @Success 200 {object}  app.Response
// @Failure 500 {object}  app.Response
// @Router /api/v1/block/generate  [post]
func GenerateBlock(c *gin.Context) {
	var (
		form ReqGenerateBlock
	)
	err := app.BindAndValid(c, &form)
	if err != nil {
		app.ErrorResp(c, e.INVALID_PARAMS, err.Error())
		return
	}
	if len(blockList) <= 0 {
		generateBlock := makeBlock(0, 0, "", "")
		blockList = append(blockList, generateBlock)
	}
	oldBlock := blockList[len(blockList)-1]
	newBlock := generate(oldBlock, form.Bpm)
	//检验区块
	if !validBlock(oldBlock, newBlock) {
		app.ErrorResp(c, e.BUSINESS_ERROR, "校验区块失败")
		return
	}
	app.SuccessResp(c, blockList)

}

func generate(oldBlock *models.Block, bpm int) *models.Block {
	newBlock := makeBlock(oldBlock.Index+1, bpm, "", oldBlock.Hash)
	newBlock.Hash = makeHash(newBlock)
	blockList = append(blockList, newBlock)
	return newBlock
}

func makeBlock(index, bpm int, hash, prevHash string) *models.Block {
	return &models.Block{
		Index:     index,
		Timestamp: time.Now().String(),
		BPM:       bpm,
		Hash:      hash,
		PrevHash:  prevHash,
	}
}

func makeHash(block *models.Block) string {
	item := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	hash := sha256.New()
	hash.Write([]byte(item))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func validBlock(oldBlock, newBlock *models.Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if makeHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}
