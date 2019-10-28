package emojipacks

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type mockPrompter struct {
	PromptMock   func(message, defaultAnswer string) string
	PasswordMock func(message string) string
}

func (m mockPrompter) Prompt(message, defaultAnswer string) string {
	return m.PromptMock(message, defaultAnswer)
}

func (m mockPrompter) Password(message string) string {
	return m.PasswordMock(message)
}

func Test_options_getSubDomain(t *testing.T) {
	type fields struct {
		subdomain string
	}
	tests := []struct {
		name      string
		inputMock inputter
		fields    fields
		want      string
	}{
		{
			name: "option subdomain is empty",
			inputMock: mockPrompter{
				PromptMock: func(message, defaultAnswer string) string {
					return "code-hex"
				},
			},
			fields: fields{
				subdomain: "",
			},
			want: "code-hex",
		},
		{
			name:      "option subdomain has value",
			inputMock: defaultPrompter{},
			fields: fields{
				subdomain: "code-hex",
			},
			want: "code-hex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prompter = tt.inputMock
			defer func() {
				prompter = defaultPrompter{}
			}()
			o := &options{
				subdomain: tt.fields.subdomain,
			}
			if got := o.getSubDomain(); got != tt.want {
				t.Errorf("options.getSubDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_getEmail(t *testing.T) {
	type fields struct {
		email string
	}
	tests := []struct {
		name      string
		inputMock inputter
		fields    fields
		want      string
	}{
		{
			name: "option email is empty",
			inputMock: mockPrompter{
				PromptMock: func(message, defaultAnswer string) string {
					return "code-hex@codehex.dev"
				},
			},
			fields: fields{
				email: "",
			},
			want: "code-hex@codehex.dev",
		},
		{
			name:      "option email has value",
			inputMock: defaultPrompter{},
			fields: fields{
				email: "code-hex@codehex.dev",
			},
			want: "code-hex@codehex.dev",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prompter = tt.inputMock
			defer func() {
				prompter = defaultPrompter{}
			}()
			o := &options{
				email: tt.fields.email,
			}
			if got := o.getEmail(); got != tt.want {
				t.Errorf("options.getEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_getPassword(t *testing.T) {
	type fields struct {
		password string
	}
	tests := []struct {
		name      string
		inputMock inputter
		fields    fields
		want      string
	}{
		{
			name: "option password is empty",
			inputMock: mockPrompter{
				PasswordMock: func(message string) string {
					return "password"
				},
			},
			fields: fields{
				password: "",
			},
			want: "password",
		},
		{
			name:      "option password has value",
			inputMock: defaultPrompter{},
			fields: fields{
				password: "password",
			},
			want: "password",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prompter = tt.inputMock
			defer func() {
				prompter = defaultPrompter{}
			}()
			o := &options{
				password: tt.fields.password,
			}
			if got := o.getPassword(); got != tt.want {
				t.Errorf("options.getPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_options_getYAMLFiles(t *testing.T) {
	type fields struct {
		yamlfiles sliceFlags
	}
	tests := []struct {
		name      string
		inputMock inputter
		fields    fields
		want      []string
	}{
		{
			name: "option yaml is empty",
			inputMock: mockPrompter{
				PromptMock: func(message, defaultAnswer string) string {
					return "yamlfile.yaml"
				},
			},
			fields: fields{
				yamlfiles: nil,
			},
			want: []string{"yamlfile.yaml"},
		},
		{
			name:      "option yaml has values",
			inputMock: defaultPrompter{},
			fields: fields{
				yamlfiles: sliceFlags{
					"yamlfile1.yaml",
					"yamlfile2.yaml",
				},
			},
			want: []string{
				"yamlfile1.yaml",
				"yamlfile2.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prompter = tt.inputMock
			defer func() {
				prompter = defaultPrompter{}
			}()
			o := &options{
				yamlfiles: tt.fields.yamlfiles,
			}
			got := o.getYAMLFiles()
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("(-want, +got)\n%s", diff)
			}
		})
	}
}
