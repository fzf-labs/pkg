package repo

import (
	"bytes"
	"go/format"
	"os"
	"regexp"
	"text/template"

	"github.com/pkg/errors"
)

const regularPerm = 0o666

// DefaultTemplate is a tool to provides the text/template operations
type DefaultTemplate struct {
	name  string
	text  string
	goFmt bool
}

// NewTemplate returns an instance of DefaultTemplate
func NewTemplate(name string) *DefaultTemplate {
	return &DefaultTemplate{
		name: name,
	}
}

// Parse accepts a source template and returns DefaultTemplate
func (t *DefaultTemplate) Parse(text string) *DefaultTemplate {
	t.text = text
	return t
}

// GoFmt sets the value to goFmt and marks the generated codes will be formatted or not
func (t *DefaultTemplate) GoFmt(fmt bool) *DefaultTemplate {
	t.goFmt = fmt
	return t
}

// SaveTo writes the codes to the target path
func (t *DefaultTemplate) SaveTo(data any, path string, forceUpdate bool) error {
	if FileExists(path) && !forceUpdate {
		return nil
	}

	output, err := t.Execute(data)
	if err != nil {
		return err
	}

	return os.WriteFile(path, output.Bytes(), regularPerm)
}

// Execute returns the codes after the template executed
func (t *DefaultTemplate) Execute(data any) (*bytes.Buffer, error) {
	tem, err := template.New(t.name).Parse(t.text)
	if err != nil {
		return nil, errors.Wrapf(err, "template parse error:%s", t.text)
	}

	buf := new(bytes.Buffer)
	if err = tem.Execute(buf, data); err != nil {
		return nil, errors.Wrapf(err, "template execute error:%s", t.text)
	}

	if !t.goFmt {
		return buf, nil
	}

	formatOutput, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, errors.Wrapf(err, "go format error:%s", buf.String())
	}

	buf.Reset()
	buf.Write(formatOutput)
	return buf, nil
}

// IsTemplateVariable returns true if the text is a template variable.
// The text must start with a dot and be a valid template.
func IsTemplateVariable(text string) bool {
	match, _ := regexp.MatchString(`(?m)^{{(\.\w+)+}}$`, text)
	return match
}

// TemplateVariable returns the variable name of the template.
func TemplateVariable(text string) string {
	if IsTemplateVariable(text) {
		return text[3 : len(text)-2]
	}
	return ""
}