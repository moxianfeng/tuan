package tools

type Goods struct {
	Name string
	Unit string
}

func NewGoods(name, unit string) *Goods {
	return &Goods{
		Name: name,
		Unit: unit,
	}
}
