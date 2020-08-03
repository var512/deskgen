package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/var512/deskgen/internal/desktop"
	"github.com/var512/deskgen/internal/flags"
)

var (
	// Command.
	filePath = flag.String("filePath", "", "filesystem path to write the file to")
	fileName = flag.String("fileName", "", "file name without extension")

	// Desktop Entry.
	typeKey              = flag.String("type", "", "entry key: Type")
	name                 = flag.String("name", "", "entry key: Name")
	version              = flag.String("version", "", "entry key: Version")
	genericName          = flag.String("genericName", "", "entry key: GenericName")
	noDisplay            = flag.Bool("noDisplay", false, "entry key: NoDisplay")
	comment              = flag.String("comment", "", "entry key: Comment")
	icon                 = flag.String("icon", "", "entry key: Icon")
	hidden               = flag.Bool("hidden", false, "entry key: Hidden")
	onlyShowIn           = flag.String("onlyShowIn", "", "entry key: OnlyShowIn")
	notShowIn            = flag.String("notShowIn", "", "entry key: NotShowIn")
	dbusActivatable      = flag.Bool("dbusActivatable", false, "entry key: DBusActivatable")
	tryExec              = flag.String("tryExec", "", "entry key: TryExec")
	exec                 = flag.String("exec", "", "entry key: Exec")
	path                 = flag.String("path", "", "entry key: Path")
	terminal             = flag.Bool("terminal", false, "entry key: Terminal")
	mimeType             = flag.String("mimeType", "", "entry key: MimeType")
	categories           = flag.String("categories", "", "entry key: Categories")
	implements           = flag.String("implements", "", "entry key: Implements")
	keywords             = flag.String("keywords", "", "entry key: Keywords")
	startupNotify        = flag.Bool("startupNotify", false, "entry key: StartupNotify")
	startupWMClass       = flag.String("startupWMClass", "", "entry key: StartupWMClass")
	URL                  = flag.String("url", "", "entry key: URL")
	prefersNonDefaultGPU = flag.Bool("prefersNonDefaultGPU", false, "entry key: PrefersNonDefaultGPU")

	// Desktop Action.
	actionName flags.ActionName
	actionIcon flags.ActionIcon
	actionExec flags.ActionExec
)

func createEntry() *desktop.Entry {
	entry, err := desktop.NewEntry(
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
	if err != nil {
		log.Fatal(err)
	}

	return entry
}

func createFile(path, name string, entry desktop.Entry) *desktop.File {
	file, err := desktop.NewFile(path, name, entry)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func main() {
	flag.Var(&actionName, "actionName", "action key: Name")
	flag.Var(&actionIcon, "actionIcon", "action key: Icon")
	flag.Var(&actionExec, "actionExec", "action key: Exec")

	flag.Parse()

	e := createEntry()
	f := createFile(*filePath, *fileName, *e)

	if *fileName != "" {
		err := f.Save()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(string(f.Content))
}
