package form

type CommonListForm struct {
	SortBy string `p:"sort_by"`
	SortWay string `p:"sort_way"`
	PageIndex int `p:"page_index"`
	EachPage int `p:"each_page"`
}

func GetCommonListFormData(data *CommonListForm) *CommonListForm{
	if data.SortBy =="" {
		data.SortBy = "id"
	}
	if data.SortWay =="" {
		data.SortWay = "desc"
	}
	if data.PageIndex ==0 {
		data.PageIndex = 1
	}
	if data.EachPage ==0 {
		data.EachPage = 20
	}
	return data
}
