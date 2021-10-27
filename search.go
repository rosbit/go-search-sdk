package gssdk

import (
	"github.com/rosbit/gnet"
	"fmt"
	"log"
	// "os"
	"net/http"
)

func Search(index string, options ...Option) (docs []Doc, pagination *Pagination, err error) {
	q := getQueryOptions(options...)
	params := q.makeQuery()
	log.Printf("[info] params: %s\n", params)
	url := fmt.Sprintf("%s/search/%s", SearcherBaseUrl, index)

	var res struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
		Result struct {
			Timeout bool `json:"timeout"`
			Docs []Doc `json:"docs"`
			*Pagination `json:"pagination"`
		} `json:"result"`
	}
	if _, err = gnet.HttpCallJ(url, &res, gnet.Params(params)/*, gnet.BodyLogger(os.Stderr)*/); err != nil {
		return
	}
	if res.Code != http.StatusOK {
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
		return
	}
	docs = res.Result.Docs
	pagination = res.Result.Pagination
	return
}
