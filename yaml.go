package emojipacks

import (
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

func unmarshalYAML(path string) (*EmojiPacks, error) {
	f, err := os.Open(path)
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

const data = `---
emojis:
- name: codehex
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex.png
- name: codehex-black
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-black.png
- name: codehex-face-jump
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-face-jump.gif
- name: codehex-ilust
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-ilust.jpg
- name: codehex-jump
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-jump.gif
- name: codehex-walking
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-walking.gif
- name: codehex-walking-2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-walking-2.gif
- name: codehex-walking-3
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-walking-3.gif
- name: codehex-walking-4
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-walking-4.gif
- name: codehex-white
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-white.png
- name: codehex1_1
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex1_1.png
- name: codehex1_2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex1_2.png
- name: codehex1_3
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex1_3.png
- name: codehex2_1
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex2_1.png
- name: codehex2_2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex2_2.png
- name: codehex2_3
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex2_3.png
- name: codehex3_1
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex3_1.png
- name: codehex3_2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex3_2.png
- name: codehex3_3
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex3_3.png
- name: codehex_body
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_body.png
- name: codehex_body_1
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_body_1.png
- name: codehex_body_2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_body_2.png
- name: codehex_body_3
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_body_3.png
- name: codehex_body_4
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_body_4.png
- name: codehex_dance
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_dance.gif
- name: codehex_face
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_face.png
- name: codehex_face_rotate_r
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_face_rotate_r.gif
- name: codehex_grass
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_grass.png
- name: codehex_grass_scroll
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_grass_scroll.gif
- name: codehex_hex
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_hex.png
- name: codehex_innocent
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_innocent.png
- name: codehex_long
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_long.png
- name: codehex_long2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_long2.png
- name: codehex_middle_1
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_middle_1.png
- name: codehex_middle_2
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_middle_2.png
- name: codehex_middle_3
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_middle_3.png
- name: codehex_middle_4
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_middle_4.png
- name: codehex_middle_long
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_middle_long.png
- name: codehex_party
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_party.gif
- name: codehex_partyparrot
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_partyparrot.gif
- name: codehex_scroll
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex_scroll.gif
title: codehex
`
