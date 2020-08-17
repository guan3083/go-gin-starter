package shoes_service

import "go-gin-starter/models"

type Shoes struct {
	Id              int64
	productId       int64
	soldNum         int64
	dataType        int
	propertyValueId int
	minSalePrice    int64
	priceType       int
	propertyValue   string
	title           string
	logoUrl         string
	goodsType       int
	brandLogoUrl    string
	subTitle        string
	articleNumber   string
	requestId       string
	spuMinSalePrice string
	price           int64
	spuId           int64
	page            int
	images          string

	PageNum  int
	PageSize int
}

func (s *Shoes) GetAll() ([]*models.Shoes, error) {
	session := models.NewSession()
	shoes, err := models.NewShoesModel(session).GetShoes(s.PageNum, s.PageSize, make(map[string]interface{}))
	return shoes, err
}

func (s *Shoes) GetTotals() (int64, error) {
	session := models.NewSession()
	total, err := models.NewShoesModel(session).GetShoesTotal(make(map[string]interface{}))
	return total, err
}
