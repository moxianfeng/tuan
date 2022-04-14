package tools

type Shelf struct {
	Goods map[string]Goods
}

func NewShelf() *Shelf {
	return &Shelf{
		Goods: map[string]Goods{},
	}
}

func (shelf *Shelf) AddGoods(goods Goods) {
	if _, exists := shelf.Goods[goods.Name]; !exists {
		shelf.Goods[goods.Name] = goods
	}
}
