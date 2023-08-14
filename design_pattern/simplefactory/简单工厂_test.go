package simplefactory

import "fmt"

type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j jsonRuleConfigParser) Parse(data []byte) {
	fmt.Println("implement")
}

type yamlRuleConfigParser struct {
}

func (j yamlRuleConfigParser) Parse(data []byte) {
	fmt.Println("implement")
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
