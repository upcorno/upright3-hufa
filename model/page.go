package model

import (
	"fmt"
	"regexp"
	"strings"

	"xorm.io/xorm"
)

// Page 分页基本数据
type Page struct {
	PageIndex int      `json:"page_index" form:"page_index" query:"page_index" validate:"gt=0"` //分页页码
	ItemNum   int      `json:"item_num" form:"item_num" query:"item_num" validate:"gt=0"`       //分页大小
	Order     []string `json:"order_arr[]" form:"order_arr[]" query:"order_arr[]"`
}

//处理分页及排序逻辑,page如:map[item_num:[8] order[]:[asc-assent_count asc-reply_count] page_index:[1]]
func (page *Page) GetResults(sess *xorm.Session, modsPtr interface{}, condiBean ...interface{}) (*PageResult, error) {
	defer sess.Close()
	sess.Limit(page.ItemNum, (page.PageIndex-1)*page.ItemNum) //分页
	orders := map[string]bool{"asc": true, "desc": true}
	regexpCompiled := regexp.MustCompile(`^[a-z,A-Z,0-9,\.,_]+$`) //正则表达式
	for _, item := range page.Order {
		arr := strings.Split(item, "-")
		if len(arr) != 2 || !orders[arr[0]] {
			return nil, fmt.Errorf("排序指令不合法:%s", item)
		}
		if !regexpCompiled.MatchString(arr[1]) {
			//由于直接使用了用户输入,所以要注意sql注入问题
			return nil, fmt.Errorf("检测到不合法字符:%s", arr[1])
		}
		if arr[0] == "asc" {
			sess.Asc(arr[1])
		} else {
			sess.Desc(arr[1])
		}
	}
	count, err := sess.FindAndCount(modsPtr)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		modsPtr = &[0]int{}
	}
	return &PageResult{Rows: modsPtr, Total: count}, nil
}

type PageResult struct {
	Rows  interface{} `json:"rows"`
	Total int64       `json:"total"`
}
