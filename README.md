# deskgen
 
Generate .desktop files following the [Desktop Entry Specification](https://specifications.freedesktop.org/desktop-entry-spec/desktop-entry-spec-latest.html).

This is my first Go project. Only what I currently need from the specification was implemented. Use at your own risk.

**Related projects**

- https://github.com/rkoesters/xdg
- https://github.com/xyproto/gendesk

### Installation

```sh
go get -u github.com/var512/deskgen/cmd/deskgen
```

### Usage

**Options**

```sh
${GOPATH}/bin/deskgen --help
```

**Stdout**

```sh
${GOPATH}/bin/deskgen \
    --type="Application" \
    --name="My script" \
    --genericName="Custom script" \
    --comment="A custom script" \
    --icon="applications-utilities" \
    --exec="my-script.sh" \
    --mimeType="inode/directory;text/html;" \
    --categories="Utility;" \
    --keywords="Internet;Development;" \
    --startupNotify="true"
```

**Save to disk**

If `filePath` is not provided the file will be saved in the current working directory.

```sh
${GOPATH}/bin/deskgen \
    --fileName="my-script" \
    --filePath "/home/user/.local/share/applications" \
    --type="Application" \
    --name="My script" \
    --genericName="Custom script" \
    --comment="A custom script" \
    --icon="applications-utilities" \
    --exec="my-script.sh" \
    --mimeType="inode/directory;text/html;" \
    --categories="Utility;" \
    --keywords="Internet;Development;" \
    --startupNotify="true"
```

**Actions**

Very basic functionality. Each action requires all fields to be set (name, icon, exec).

```sh
${GOPATH}/bin/deskgen \
    --type="Application" \
    --name="My script" \
    --actionName="Example 1" \
    --actionIcon="" \
    --actionExec="" \
    --actionName="Example 2" \
    --actionIcon="arrow-right" \
    --actionExec="" \
    --actionName="New Private Window" \
    --actionIcon="icon-private" \
    --actionExec='/bin/example --new-private-window --with="double-quotes" %u'
```
