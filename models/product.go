package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Product struct {

	gorm.Model

	Name string /*名称*/
	Title string /*标题*/
	Description string /*描述*/
	Picture string /*主要图片*/

	Price float64 /*原价 RMB*/
	DiscountPrice float64/*折后价*/
	DiscountPercent float64/*打折力度0-1*/

	ProductionDate time.Time/*生产日期*/
	ExpiringDate time.Time /*过期日期*/
	GuaranteePeriod string /*保质期*/

	//...更多

	CategoryID uint/*所属类目id*/

}

type RepertoryRecord struct {

	gorm.Model

	ProductID uint /*商品id*/

	ChangeNumber int /*变化数量*/
	Surplus int /*变化后剩余数量*/

	Remark string /*备注*/
	OperaterID uint /*操作人*/
}
