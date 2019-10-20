package emojipacks

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	loginslack "github.com/Code-Hex/login-slack"
	"github.com/Songmu/prompter"
	"github.com/logrusorgru/aurora"
)

func Run(args []string) int {
	progName := args[0]
	if err := run(args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run %s: %v\n", progName, err)
		return 1
	}
	return 0
}

var _ flag.Value = (sliceFlags)(nil)

type sliceFlags []string

func (s sliceFlags) String() string {
	return strings.Join(s, ", ")
}

func (s sliceFlags) Set(v string) error {
	s = append(s, v)
	return nil
}

type options struct {
	subdomain string
	email     string
	password  string
	yamlfiles sliceFlags
}

func parse(args []string) (*options, error) {
	o := &options{}

	flg := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flg.StringVar(&o.subdomain, "s", "", "subdomain (workspace) of your slack")
	flg.StringVar(&o.email, "e", "", "email of your account")
	flg.StringVar(&o.password, "p", "", "password of your account")
	flg.Var(&o.yamlfiles, "y", "batch uploading of multiple yaml files")

	if err := flg.Parse(args[1:]); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *options) getSubDomain() string {
	if o.subdomain != "" {
		return o.subdomain
	}
	return prompter.Prompt("Slack subdomain", "")
}

func (o *options) getEmail() string {
	if o.email != "" {
		return o.email
	}
	return prompter.Prompt("Email address login", "")
}

func (o *options) getPassword() string {
	if o.password != "" {
		return o.password
	}
	return prompter.Password("Password")
}

func (o *options) getYAMLFiles() []string {
	if len(o.yamlfiles) > 0 {
		return o.yamlfiles
	}
	return []string{prompter.Prompt("Path or URL of Emoji yaml file", "")}
}

func run(args []string) error {
	o, err := parse(args)
	if err != nil {
		return err
	}
	subdomain := o.getSubDomain()
	email := o.getEmail()
	password := o.getPassword()
	yamlFiles := o.getYAMLFiles()

	packs, err := unmarshalYAMLs(yamlFiles)
	if err != nil {
		return err
	}

	accessToken, err := loginslack.Login(email, password, subdomain)
	if err != nil {
		return err
	}

	u := &uploader{accessToken: accessToken}

	return u.runUploadEmojiPacks(packs)
}

type uploader struct {
	accessToken string
}

func (u *uploader) runUploadEmojiPacks(packs []*EmojiPacks) error {
	for _, pack := range packs {
		if err := u.runUploadEmojiPack(pack); err != nil {
			return err
		}
	}
	return nil
}

func (u *uploader) runUploadEmojiPack(pack *EmojiPacks) error {
	fmt.Println(aurora.BrightCyan("start upload: " + pack.Title))
	for _, emoji := range pack.Emojis {
		err := u.pipeDownloadAndUpload(emoji.Src, emoji.Name)
		if err != nil {
			return err
		}
		fmt.Println(aurora.BrightGreen("complete: " + emoji.Name))
	}
	return nil
}

func (u *uploader) pipeDownloadAndUpload(srcURL, emojiName string) error {
	resp, err := http.Get(srcURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return u.uploadEmoji(resp.Body, resp.ContentLength, emojiName)
}
