package models

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

type Pagination struct {
    Limit        int         `json:"limit,omitempty;query:limit"`
    Page         int         `json:"page,omitempty;query:page"`
    Sort         string      `json:"sort,omitempty;query:sort"`
    TotalRows    int64       `json:"total_rows"`
    TotalPages   int         `json:"total_pages"`
    Rows         interface{} `json:"rows"`
    Context      interface{} `json:"context"`
}

type PaginationShort struct {
    Limit        int         `json:"limit,omitempty;query:limit"`
    Page         int         `json:"page,omitempty;query:page"`
    Sort         string      `json:"sort,omitempty;query:sort"`
    TotalRows    int64       `json:"total_rows"`
    TotalPages   int         `json:"total_pages"`
}

func (p *Pagination) GetOffset() int {
    return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
    if p.Limit == 0 {
        p.Limit = 10
    }
    return p.Limit
}

func (p *Pagination) GetPage() int {
    if p.Page == 0 {
        p.Page = 1
    }
    if p.Page > p.TotalPages {
      p.Page = p.TotalPages
    }
    return p.Page
}

func (p *Pagination) GetSort() string {
    if p.Sort == "" {
        p.Sort = "id desc"
    }
    return p.Sort
}
