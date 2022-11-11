package entities

// Category 商品分类, 分类示例
type Category struct {
	Id   string
	Name string
}

// TableName 表名
func (Category) TableName() string {
	return "categories"
}
