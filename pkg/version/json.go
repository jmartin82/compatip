package version

import "github.com/tidwall/gjson"

type JsonPostProcessor struct {
	path string
}

func NewJsonPostProcessor(path string) *JsonPostProcessor {
	return &JsonPostProcessor{path: path}

}

func (jp *JsonPostProcessor) Process(in string) string {
	res := gjson.Get(in, jp.path)
	return res.String()

}
