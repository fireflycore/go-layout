package base

type PageEntity struct {
	Page     uint `json:"page"`
	PageSize uint `json:"page_size"`
}

type ListEntity struct {
	Total uint        `json:"total"`
	List  interface{} `json:"list"`
}

func (p *PageEntity) Restrict() {
	if p.Page != 0 && p.Page < 1 {
		p.Page = 1
	}
	if p.PageSize > 100 {
		p.PageSize = 100
	}
	if p.PageSize == 0 || p.PageSize < 1 {
		p.PageSize = 1
	}
}
