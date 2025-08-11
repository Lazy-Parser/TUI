// go run ./cmd/scaffold/scaffold.go -name=viewer

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "internal", "tui", "components", "pages")
	name := flag.String("name", "", "page name (e.g. users)")
	root := flag.String("root", path, "pages root")
	pkg := flag.String("pkg", "pages", "import path alias for pages package (if needed)")
	flag.Parse()

	if *name == "" {
		fmt.Fprintln(os.Stderr, "Usage: scaffold -name=<page>")
		os.Exit(2)
	}

	pageName := strings.ToLower(*name)
	dir := filepath.Join(*root, pageName)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		panic(err)
	}

	// files to create
	files := map[string]string{
		pageName + ".go": pageTpl,
		"header.go":      headerTpl,
		"body.go":        mainTpl,
		"footer.go":      footerTpl,
	}

	data := struct {
		PageName string // e.g. "users"
		Title    string // e.g. "Users"
		PkgPages string // e.g. "tui/internal/tui/components/pages"
	}{
		PageName: pageName,
		Title:    strings.Title(pageName),
		PkgPages: *pkg,
	}

	for name, tpl := range files {
		path := filepath.Join(dir, name)
		if _, err := os.Stat(path); err == nil {
			fmt.Println("skip (exists):", path)
			continue
		}
		if err := renderToFile(path, tpl, data); err != nil {
			panic(err)
		}
		fmt.Println("created:", path)
	}
}

func renderToFile(path, tpl string, data any) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	t := template.Must(template.New("").Parse(tpl))
	return t.Execute(f, data)
}

const pageTpl = `package page_{{.PageName}}

import (
	"tui/internal/tui/components/pages"
)

func NewPage() *pages.Page {
	return &pages.Page{
		Header: NewHeader(),
		Main:   NewMain(),
		Footer: NewFooter(),
	}
}
`

const headerTpl = `package page_{{.PageName}}

import (
	tea "github.com/charmbracelet/bubbletea"
)

type header struct{}

func NewHeader() tea.Model { return &header{} }

func (h *header) Init() tea.Cmd                            { return nil }
func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd)  { return h, nil }
func (h *header) View() string                             { return "Header" }
`

const mainTpl = `package page_{{.PageName}}

import (
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct{}

func NewMain() tea.Model                                   { return &mainView{} }
func (m *mainView) Init() tea.Cmd                          { return nil }
func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd){ return m, nil }
func (m *mainView) View() string                           { return "Content goes here" }
`

const footerTpl = `package page_{{.PageName}}

import tea "github.com/charmbracelet/bubbletea"

type footer struct{}

func NewFooter() tea.Model                                  { return &footer{} }
func (f *footer) Init() tea.Cmd                             { return nil }
func (f *footer) Update(msg tea.Msg) (tea.Model, tea.Cmd)   { return f, nil }
func (f *footer) View() string                              { return "Footer" }
`
