package backend

import "github.com/gogf/gf/v2/frame/g"

type DateHeadReq struct {
	g.Meta `path:"/data/head" method:"get" tags:"数据data" summary:"数据大屏的头部卡片"`
}

type DateHeadRes struct {
	TodayOrderCount int `json:"today_order_count" desc:"今日订单总数"`
	DAU             int `json:"dau" desc:"今日日活"`
	ConversionRate  int `json:"conversion_rate" desc:"转化率"`
}
