package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// ConfigurarEstilosAltoContraste ajusta la paleta global de colores de tview a alto contraste.
func ConfigurarEstilosAltoContraste() {
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	tview.Styles.ContrastBackgroundColor = tcell.ColorBlue
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorDarkBlue
	tview.Styles.BorderColor = tcell.ColorTeal
	tview.Styles.TitleColor = tcell.ColorYellow
	tview.Styles.GraphicsColor = tcell.ColorYellow
	tview.Styles.PrimaryTextColor = tcell.ColorWhite
	tview.Styles.SecondaryTextColor = tcell.ColorLightCyan
	tview.Styles.TertiaryTextColor = tcell.ColorGreen
	tview.Styles.InverseTextColor = tcell.ColorBlack
	tview.Styles.ContrastSecondaryTextColor = tcell.ColorYellow
}

// AppRouter gestiona el flujo de navegación entre pantallas utilizando tview.Pages.
type AppRouter struct {
	app        *tview.Application
	pages      *tview.Pages
	primitives map[string]tview.Primitive
}

// NewAppRouter inicializa una nueva instancia de AppRouter con paleta de alto contraste.
func NewAppRouter() *AppRouter {
	ConfigurarEstilosAltoContraste()

	app := tview.NewApplication()
	pages := tview.NewPages()
	return &AppRouter{
		app:        app,
		pages:      pages,
		primitives: make(map[string]tview.Primitive),
	}
}

// AgregarPantalla registra un nuevo primitivo tview bajo un nombre de pantalla.
func (r *AppRouter) AgregarPantalla(nombre string, primitive tview.Primitive) {
	r.primitives[nombre] = primitive
	r.pages.AddPage(nombre, primitive, true, false)
}

// CambiarPantalla conmuta la pantalla visible activa y transfiere explícitamente el foco.
func (r *AppRouter) CambiarPantalla(nombre string) {
	r.pages.SwitchToPage(nombre)
	r.app.SetFocus(r.pages)
}

// MostrarModalError muestra una ventana emergente interactiva de aviso u error con alto contraste.
func (r *AppRouter) MostrarModalError(msg string) {
	previousFocus := r.app.GetFocus()

	modal := tview.NewModal().
		SetText("\n[yellow:][bold]" + msg).
		SetBackgroundColor(tcell.ColorDarkBlue).
		SetTextColor(tcell.ColorWhite).
		AddButtons([]string{"Aceptar"}).
		SetButtonBackgroundColor(tcell.ColorBlue).
		SetButtonTextColor(tcell.ColorYellow)

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		r.pages.RemovePage("error_modal")
		if previousFocus != nil {
			r.app.SetFocus(previousFocus)
		}
	})

	r.pages.AddPage("error_modal", modal, false, true)
	r.app.SetFocus(modal)
}

// GetApplication retorna el puntero al tview.Application principal.
func (r *AppRouter) GetApplication() *tview.Application {
	return r.app
}

// Salir detiene la aplicación tview.
func (r *AppRouter) Salir() {
	r.app.Stop()
}

// Run inicia el bucle principal de la aplicación tview.
func (r *AppRouter) Run(paginaInicial string) error {
	r.CambiarPantalla(paginaInicial)
	return r.app.SetRoot(r.pages, true).EnableMouse(true).Run()
}
