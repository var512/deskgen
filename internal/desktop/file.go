package desktop

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type File struct {
	Name      string
	Path      string
	Extension string
	FullPath  string
	Content   []byte
	Entry     Entry
}

func (f File) parseContent() []byte {
	// TODO RTFD.
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
{{ if true }}
NoDisplay={{ .Entry.NoDisplay }}
{{- end -}}
{{ if .Entry.Comment }}
Comment={{ .Entry.Comment }}
{{- end -}}
{{ if .Entry.Icon }}
Icon={{ .Entry.Icon }}
{{- end -}}
{{ if true }}
Hidden={{ .Entry.Hidden }}
{{- end -}}
{{ if .Entry.OnlyShowIn }}
OnlyShowIn={{ .Entry.OnlyShowIn }}
{{- end -}}
{{ if .Entry.NotShowIn }}
NotShowIn={{ .Entry.NotShowIn }}
{{- end -}}
{{ if true }}
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
{{ if true }}
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
{{ if true }}
StartupNotify={{ .Entry.StartupNotify }}
{{- end -}}
{{ if .Entry.StartupWMClass }}
StartupWMClass={{ .Entry.StartupWMClass }}
{{- end -}}
{{ if .Entry.URL }}
URL={{ .Entry.URL }}
{{- end -}}
{{ if true }}
PrefersNonDefaultGPU={{ .Entry.PrefersNonDefaultGPU }}
{{- end -}}
{{ if .Entry.Actions }}
Actions={{ range $v := .Entry.Actions }}{{ $v.Name }};{{ end }}
{{- range $v := .Entry.Actions }}

[Desktop Action {{$v.Name}}]
Name={{$v.Name}}
Icon={{$v.Icon}}
Exec={{$v.Exec}}
{{- end -}}
{{- end -}}
{{- "" }}
`
	buf := &bytes.Buffer{}

	tmpl, err := template.New("DesktopEntry").Parse(desktopEntry)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(buf, f)
	if err != nil {
		log.Fatal(err)
	}

	return buf.Bytes()
}

func (f File) Save() {
	if _, err := os.Stat(f.FullPath); err == nil {
		log.Fatalf("file already exists: %v", f.FullPath)
	}

	if _, err := os.Stat(f.Path); os.IsNotExist(err) {
		if err != nil {
			log.Fatal(fmt.Errorf("path error: %w", err))
		}
	}

	err := ioutil.WriteFile(f.FullPath, f.Content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func NewFile(path, name string, entry Entry) *File {
	f := &File{
		Path:  path,
		Name:  name,
		Entry: entry,
	}

	// Set File.Extension.
	switch entry.TypeKey {
	case "Application":
		f.Extension = "desktop"
	case "Directory":
		f.Extension = "directory"
	default:
		log.Fatal("invalid type")
	}

	// Set File.FullPath: use path or working directory.
	if f.Path == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		f.Path = wd
	}

	f.FullPath = f.Path + "/" + f.Name + "." + f.Extension

	// Set File.Content.
	f.Content = f.parseContent()

	return f
}
