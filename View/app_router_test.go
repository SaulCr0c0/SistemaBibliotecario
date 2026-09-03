package view

import (
	"testing"
	"github.com/rivo/tview"
)

func TestAppRouter_CambiarPantallaFoco(t *testing.T) {
	router := NewAppRouter()

	input1 := tview.NewInputField()
	input2 := tview.NewInputField()

	router.AgregarPantalla("p1", input1)
	router.AgregarPantalla("p2", input2)

	router.GetApplication().SetRoot(router.pages, true)

	router.CambiarPantalla("p1")
	if router.GetApplication().GetFocus() == nil {
		t.Errorf("Esperaba foco en p1 al cambiar a p1")
	}

	router.CambiarPantalla("p2")
	if router.GetApplication().GetFocus() == nil {
		t.Errorf("Esperaba foco en p2 al cambiar a p2")
	}
}
