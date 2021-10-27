package gssdk

import (
	"github.com/rosbit/gnet"
	"fmt"
	"net/http"
)

func AddIndexDoc(index string, doc map[string]interface{}) (docId string, err error) {
	url := fmt.Sprintf("%s/doc/%s", SearcherBaseUrl, index)
	return updateDoc(url, doc)
}

func UpdateIndexDoc(index string, doc map[string]interface{}) (docId string, err error) {
	url := fmt.Sprintf("%s/update/%s", SearcherBaseUrl, index)
	return updateDoc(url, doc)
}

func updateDoc(url string, doc map[string]interface{}) (docId string, err error) {
	var res struct {
		Code int
		Msg string
		Id string
	}
	if _, err = gnet.JSONCallJ(url, &res, gnet.Params(doc), gnet.M("PUT")); err != nil {
		return
	}
	if res.Code != http.StatusOK {
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
		return
	}
	docId = res.Id
	return
}

func DeleteIndexDoc(index string, docId string) (err error) {
	url := fmt.Sprintf("%s/doc/%s", SearcherBaseUrl, index)
	var res struct {
		Code int
		Msg string
	}
	if _, err = gnet.JSONCallJ(url, &res, gnet.Params(map[string]interface{}{"id":docId}), gnet.M("DELETE")); err != nil {
		return
	}
	if res.Code != http.StatusOK {
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
	}
	return
}
