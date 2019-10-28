package emojipacks

import (
	"errors"
	"io"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

const data = `---
title: codehex
emojis:
- name: codehex
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex.png
- name: codehex-black
  src: https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-black.png
`

var validEmojiPack = &EmojiPacks{
	Title: "codehex",
	Emojis: []struct {
		Name string
		Src  string
	}{
		{
			Name: "codehex",
			Src:  "https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex.png",
		},
		{
			Name: "codehex-black",
			Src:  "https://raw.githubusercontent.com/Code-Hex/code-hex.github.io/master/slack/assets/codehex-black.png",
		},
	},
}

func testOpen(path string) (io.ReadCloser, error) {
	rd := strings.NewReader(data)
	return ioutil.NopCloser(rd), nil
}

func Test_unmarshalYAMLs(t *testing.T) {
	tests := []struct {
		name      string
		openMock  func(path string) (io.ReadCloser, error)
		paths     []string
		wantPacks []*EmojiPacks
		wantErr   bool
	}{
		{
			name:     "valid",
			openMock: testOpen,
			paths: []string{
				"path1",
				"path2",
			},
			wantPacks: []*EmojiPacks{
				validEmojiPack,
				validEmojiPack,
			},
			wantErr: false,
		},
		{
			name: "invalid",
			openMock: func(path string) (io.ReadCloser, error) {
				return nil, errors.New("error")
			},
			paths:     []string{"path1"},
			wantPacks: nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open = tt.openMock
			defer func() {
				open = defaultOpen
			}()
			gotPacks, err := unmarshalYAMLs(tt.paths)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalYAMLs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotPacks, tt.wantPacks) {
				t.Errorf("unmarshalYAMLs() = %v, want %v", gotPacks, tt.wantPacks)
			}
		})
	}
}

func Test_unmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		openMock func(path string) (io.ReadCloser, error)
		path     string
		want     *EmojiPacks
		wantErr  bool
	}{
		{
			name:     "valid",
			openMock: testOpen,
			path:     "path",
			want:     validEmojiPack,
			wantErr:  false,
		},
		{
			name: "invalid open",
			openMock: func(path string) (io.ReadCloser, error) {
				return nil, errors.New("error")
			},
			path:    "path",
			want:    nil,
			wantErr: true,
		},
		{
			name: "invalid yaml",
			openMock: func(path string) (io.ReadCloser, error) {
				rd := strings.NewReader(`invalid`)
				return ioutil.NopCloser(rd), nil
			},
			path:    "path",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			open = tt.openMock
			defer func() {
				open = defaultOpen
			}()
			got, err := unmarshalYAML(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshalYAML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unmarshalYAML() = %v, want %v", got, tt.want)
			}
		})
	}
}
