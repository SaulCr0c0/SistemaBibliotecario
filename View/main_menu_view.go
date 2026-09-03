package view

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// MainMenuView representa el menú principal adaptado según el rol del usuario con alto contraste.
type MainMenuView struct {
	layout   *tview.Flex
	header   *tview.TextView
	menu     *tview.List
	onSelect func(opc string)
}

// NewMainMenuView instancia la vista del menú principal con colores de alto contraste.
func NewMainMenuView() *MainMenuView {
	mv := &MainMenuView{}

	mv.header = tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter)

	mv.menu = tview.NewList().
		SetMainTextColor(tcell.ColorWhite).
		SetSecondaryTextColor(tcell.ColorLightCyan).
		SetShortcutColor(tcell.ColorYellow).
		SetSelectedBackgroundColor(tcell.ColorBlue).
		SetSelectedTextColor(tcell.ColorYellow)

	mv.menu.SetBorder(true).
		SetBorderColor(tcell.ColorTeal).
		SetTitle(" 📚 Menú Principal - Sistema Bibliotecario ").
		SetTitleColor(tcell.ColorYellow).
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

	// Configuración inicial por defecto
	mv.ConfigurarMenuPorRol("")

	return mv
}

// ConfigurarMenuPorRol adapta dinámicamente las opciones visibles según el rol.
func (mv *MainMenuView) ConfigurarMenuPorRol(rol string) {
	mv.menu.Clear()

	// La opción de 'Gestionar Usuarios' solo está disponible para Administradores
	if rol == "Administrador" {
		mv.menu.AddItem("Gestionar Usuarios", "Alta, modificación y control de acceso de usuarios", 'u', nil)
	}

	// Opciones disponibles para todos los roles (Bibliotecario / Administrador)
	mv.menu.AddItem("Gestionar Materiales", "Libros, Tesis y Revistas", 'm', nil).
		AddItem("Gestionar Lectores", "Profesores y Estudiantes", 'l', nil).
		AddItem("Gestionar Préstamos", "Préstamos y Sanciones", 'p', nil).
		AddItem("Generar Reportes", "Métricas e Infracciones", 'r', nil).
		AddItem("Cerrar Sesión", "Volver a la pantalla de acceso", 'c', nil)
}

// ActualizarInfoUsuario actualiza el encabezado y adapta el menú al rol autenticado.
func (mv *MainMenuView) ActualizarInfoUsuario(username, rol string) {
	mv.header.SetText(fmt.Sprintf("\n[yellow:][bold]Usuario activo: [white:][bold]%s [green:][bold]| Rol: %s", username, rol))
	mv.ConfigurarMenuPorRol(rol)
}

// SetOnSelect establece el manejador de eventos de selección del menú.
func (mv *MainMenuView) SetOnSelect(fn func(opc string)) {
	mv.onSelect = fn
}

// GetPrimitive retorna el elemento raíz de la vista del menú.
func (mv *MainMenuView) GetPrimitive() tview.Primitive {
	return mv.layout
}
