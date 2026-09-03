package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// LoginView implementa la interfaz ComponentView para la pantalla de inicio de sesión.
type LoginView struct {
	form          *tview.Form
	usernameField *tview.InputField
	passwordField *tview.InputField
	errorLabel    *tview.TextView
	onLogin       func(u, p string)
	onExit        func()
	layout        *tview.Flex
}

// NewLoginView crea y construye la interfaz visual del Login.
func NewLoginView() *LoginView {
	lv := &LoginView{}

	lv.usernameField = tview.NewInputField().
		SetLabel("Usuario: ").
		SetFieldWidth(25)

	lv.passwordField = tview.NewInputField().
		SetLabel("Contraseña: ").
		SetMaskCharacter('*').
		SetFieldWidth(25)

	lv.passwordField.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter && lv.onLogin != nil {
			u, p := lv.ObtenerCredenciales()
			lv.onLogin(u, p)
		}
	})

	lv.errorLabel = tview.NewTextView().
		SetTextColor(tcell.ColorRed).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	lv.form = tview.NewForm().
		AddFormItem(lv.usernameField).
		AddFormItem(lv.passwordField).
		AddButton("Iniciar Sesión", func() {
			if lv.onLogin != nil {
				u, p := lv.ObtenerCredenciales()
				lv.onLogin(u, p)
			}
		}).
		AddButton("Limpiar", func() {
			lv.usernameField.SetText("")
			lv.passwordField.SetText("")
			lv.errorLabel.SetText("")
		}).
		AddButton("Salir", func() {
			if lv.onExit != nil {
				lv.onExit()
			}
		})

	lv.form.SetBorder(true).
		SetTitle(" 🔒 Sistema Bibliotecario - Acceso ").
		SetTitleAlign(tview.AlignCenter)

	// Layout centrado en pantalla
	lv.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(lv.form, 11, 1, true).
				AddItem(lv.errorLabel, 2, 1, false), 50, 1, true).
			AddItem(nil, 0, 1, false), 13, 1, true).
		AddItem(nil, 0, 1, false)

	return lv
}

// ObtenerCredenciales retorna el usuario y contraseña ingresados en el formulario.
func (lv *LoginView) ObtenerCredenciales() (string, string) {
	return lv.usernameField.GetText(), lv.passwordField.GetText()
}

// MostrarError actualiza el mensaje de error en la vista.
func (lv *LoginView) MostrarError(msg string) {
	lv.errorLabel.SetText(msg)
}

// SetOnLogin asigna el callback para procesar la autenticación.
func (lv *LoginView) SetOnLogin(fn func(u, p string)) {
	lv.onLogin = fn
}

// SetOnExit asigna el callback para salir de la aplicación.
func (lv *LoginView) SetOnExit(fn func()) {
	lv.onExit = fn
}

// GetPrimitive retorna el elemento raíz tview de la vista.
func (lv *LoginView) GetPrimitive() tview.Primitive {
	return lv.layout
}
