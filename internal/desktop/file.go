package desktop

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/var512/deskgen/internal/flags"
)

type File struct {
	Name      string
	Path      string
	Extension string
	FullPath  string
	Content   []byte
	Entry     Entry
}

func (f *File) parseContent() ([]byte, error) {
	const desktopEntry = `
{{- "[Desktop Entry]" -}}
{{ if .Entry.TypeKey }}
Type={{ .Entry.TypeKey }}
{{- end -}}
{{ if .Entry.Version }}
Version={{ .Entry.Version }}
{{- end -}}
{{ if .Entry.Name }}
Name={{ .Entry.Name }}
{{- end -}}
{{ if .Entry.GenericName }}
GenericName={{ .Entry.GenericName }}
{{- end -}}
{{ if flagIsSet "noDisplay" }}
NoDisplay={{ .Entry.NoDisplay }}
{{- end -}}
{{ if .Entry.Comment }}
Comment={{ .Entry.Comment }}
{{- end -}}
{{ if .Entry.Icon }}
Icon={{ .Entry.Icon }}
{{- end -}}
{{ if flagIsSet "hidden" }}
Hidden={{ .Entry.Hidden }}
{{- end -}}
{{ if .Entry.OnlyShowIn }}
OnlyShowIn={{ .Entry.OnlyShowIn }}
{{- end -}}
{{ if .Entry.NotShowIn }}
NotShowIn={{ .Entry.NotShowIn }}
{{- end -}}
{{ if flagIsSet "dbusActivatable" }}
DbusActivatable={{ .Entry.DbusActivatable }}
{{- end -}}
{{ if .Entry.TryExec }}
TryExec={{ .Entry.TryExec }}
{{- end -}}
{{ if .Entry.Exec }}
Exec={{ .Entry.Exec }}
{{- end -}}
{{ if .Entry.Path }}
Path={{ .Entry.Path }}
{{- end -}}
{{ if flagIsSet "terminal" }}
Terminal={{ .Entry.Terminal }}
{{- end -}}
{{ if .Entry.MimeType }}
MimeType={{ .Entry.MimeType }}
{{- end -}}
{{ if .Entry.Categories }}
Categories={{ .Entry.Categories }}
{{- end -}}
{{ if .Entry.Implements }}
Implements={{ .Entry.Implements }}
{{- end -}}
{{ if .Entry.Keywords }}
Keywords={{ .Entry.Keywords }}
{{- end -}}
{{ if flagIsSet "startupNotify" }}
StartupNotify={{ .Entry.StartupNotify }}
{{- end -}}
{{ if .Entry.StartupWMClass }}
StartupWMClass={{ .Entry.StartupWMClass }}
{{- end -}}
{{ if .Entry.URL }}
URL={{ .Entry.URL }}
{{- end -}}
{{ if flagIsSet "prefersNonDefaultGPU" }}
PrefersNonDefaultGPU={{ .Entry.PrefersNonDefaultGPU }}
{{- end -}}
{{ if .Entry.Actions }}
Actions={{ range $v := .Entry.Actions }}{{ $v.Name }};{{ end }}
{{- range $v := .Entry.Actions }}

[Desktop Action {{ $v.Name }}]
Name={{ $v.Name -}}
{{ if $v.Icon }}
Icon={{ $v.Icon }}
{{- end -}}
{{ if $v.Exec }}
Exec={{ $v.Exec }}
{{- end -}}
{{- end -}}
{{- end -}}
{{- "" }}
`
	buf := &bytes.Buffer{}

	tmpl, err := template.New("DesktopEntry").Funcs(template.FuncMap{"flagIsSet": flags.IsSet}).Parse(desktopEntry)
	if err != nil {
		return nil, err
	}

	err = tmpl.Execute(buf, f)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (f *File) Save() error {
	if _, err := os.Stat(f.FullPath); err == nil {
		return fmt.Errorf("file already exists: %v: %w", f.FullPath, err)
	}

	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		return fmt.Errorf("path error: %v: %w", f.Path, err)
	}

	err := ioutil.WriteFile(f.FullPath, f.Content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func NewFile(path, name string, entry Entry) (*File, error) {
	file := &File{
		Path:  path,
		Name:  name,
		Entry: entry,
	}

	// Set File.Extension.
	file.Extension = typeExtension[entry.TypeKey]

	// Set File.FullPath: use path or working directory.
	if file.Path == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		file.Path = wd
	}

	file.FullPath = file.Path + "/" + file.Name + "." + file.Extension

	// Set File.Content.
	content, err := file.parseContent()
	if err != nil {
		return nil, err
	}

	file.Content = content

	return file, nil
}
