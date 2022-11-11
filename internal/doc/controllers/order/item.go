package order

import (
	"github.com/gin-gonic/gin"
	"shequn1/foundation/app"
	"shequn1/foundation/database/category"
	"shequn1/foundation/database/mongo"
	"shequn1/foundation/database/orm"
	"shequn1/foundation/pager"
	"shequn1/internal/entities"
	"sync"
)

// Category 分类控制器, 示例
type Category struct {
	entity entities.Category
}

var once sync.Once

var tempData = []entities.Category{}

// NewCategory 实例化控制器
func NewCategory() *Category {
	orm.Master().AutoMigrate(entities.Category{})
	once.Do(func() {
		for _, data := range tempData {
			app.Logger().Debug(orm.Master().Create(&data).Error)
		}
		mongo.Collection(entities.Category{}).InsertMany(&tempData)
	})
	return &Category{entity: entities.Category{}}
}

// Mgo 使用 mgo
func (c *Category) Mgo(ctx *gin.Context) {
	app.NewResponse(app.Success, category.New().Table(c.entity).WithMgo().Categories()).End(ctx)
}

// Mongo 使用 mongo
func (c *Category) Mongo(ctx *gin.Context) {
	app.NewResponse(app.Success, category.New().Table(c.entity).WithMongo().Categories()).End(ctx)
}

// Mysql 使用 gorm
func (c *Category) Mysql(ctx *gin.Context) {
	app.NewResponse(app.Success, category.New().Table(c.entity).WithMysql().Categories()).End(ctx)
}

// ListMongo 分页功能
func (c *Category) ListMongo(ctx *gin.Context) {
	app.NewResponse(app.Success, pager.New(ctx, pager.NewMongoDriver()).SetIndex(c.entity.TableName()).Find(c.entity).Result()).End(ctx)
}

// ListMgo 分页功能
func (c *Category) ListMgo(ctx *gin.Context) {
	app.NewResponse(app.Success, pager.New(ctx, pager.NewMgoDriver()).SetIndex(c.entity.TableName()).Find(c.entity).Result()).End(ctx)
}

// ListMysql 分页功能
func (c *Category) ListMysql(ctx *gin.Context) {
	app.NewResponse(app.Success, pager.New(ctx, pager.NewGormDriver()).Where(pager.Where{"level": 2}).SetIndex(c.entity.TableName()).Find(c.entity).Result()).End(ctx)
}
