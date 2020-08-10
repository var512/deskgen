package desktop

import (
	"strings"
	"testing"
)

const (
	minValidTotalLines = 4 // Header, Type, Name, NL@EOF.
	desktopEntryHeader = "[Desktop Entry]"
)

func TestParsedContent(t *testing.T) {
	entry, err := NewEntry("Application", "Test")
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	file, err := NewFile("", "", *entry)
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	content, err := file.parseContent()
	if err != nil {
		t.Fatalf("unexpected error: %#v", err)
	}

	lines := strings.Split(string(content), "\n")

	if len(lines) < minValidTotalLines {
		t.Errorf("want %v or more lines, got: %#v", minValidTotalLines, len(lines))
	}

	if lines[0] != desktopEntryHeader {
		t.Errorf("want header %v, got %#v", desktopEntryHeader, lines[0])
	}

	for i := 1; i < len(lines); i++ {
		if lines[i] == "" && lines[i-1] == "" {
			t.Error("want single empty new line, got many")
		}
	}

	if lines[len(lines)-1] != "" {
		t.Errorf("want new line at EOF, got: %#v", lines[len(lines)-1])
	}
}

func TestFileExtension(t *testing.T) {
	for typeKey, extension := range TypeExtension() {
		entry, err := NewEntry(typeKey, "Test")
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		file, err := NewFile("", "", *entry)
		if err != nil {
			t.Fatalf("unexpected error: %#v", err)
		}

		if file.Extension != extension {
			t.Errorf("want extension %v, got: %#v", extension, file.Extension)
		}
	}
}
