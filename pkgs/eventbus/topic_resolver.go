package eventbus

import "fmt"

type TopicKey = string

type TopicResolver struct {
	KeyToName map[TopicKey]string
	NameToKey map[string]TopicKey
}

func NewTopicResolver(keyToName map[TopicKey]string) (*TopicResolver, error) {
	nameToKey := make(map[string]TopicKey)
	for k, name := range keyToName {
		if name == "" {
			return nil, fmt.Errorf("topic name is empty for key %q", k)
		}
		if _, dup := nameToKey[name]; dup {
			return nil, fmt.Errorf("duplicate topic name %q", name)
		}
		nameToKey[name] = k
	}
	return &TopicResolver{
		KeyToName: keyToName,
		NameToKey: nameToKey,
	}, nil
}

func (t *TopicResolver) GetName(key TopicKey) (string, bool) {
	v, ok := t.KeyToName[key]
	return v, ok
}

func (t *TopicResolver) GetKey(name string) (TopicKey, bool) {
	v, ok := t.NameToKey[name]
	return v, ok
}
