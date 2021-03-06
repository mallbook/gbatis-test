package bean

import (
	"encoding/json"

	"github.com/mallbook/gbatis"
)

// Mall model mall
type Mall struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Avatar    string `json:"avatar"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
	Story     string `json:"story"`
}

// String serial mall object
func (m Mall) String() string {
	s, err := json.Marshal(m)
	if err != nil {
		return err.Error()
	}
	return string(s)
}

// NewMall create a mall object
func NewMall() *Mall {
	return &Mall{}
}

// Shop model shop
type Shop struct {
	ID      string `json:"id"`
	BrandID string `json:"brandID"`
}

// String serial a shop object
func (shop Shop) String() string {
	s, err := json.Marshal(shop)
	if err != nil {
		return err.Error()
	}
	return string(s)
}

// NewShop new a shop
func NewShop() *Shop {
	return &Shop{}
}

// Brand model brands
type Brand struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Story  string `json:"story"`
}

func (b Brand) String() string {
	s, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}
	return string(s)
}

// NewBrand create a brand object
func NewBrand() *Brand {
	return &Brand{}
}

func init() {
	gbatis.RegisterBean("bean.Mall", NewMall)
	gbatis.RegisterBean("bean.Shop", NewShop)
	gbatis.RegisterBean("bean.Brand", NewBrand)
}
