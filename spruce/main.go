package spruce

import (
	"github.com/geofffranks/simpleyaml"
	"github.com/geofffranks/yaml"
	"github.com/geofffranks/spruce"
)

func Merge(docs ...[]byte) ([]byte, error) {
	root := make(map[interface{}]interface{})
	m := &spruce.Merger{
		AppendByDefault: false,
	}
	for _, doc := range docs {
		parsed, err := simpleyaml.NewYaml(doc)
		if err != nil {
			return nil, err
		}
		doc, err := parsed.Map()
		if err != nil {
			return nil, err
		}
		m.Merge(root, doc)
	}
	return yaml.Marshal(root)
}