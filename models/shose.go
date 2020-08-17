package models

import "github.com/jinzhu/gorm"

type Shoes struct {
	Id              int64    `json:"id"  gorm:"primary_key" gorm:"column:id"`
	ProductId       int64    `json:"product_id" gorm:"column:product_id"`
	SoldNum         int64    `json:"sold_num" gorm:"column:sold_num"`
	DataType        int      `json:"data_type" gorm:"column:data_type"`
	PropertyValueId int      `json:"property_value_id" gorm:"column:data_type"`
	MinSalePrice    int64    `json:"min_sale_price" gorm:"column:min_sale_price"`
	PriceType       int      `json:"price_type" gorm:"column:price_type"`
	PropertyValue   string   `json:"property_value" gorm:"column:property_value"`
	Title           string   `json:"title" gorm:"column:title"`
	LogoUrl         string   `json:"logo_url" gorm:"column:logo_url"`
	GoodsType       int      `json:"goods_type" gorm:"column:goods_type"`
	BrandLogoUrl    string   `json:"brand_logo_url" gorm:"column:brand_logo_url"`
	SubTitle        string   `json:"sub_title" gorm:"column:sub_title"`
	ArticleNumber   string   `json:"article_number" gorm:"column:article_number"`
	RequestId       string   `json:"request_id" gorm:"column:request_id"`
	SpuMinSalePrice string   `json:"spu_min_sale_price" gorm:"column:spu_min_sale_price"`
	Price           int64    `json:"price" gorm:"column:price"`
	SpuId           int64    `json:"spu_id" gorm:"column:spu_id"`
	Page            int      `json:"page" gorm:"column:page"`
	Images          string   `json:"images" gorm:"column:images"`
	Session         *Session `json:"-" gorm:"-"`
}

// 设置Shoes的表名为`Shoes`
func (Shoes) TableName() string {
	return "shoes"
}

func NewShoesModel(session *Session) *Shoes {
	return &Shoes{Session: session}
}

// GetShoesTotal gets the total number of Shoes based on the constraints
func (a *Shoes) GetShoesTotal(maps interface{}) (int64, error) {
	var count int64
	err := a.Session.db.Model(&Shoes{}).Where(maps).Count(&count).Error
	return count, err
}

// GetShoes gets a list of Shoes based on paging constraints
func (a *Shoes) GetShoes(pageNum int, pageSize int, maps interface{}) ([]*Shoes, error) {
	var shoes []*Shoes
	err := a.Session.db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&shoes).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return shoes, nil
}
