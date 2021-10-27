package main

import (
	gssdk "github.com/rosbit/go-search-sdk"
	"flag"
	"os"
	"fmt"
	"strings"
	"encoding/json"
)

func init() {
	searcherUrl := os.Getenv("SEARCHER_URL")
	if len(searcherUrl) == 0 {
		panic("env SEARCHER_URL not found")
	}
	gssdk.SetSearcherBaseUrl(searcherUrl)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> <options>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Where <command> are:\n")
		fmt.Fprintf(os.Stderr, " create-schema\n")
		fmt.Fprintf(os.Stderr, " delete-schema, dump-schema\n")
		fmt.Fprintf(os.Stderr, " rename-schema\n")
		fmt.Fprintf(os.Stderr, " add-index-doc update-index-doc\n")
		fmt.Fprintf(os.Stderr, " delete-index-doc\n")
		fmt.Fprintf(os.Stderr, " search\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-schema":
		createSchema()
	case "delete-schema":
		deleteSchema()
	case "dump-schema":
		dumpSchema()
	case "rename-schema":
		renameSchema()
	case "add-index-doc":
		addIndexDoc()
	case "update-index-doc":
		updateIndexDoc()
	case "delete-index-doc":
		deleteIndexDoc()
	case "search":
		search()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %s\n", os.Args[1])
		os.Exit(2)
	}
}

func createSchema() {
	f := flag.NewFlagSet("create-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	file := f.String("schema-file", "", "specify schema-file")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*file) == 0 {
		fmt.Fprintf(os.Stderr, "%s create-schema -index=xxx -schema-file=xxx\n", os.Args[0])
		os.Exit(4)
	}
	err := gssdk.CreateSchema(*index, *file)
	fmt.Printf("err: %v\n", err)
}

func deleteSchema() {
	f := flag.NewFlagSet("delete-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 {
		fmt.Fprintf(os.Stderr, "%s delete-schema -index=xxx\n", os.Args[0])
		os.Exit(4)
	}
	err := gssdk.DeleteSchema(*index)
	fmt.Printf("err: %v\n", err)
}

func dumpSchema() {
	f := flag.NewFlagSet("dump-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 {
		fmt.Fprintf(os.Stderr, "%s dump-schema -index=xxx\n", os.Args[0])
		os.Exit(4)
	}
	schema, err := gssdk.DumpSchema(*index)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(5)
	}
	fmt.Printf("%s\n", schema)
}

func renameSchema() {
	f := flag.NewFlagSet("rename-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	newIndex := f.String("new-index", "", "specify new index name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*newIndex) == 0 {
		fmt.Fprintf(os.Stderr, "%s rename-schema -index=xxx -new-index=xxx\n", os.Args[0])
		os.Exit(4)
	}
	err := gssdk.RenameSchema(*index, *newIndex)
	fmt.Printf("err: %v\n", err)
}

func addIndexDoc() {
	f := flag.NewFlagSet("add-index-doc", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	doc := f.String("doc", "", "specify doc with A JSON object")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*doc) == 0 {
		fmt.Fprintf(os.Stderr, "%s add-index-doc -index=xxx -doc={JSON}\n", os.Args[0])
		os.Exit(4)
	}
	var d map[string]interface{}
	if err := json.Unmarshal([]byte(*doc), &d); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	docId, err := gssdk.AddIndexDoc(*index, d)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(6)
	}
	fmt.Printf("docId: %s\n", docId)
}

func updateIndexDoc() {
	f := flag.NewFlagSet("update-index-doc", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	doc := f.String("doc", "", "specify doc with A JSON object")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*doc) == 0 {
		fmt.Fprintf(os.Stderr, "%s update-index-doc -index=xxx -doc={JSON}\n", os.Args[0])
		os.Exit(4)
	}
	var d map[string]interface{}
	if err := json.Unmarshal([]byte(*doc), &d); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	docId, err := gssdk.UpdateIndexDoc(*index, d)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(6)
	}
	fmt.Printf("docId: %s\n", docId)
}

func deleteIndexDoc() {
	f := flag.NewFlagSet("delete-index-doc", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	docId := f.String("doc-id", "", "specify docId returned when calling add/update-index-doc")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*docId) == 0 {
		fmt.Fprintf(os.Stderr, "%s delete-index-doc -index=xxx -doc-id=xxx\n", os.Args[0])
		os.Exit(4)
	}
	err := gssdk.DeleteIndexDoc(*index, *docId)
	fmt.Printf("err: %v\n", err)
}

func search() {
	var options []gssdk.Option

	f := flag.NewFlagSet("search", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	f.Func("q", "specify query string", func(q string)error{
		if len(q) > 0 {
			options = append(options, gssdk.Q(q))
		}
		return nil
	})
	f.Func("s", "specify sorting, format: field[:asc]", func(s string)error{
		if len(s) == 0 {
			return nil
		}
		ss := strings.Split(s, ":")
		switch len(ss) {
		case 0: break
		case 1: options = append(options, gssdk.Sorting(ss[0]))
		default:
			scend := ss[1] == "asc"
			options = append(options, gssdk.Sorting(ss[0], scend))
		}
		return nil
	})
	page := f.Int("page", 0, "specify page")
	pageSize := f.Int("page-size", 0, "specify page size")
	f.Func("fq", "specify query in field, format: filed:q", func(fq string)error{
		if len(fq) == 0 {
			return nil
		}
		fqs := strings.Split(fq, ":")
		switch len(fqs) {
		case 2:
			options = append(options, gssdk.FieldQuery(fqs[0], fqs[1]))
		default:
		}
		return nil
	})
	f.Func("f", "specify field filter, format: field:val1,val2,...", func(filter string)error{
		if len(filter) == 0 {
			return nil
		}
		fs := strings.Split(filter, ":")
		switch len(fs) {
		case 2:
			vs := strings.FieldsFunc(fs[1], func(c rune)bool{
				return c == ',' || c == ' ' || c == '\t'
			})
			if len(vs) == 0 {
				return nil
			}
			options = append(options, gssdk.Filter(fs[0], vs))
		default:
		}
		return nil
	})
	f.Func("fl", "specify field list", func(fl string)error{
		if len(fl) == 0 {
			return nil
		}
		fls := strings.FieldsFunc(fl, func(c rune)bool{
			return c == ',' || c == ' ' || c == '\t'
		})
		if len(fls) == 0 {
			return nil
		}
		options = append(options, gssdk.OutputFields(fls))
		return nil
	})
	pretty := f.Bool("pretty", false, "specify pretty output")

	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 {
		fmt.Fprintf(os.Stderr, "%s search -index=xxx -q=xx -s=xx -page=xx -page-size=xx fq=xxx -fl=xxx\n", os.Args[0])
		os.Exit(4)
	}
	if *page > 0 {
		options = append(options, gssdk.PageNo(*page))
	}
	if *pageSize > 0 {
		options = append(options, gssdk.PageSize(*pageSize))
	}
	if *pretty {
		options = append(options, gssdk.WithPretty())
	}

	docs, pagination, err := gssdk.Search(*index, options...)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(5)
	}
	jEnc := json.NewEncoder(os.Stdout)
	jEnc.Encode(pagination)
	jEnc.Encode(docs)
}
