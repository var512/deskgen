package main

import (
	"flag"
	"fmt"

	"github.com/var512/deskgen/internal/desktop"
	"github.com/var512/deskgen/internal/flags"
)

var (
	// Command.
	filePath = flag.String("filePath", "", "FilePath")
	fileName = flag.String("fileName", "", "FileName")

	// Desktop Entry Specification.
	typeKey              = flag.String("type", "", "Type")
	name                 = flag.String("name", "", "Name")
	version              = flag.String("version", "", "Version")
	genericName          = flag.String("genericName", "", "GenericName")
	noDisplay            = flag.Bool("noDisplay", false, "NoDisplay")
	comment              = flag.String("comment", "", "Comment")
	icon                 = flag.String("icon", "", "Icon")
	hidden               = flag.Bool("hidden", false, "Hidden")
	onlyShowIn           = flag.String("onlyShowIn", "", "OnlyShowIn")
	notShowIn            = flag.String("notShowIn", "", "NotShowIn")
	dbusActivatable      = flag.Bool("dbusActivatable", false, "DBusActivatable")
	tryExec              = flag.String("tryExec", "", "TryExec")
	exec                 = flag.String("exec", "", "Exec")
	path                 = flag.String("path", "", "Path")
	terminal             = flag.Bool("terminal", false, "Terminal")
	mimeType             = flag.String("mimeType", "", "MimeType")
	categories           = flag.String("categories", "", "Categories")
	implements           = flag.String("implements", "", "Implements")
	keywords             = flag.String("keywords", "", "Keywords")
	startupNotify        = flag.Bool("startupNotify", false, "StartupNotify")
	startupWMClass       = flag.String("startupWMClass", "", "StartupWMClass")
	URL                  = flag.String("url", "", "URL")
	prefersNonDefaultGPU = flag.Bool("prefersNonDefaultGPU", false, "PrefersNonDefaultGPU")

	actionName flags.ActionName
	actionIcon flags.ActionIcon
	actionExec flags.ActionExec
)

func createEntry() *desktop.Entry {
	e := desktop.NewEntry(
		*typeKey,
		*name,
		desktop.Version(*version),
		desktop.GenericName(*genericName),
		desktop.NoDisplay(*noDisplay),
		desktop.Comment(*comment),
		desktop.Icon(*icon),
		desktop.Hidden(*hidden),
		desktop.OnlyShowIn(*onlyShowIn),
		desktop.NotShowIn(*notShowIn),
		desktop.DbusActivatable(*dbusActivatable),
		desktop.TryExec(*tryExec),
		desktop.Exec(*exec),
		desktop.Path(*path),
		desktop.Terminal(*terminal),
		desktop.MimeType(*mimeType),
		desktop.Categories(*categories),
		desktop.Implements(*implements),
		desktop.Keywords(*keywords),
		desktop.StartupNotify(*startupNotify),
		desktop.StartupWMClass(*startupWMClass),
		desktop.URL(*URL),
		desktop.PrefersNonDefaultGPU(*prefersNonDefaultGPU),
		desktop.Actions(actionName, actionIcon, actionExec),
	)

	return e
}

func createFile(path, name string, entry desktop.Entry) *desktop.File {
	f := desktop.NewFile(path, name, entry)

	return f
}

func main() {
	flag.Var(&actionName, "actionName", "Action.Name")
	flag.Var(&actionIcon, "actionIcon", "Action.Icon")
	flag.Var(&actionExec, "actionExec", "Action.Exec")

	flag.Parse()

	e := createEntry()
	f := createFile(*filePath, *fileName, *e)

	if *fileName != "" {
		f.Save()
	}

	fmt.Println(string(f.Content))
}
