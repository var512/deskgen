package desktop

import (
	"errors"

	"github.com/var512/deskgen/internal/flags"
)

type Entry struct {
	TypeKey              string
	Name                 string
	Version              string
	GenericName          string
	NoDisplay            bool
	Comment              string
	Icon                 string
	Hidden               bool
	OnlyShowIn           string
	NotShowIn            string
	DbusActivatable      bool
	TryExec              string
	Exec                 string
	Path                 string
	Terminal             bool
	MimeType             string
	Categories           string
	Implements           string
	Keywords             string
	StartupNotify        bool
	StartupWMClass       string
	URL                  string
	PrefersNonDefaultGPU bool
	Actions              []Action
}

// TODO validation.
// Functional options.
type Option func(*Entry) error

func Version(version string) Option {
	return func(e *Entry) error {
		e.Version = version
		return nil
	}
}

func GenericName(genericName string) Option {
	return func(e *Entry) error {
		e.GenericName = genericName
		return nil
	}
}

func NoDisplay(noDisplay bool) Option {
	return func(e *Entry) error {
		e.NoDisplay = noDisplay
		return nil
	}
}

func Comment(comment string) Option {
	return func(e *Entry) error {
		e.Comment = comment
		return nil
	}
}

func Icon(icon string) Option {
	return func(e *Entry) error {
		e.Icon = icon
		return nil
	}
}

func Hidden(hidden bool) Option {
	return func(e *Entry) error {
		e.Hidden = hidden
		return nil
	}
}

func OnlyShowIn(onlyShowIn string) Option {
	return func(e *Entry) error {
		e.OnlyShowIn = onlyShowIn
		return nil
	}
}

func NotShowIn(notShowIn string) Option {
	return func(e *Entry) error {
		e.NotShowIn = notShowIn
		return nil
	}
}

func DbusActivatable(dbusActivatable bool) Option {
	return func(e *Entry) error {
		e.DbusActivatable = dbusActivatable
		return nil
	}
}

func TryExec(tryExec string) Option {
	return func(e *Entry) error {
		e.TryExec = tryExec
		return nil
	}
}

func Exec(exec string) Option {
	return func(e *Entry) error {
		e.Exec = exec
		return nil
	}
}

func Path(path string) Option {
	return func(e *Entry) error {
		e.Path = path
		return nil
	}
}

func Terminal(terminal bool) Option {
	return func(e *Entry) error {
		e.Terminal = terminal
		return nil
	}
}

func MimeType(mimeType string) Option {
	return func(e *Entry) error {
		e.MimeType = mimeType
		return nil
	}
}

func Categories(categories string) Option {
	return func(e *Entry) error {
		e.Categories = categories
		return nil
	}
}

func Implements(implements string) Option {
	return func(e *Entry) error {
		e.Implements = implements
		return nil
	}
}

func Keywords(keywords string) Option {
	return func(e *Entry) error {
		e.Keywords = keywords
		return nil
	}
}

func StartupNotify(startupNotify bool) Option {
	return func(e *Entry) error {
		e.StartupNotify = startupNotify
		return nil
	}
}

func StartupWMClass(startupWMClass string) Option {
	return func(e *Entry) error {
		e.StartupWMClass = startupWMClass
		return nil
	}
}

func URL(url string) Option {
	return func(e *Entry) error {
		e.URL = url
		return nil
	}
}

func PrefersNonDefaultGPU(prefersNonDefaultGPU bool) Option {
	return func(e *Entry) error {
		e.PrefersNonDefaultGPU = prefersNonDefaultGPU
		return nil
	}
}

// Do it properly with a 3rd-party flags array/groups dep if this gets too complex.
func Actions(name flags.ActionName, icon flags.ActionIcon, exec flags.ActionExec) Option {
	return func(e *Entry) error {
		if (len(name)+len(icon)+len(exec))%3 != 0 {
			return errors.New("all action fields are required: name, icon, exec")
		}

		for i := range name {
			// Action spec requires a name.
			if name[i] == "" {
				return errors.New("action name is required")
			}

			e.Actions = append(e.Actions, Action{
				Name: name[i],
				Icon: icon[i],
				Exec: exec[i],
			})
		}

		return nil
	}
}

func NewEntry(typeKey, name string, opts ...Option) (*Entry, error) {
	entry := &Entry{
		TypeKey: typeKey,
		Name:    name,
	}

	for _, opt := range opts {
		err := opt(entry)
		if err != nil {
			return nil, err
		}
	}

	if entry.Version == "" {
		entry.Version = "1.0"
	}

	return entry, nil
}
