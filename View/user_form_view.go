package view

import (
	"strconv"

	"biblioteca/Model"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// UserFormView representa el formulario para crear y editar usuarios con estilos de alto contraste.
type UserFormView struct {
	form             *tview.Form
	usernameField    *tview.InputField
	passwordField    *tview.InputField
	rolDropDown      *tview.DropDown
	nivelAccesoField *tview.InputField
	turnoField       *tview.InputField
	errorLabel       *tview.TextView
	layout           *tview.Flex
	usuarioID        int // 0 si es un nuevo usuario
	onGuardar        func(id int, username, password, rol string, nivelAcceso int, turno string)
	onCancelar       func()
}

// NewUserFormView instancia la vista de formulario de usuario con paleta de alto contraste.
func NewUserFormView() *UserFormView {
	ufv := &UserFormView{}

	ufv.usernameField = tview.NewInputField().
		SetLabel("Nombre de Usuario: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(25)

	ufv.passwordField = tview.NewInputField().
		SetLabel("Contraseña: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetMaskCharacter('*').
		SetFieldWidth(25)

	ufv.rolDropDown = tview.NewDropDown().
		SetLabel("Rol: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite),
			tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorYellow).Bold(true),
		).
		SetOptions([]string{"Administrador", "Bibliotecario"}, func(option string, optionIndex int) {
			ufv.actualizarVisibilidadCampos(option)
		}).
		SetCurrentOption(0)

	ufv.nivelAccesoField = tview.NewInputField().
		SetLabel("Nivel de Acceso (1-10): ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(10).
		SetText("1")

	ufv.turnoField = tview.NewInputField().
		SetLabel("Turno (Mañana/Tarde/Noche): ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(20).
		SetText("Mañana")

	ufv.errorLabel = tview.NewTextView().
		SetTextColor(tcell.ColorRed).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	ufv.form = tview.NewForm().
		AddFormItem(ufv.usernameField).
		AddFormItem(ufv.passwordField).
		AddFormItem(ufv.rolDropDown).
		AddFormItem(ufv.nivelAccesoField).
		AddFormItem(ufv.turnoField).
		AddButton(" Guardar ", func() {
			if ufv.onGuardar != nil {
				id, u, p, r, nivel, turno, errStr := ufv.ObtenerDatos()
				if errStr != "" {
					ufv.MostrarError(errStr)
					return
				}
				ufv.onGuardar(id, u, p, r, nivel, turno)
			}
		}).
		AddButton(" Cancelar ", func() {
			if ufv.onCancelar != nil {
				ufv.onCancelar()
			}
		})

	ufv.form.SetButtonBackgroundColor(tcell.ColorBlue)
	ufv.form.SetButtonTextColor(tcell.ColorYellow)
	ufv.form.SetFieldBackgroundColor(tcell.ColorDarkBlue)
	ufv.form.SetFieldTextColor(tcell.ColorWhite)
	ufv.form.SetLabelColor(tcell.ColorYellow)

	ufv.form.SetBorder(true)
	ufv.form.SetBorderColor(tcell.ColorTeal)
	ufv.form.SetTitle(" 📝 Formulario de Usuario ")
	ufv.form.SetTitleColor(tcell.ColorYellow)
	ufv.form.SetTitleAlign(tview.AlignCenter)

	ufv.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(ufv.form, 16, 1, true).
				AddItem(ufv.errorLabel, 2, 1, false), 60, 1, true).
			AddItem(nil, 0, 1, false), 18, 1, true).
		AddItem(nil, 0, 1, false)

	return ufv
}

func (ufv *UserFormView) actualizarVisibilidadCampos(rol string) {
}

// CargarDatosModoCrear prepara el formulario para agregar un usuario.
func (ufv *UserFormView) CargarDatosModoCrear() {
	ufv.usuarioID = 0
	ufv.form.SetTitle(" ➕ Crear Nuevo Usuario ")
	ufv.usernameField.SetText("")
	ufv.passwordField.SetText("")
	ufv.rolDropDown.SetCurrentOption(0)
	ufv.nivelAccesoField.SetText("1")
	ufv.turnoField.SetText("Mañana")
	ufv.errorLabel.SetText("")
	ufv.form.SetFocus(0)
}

// CargarDatosModoEditar puebla los campos del formulario para modificar a un usuario existente.
func (ufv *UserFormView) CargarDatosModoEditar(u *model.Usuario) {
	ufv.usuarioID = u.ID
	ufv.form.SetTitle(" ✏️ Editar Usuario ")
	ufv.usernameField.SetText(u.Username)
	ufv.passwordField.SetText(u.GetPassword())

	if u.Rol == "Administrador" {
		ufv.rolDropDown.SetCurrentOption(0)
	} else {
		ufv.rolDropDown.SetCurrentOption(1)
	}

	ufv.nivelAccesoField.SetText(strconv.Itoa(u.NivelAcceso))
	ufv.turnoField.SetText(u.Turno)
	ufv.errorLabel.SetText("")
	ufv.form.SetFocus(0)
}

// ObtenerDatos extrae y valida los datos ingresados en el formulario.
func (ufv *UserFormView) ObtenerDatos() (int, string, string, string, int, string, string) {
	username := ufv.usernameField.GetText()
	password := ufv.passwordField.GetText()
	_, rol := ufv.rolDropDown.GetCurrentOption()

	if username == "" {
		return 0, "", "", "", 0, "", "El nombre de usuario es obligatorio"
	}

	nivelAcceso := 1
	if rol == "Administrador" {
		val, err := strconv.Atoi(ufv.nivelAccesoField.GetText())
		if err != nil || val < 1 {
			return 0, "", "", "", 0, "", "Nivel de acceso inválido (debe ser un entero >= 1)"
		}
		nivelAcceso = val
	}

	turno := ufv.turnoField.GetText()
	if rol == "Bibliotecario" && turno == "" {
		turno = "Mañana"
	}

	return ufv.usuarioID, username, password, rol, nivelAcceso, turno, ""
}

// MostrarError publica un mensaje de error en la vista del formulario con resaltado.
func (ufv *UserFormView) MostrarError(msg string) {
	if msg != "" {
		ufv.errorLabel.SetText("[red:][bold]❌ " + msg)
	} else {
		ufv.errorLabel.SetText("")
	}
}

// SetOnGuardar asigna el callback para guardar los cambios del usuario.
func (ufv *UserFormView) SetOnGuardar(fn func(id int, username, password, rol string, nivelAcceso int, turno string)) {
	ufv.onGuardar = fn
}

// SetOnCancelar asigna el callback para cancelar y volver.
func (ufv *UserFormView) SetOnCancelar(fn func()) {
	ufv.onCancelar = fn
}

// GetPrimitive retorna el elemento raíz de la vista del formulario.
func (ufv *UserFormView) GetPrimitive() tview.Primitive {
	return ufv.layout
}
