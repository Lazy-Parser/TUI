package common

import (
	"reflect"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func CmdHandler(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}

func KeyMapToSlice(t any) (bindings []key.Binding) {
	typ := reflect.TypeOf(t)
	if typ.Kind() != reflect.Struct {
		return nil
	}
	for i := 0; i < typ.NumField(); i++ {
		v := reflect.ValueOf(t).Field(i)
		bindings = append(bindings, v.Interface().(key.Binding))
	}
	return
}
