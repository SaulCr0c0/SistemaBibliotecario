package view

import (
	"fmt"
	"github.com/rivo/tview"
)

// MainMenuView representa el menú principal tras iniciar sesión.
type MainMenuView struct {
	layout   *tview.Flex
	header   *tview.TextView
	menu     *tview.List
	onSelect func(opc string)
}

// NewMainMenuView instancia la vista del menú principal.
func NewMainMenuView() *MainMenuView {
	mv := &MainMenuView{}

	mv.header = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	mv.menu = tview.NewList().
		AddItem("Gestionar Materiales", "Libros, Tesis y Revistas", 'm', nil).
		AddItem("Gestionar Lectores", "Profesores y Estudiantes", 'l', nil).
		AddItem("Gestionar Préstamos", "Préstamos y Sanciones", 'p', nil).
		AddItem("Generar Reportes", "Métricas e Infracciones", 'r', nil).
		AddItem("Cerrar Sesión", "Volver a la pantalla de acceso", 'c', nil)

	mv.menu.SetBorder(true).
		SetTitle(" 📚 Menú Principal - Sistema Bibliotecario ").
		SetTitleAlign(tview.AlignCenter)

	mv.menu.SetSelectedFunc(func(index int, mainText, secondaryText string, shortcut rune) {
		if mv.onSelect != nil {
			mv.onSelect(mainText)
		}
	})

	mv.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mv.header, 3, 1, false).
		AddItem(mv.menu, 0, 1, true)

	return mv
}

// ActualizarInfoUsuario actualiza el encabezado con la información del usuario autenticado.
func (mv *MainMenuView) ActualizarInfoUsuario(username, rol string) {
	mv.header.SetText(fmt.Sprintf("\n[yellow]Usuario activo: [white]%s [green]| Rol: %s", username, rol))
}

// SetOnSelect establece el manejador de eventos de selección del menú.
func (mv *MainMenuView) SetOnSelect(fn func(opc string)) {
	mv.onSelect = fn
}

// GetPrimitive retorna el primitiva principal del menú.
func (mv *MainMenuView) GetPrimitive() tview.Primitive {
	return mv.layout
}
