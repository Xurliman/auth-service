package pagination

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Pagination struct {
	Page       int                    `json:"page" query:"page"`
	Limit      int                    `json:"limit" query:"limit"`
	SortBy     string                 `json:"-" query:"sort_by"`
	SortDir    string                 `json:"-" query:"sort_dir"`
	Search     string                 `json:"-" query:"search"`
	TotalRows  int64                  `json:"total_rows"`
	TotalPages int                    `json:"total_pages"`
	Filters    map[string]interface{} `json:"-"`
	Rows       interface{}            `json:"-"`
}

type SearchParams struct {
	Search string
}

func GetSearchParams(ctx *fiber.Ctx) SearchParams {
	params := new(SearchParams)
	if err := ctx.QueryParser(params); err != nil {
		return SearchParams{}
	}
	return *params
}

func GetPaginationParams(ctx *fiber.Ctx) Pagination {
	params := new(Pagination)
	if err := ctx.QueryParser(params); err != nil {
		return Pagination{}
	}
	return *params
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
	return p.Page
}

func (p *Pagination) GetSort() string {
	sortDir := strings.ToUpper(p.SortDir)
	if p.SortBy == "" || (sortDir != "ASC" && sortDir != "DESC") {
		return ""
	}

	return p.SortBy + " " + sortDir
}
