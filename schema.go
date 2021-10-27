package gssdk

import (
	"github.com/rosbit/gnet"
	"os"
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)

func CreateSchema(index string, schemaFile string) (err error) {
	fp, e := os.Open(schemaFile)
	if e != nil {
		err = e
		return
	}
	defer fp.Close()

	var res struct {
		Code int
		Msg string
	}
	url := fmt.Sprintf("%s/schema/%s", SearcherBaseUrl, index)
	if _, err = gnet.JSONCallJ(url, &res, gnet.Params(fp)); err != nil {
		return
	}
	if res.Code != http.StatusOK {
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
		return
	}
	return
}

func DeleteSchema(index string) (err error) {
	url := fmt.Sprintf("%s/schema/%s", SearcherBaseUrl, index)
	var res struct {
		Code int
		Msg string
	}
	if _, err = gnet.HttpCallJ(url, &res, gnet.M("DELETE")); err != nil {
		return
	}
	if res.Code != http.StatusOK {
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
		return
	}
	return
}

func DumpSchema(index string) (schema []byte, err error) {
	url := fmt.Sprintf("%s/schema/%s", SearcherBaseUrl, index)
	status, _, resp, e := gnet.Http(url, gnet.DontReadRespBody())
	if e != nil {
		err = e
		return
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if status != http.StatusOK {
		var res struct {
			Code int
			Msg string
		}
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
		return
	}
	return io.ReadAll(resp.Body)
}

func RenameSchema(oldName, newName string) (err error) {
	url := fmt.Sprintf("%s/schema/%s/%s", SearcherBaseUrl, oldName, newName)
	var res struct {
		Code int
		Msg string
	}
	if _, err = gnet.HttpCallJ(url, &res, gnet.M("PUT")); err != nil {
		return
	}
	if res.Code != http.StatusOK {
		err = fmt.Errorf("code: %d, msg: %s", res.Code, res.Msg)
		return
	}
	return
}
