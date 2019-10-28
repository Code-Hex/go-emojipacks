package emojipacks

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type EmojiPacks struct {
	Title  string
	Emojis []struct {
		Name string
		Src  string
	}
}

func unmarshalYAMLs(paths []string) (packs []*EmojiPacks, _ error) {
	for _, path := range paths {
		pack, err := unmarshalYAML(path)
		if err != nil {
			return nil, err
		}
		packs = append(packs, pack)
	}
	return packs, nil
}

var open = defaultOpen

func unmarshalYAML(path string) (*EmojiPacks, error) {
	f, err := open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var typ EmojiPacks
	if err := yaml.NewDecoder(f).Decode(&typ); err != nil {
		return nil, err
	}
	return &typ, nil
}

func defaultOpen(path string) (io.ReadCloser, error) {
	return os.Open(path)
}
