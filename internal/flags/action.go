package flags

import (
	"strings"
)

type (
	// Flag Value interface.
	ActionName []string
	ActionIcon []string
	ActionExec []string
)

func (i *ActionName) String() string {
	return ""
}

func (i *ActionIcon) String() string {
	return ""
}

func (i *ActionExec) String() string {
	return ""
}

func (i *ActionName) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func (i *ActionIcon) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}

func (i *ActionExec) Set(value string) error {
	*i = append(*i, strings.TrimSpace(value))
	return nil
}
