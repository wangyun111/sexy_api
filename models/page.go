package models

type Page struct {
	PageNo     int64       `json:"pageNo"     description:"第几页"`
	PageSize   int64       `json:"pageSize"   description:"每页多少条"`
	TotalPage  int64       `json:"totalPage"  description:"总页数"`
	TotalCount int64       `json:"totalCount" description:"总条数"`
	FirstPage  bool        `json:"firstPage"  description:"是否是第一页"`
	LastPage   bool        `json:"lastPage"   description:"是否是最后一页"`
	List       interface{} `json:"list"       description:"内容list"`
}

type CountModel struct {
	Count int64
}

func PageUtil(count, pageNo, pageSize int64, list interface{}) *Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return &Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}
