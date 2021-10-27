package gssdk

import (
	"strings"
	"fmt"
)

type query struct {
	q string
	sorting []string
	f []string
	page, pageSize int
	fq []string
	fl []string
	pretty bool
}

type Option func(*query)
func getQueryOptions(options ...Option) *query {
    var q query
    for _, o := range options {
        o(&q)
    }
    return &q
}
func (q *query) makeQuery() map[string]interface{} {
	res := make(map[string]interface{})
	if len(q.q) > 0 {
		res["q"] = q.q
	}
	if len(q.sorting) > 0 {
		res["s"] = strings.Join(q.sorting, ",")
	}
	if len(q.f) > 0 {
		res["f"] = strings.Join(q.f, ";")
	}
	if q.page > 0 {
		res["page"] = q.page
	}
	if q.pageSize > 0 && q.pageSize <= 100 {
		res["pagesize"] = q.pageSize
	}
	if len(q.fq) > 0 {
		res["fq"] = strings.Join(q.fq, ",")
	}
	if len(q.fl) > 0 {
		res["fl"] = strings.Join(q.fl, ",")
	}
	if q.pretty {
		res["pretty"] = true
	}
	return res
}

func Q(qStr string) Option {
	return func(q *query) {
		q.q = qStr
	}
}

func Sorting(field string, isAsc ...bool) Option {
	return func(q *query) {
		scend := func()string{
			if len(isAsc) > 0 && isAsc[0]  {
				return "asc"
			}
			return "desc"
		}()
		q.sorting = append(q.sorting, fmt.Sprintf("%s:%s", field, scend))
	}
}

func Filter(field string, vals []string) Option {
	return func(q *query) {
		if len(vals) == 0 {
			return
		}
		q.f = append(q.f, fmt.Sprintf("%s:%s", field, strings.Join(vals, ",")))
	}
}

func PageNo(page int) Option {
	return func(q *query) {
		q.page = page
	}
}

func PageSize(pageSize int) Option {
	return func(q *query) {
		q.pageSize = pageSize
	}
}

func FieldQuery(field string, qStr string) Option {
	return func(q *query) {
		q.fq = append(q.fq, fmt.Sprintf("%s:%s", field, qStr))
	}
}

func OutputFields(fieldNames []string) Option {
	return func(q *query) {
		if len(fieldNames) > 0 {
			q.fl = append(q.fl, fieldNames...)
		}
	}
}

func WithPretty() Option {
	return func(q *query) {
		q.pretty = true
	}
}
