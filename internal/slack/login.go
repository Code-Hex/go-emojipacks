package slack

import (
	"bytes"
	"strings"

	"github.com/Songmu/prompter"
	"github.com/headzoo/surf/browser"
	"gopkg.in/headzoo/surf.v1"
)

var defaultUA = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Mobile Safari/537.36"

var workspaceURL = func(workspace string) string {
	return "https://" + workspace + ".slack.com"
}

type options struct {
	useragent        string
	twoFactorHandler func() string
}

// Option indicates functional options.
type Option func(*options)

// WithTwoFactorHandler indicates option to invoke passed handler
// when called two factor authentication.
// It should returns 6-digits string.
func WithTwoFactorHandler(f func() string) Option {
	return func(o *options) {
		if f != nil {
			o.twoFactorHandler = f
		} else {
			o.twoFactorHandler = defaultTwoFactorHandler()
		}
	}
}

// WithUserAgent indicates option to change user-agent
func WithUserAgent(ua string) Option {
	return func(o *options) {
		o.useragent = ua
	}
}

// Login gets api-token (for only login user) after logged-in.
func Login(email, password, workspace string, opts ...Option) (string, error) {
	o := getOptions(opts...)

	bow := surf.NewBrowser()
	bow.SetUserAgent(o.useragent)
	baseURL := workspaceURL(workspace)
	if err := bow.Open(baseURL); err != nil {
		return "", err
	}

	if err := basicAuth(bow, email, password); err != nil {
		return "", &OpError{
			Op:  "basic authentication",
			Err: err,
		}
	}

	body := bow.Body()

	token, err := readToken(body)
	if err != nil {
		return "", &OpError{
			Op:  "read token after basic authenticated",
			Err: err,
		}
	}
	if token != "" {
		return token, nil
	}

	// 2fa check
	if strings.Index(body, "Enter your authentication code") >= 0 {
		sixDigit := o.twoFactorHandler()
		if err := twoFactorAuth(bow, sixDigit); err != nil {
			return "", &OpError{
				Op:  "two factor authentication",
				Err: err,
			}
		}
		body = bow.Body()
	}

	token, err = readToken(body)
	if err != nil {
		return "", &OpError{
			Op:  "read token after two factor authenticated",
			Err: err,
		}
	}
	if token == "" {
		return "", &OpError{
			Op:  "read token after basic authenticated",
			Err: ErrSomeThingWrong,
		}
	}
	return token, nil
}

func readToken(body string) (string, error) {
	const apiTokenKey = `"api_token"`
	idx := strings.Index(body, apiTokenKey)
	if idx < 0 {
		return "", ErrTokenNotFound
	}
	idx += len(apiTokenKey)

	var (
		buf          bytes.Buffer
		readingToken bool // = false
	)
	for i := idx; i < len(body); i++ {
		switch c := body[i]; c {
		case '"':
			if readingToken {
				return buf.String(), nil
			}
			readingToken = !readingToken
		default:
			if readingToken {
				buf.WriteByte(c)
			}
		}
	}
	return "", ErrReadLimitExceeded
}

func basicAuth(bow *browser.Browser, email, password string) error {
	signinForm, err := bow.Form("#signin_form")
	if err != nil {
		return err
	}
	if err := signinForm.Input("email", email); err != nil {
		return err
	}
	if err := signinForm.Input("password", password); err != nil {
		return err
	}
	if err := signinForm.Submit(); err != nil {
		return err
	}
	return nil
}

func twoFactorAuth(bow *browser.Browser, sixDigit string) error {
	twoFactorForm, err := bow.Form("div.two_factor_signin > form")
	if err != nil {
		return err
	}
	if err := twoFactorForm.Input("2fa_code", sixDigit); err != nil {
		return err
	}
	if err := twoFactorForm.Submit(); err != nil {
		return err
	}
	return nil
}

func getOptions(opts ...Option) *options {
	o := &options{
		useragent:        defaultUA,
		twoFactorHandler: defaultTwoFactorHandler(),
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func defaultTwoFactorHandler() func() string {
	return func() string {
		return prompter.Password("Enter 2FA Code")
	}
}
