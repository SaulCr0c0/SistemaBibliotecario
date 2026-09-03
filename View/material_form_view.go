package view

import (
	"fmt"
	"strconv"

	model "biblioteca/Model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// MaterialFormView representa el formulario para crear y editar materiales (Libros, Tesis, Revistas).
type MaterialFormView struct {
	form           *tview.Form
	tipoDropDown   *tview.DropDown
	tituloField    *tview.InputField
	anioField      *tview.InputField
	disponField    *tview.DropDown
	extra1Field    *tview.InputField // ISBN (Libro) / Carrera (Tesis) / ISSN (Revista)
	extra2Field    *tview.InputField // Editorial (Libro) / Asesor (Tesis) / Edición (Revista)
	errorLabel     *tview.TextView
	layout         *tview.Flex
	materialID     int
	onGuardar      func(id int, tipo, titulo string, anio int, disponible bool, extra1, extra2 string)
	onCancelar     func()
}

// NewMaterialFormView instancia la vista del formulario de materiales.
func NewMaterialFormView() *MaterialFormView {
	mfv := &MaterialFormView{}

	mfv.tipoDropDown = tview.NewDropDown().
		SetLabel("Tipo de Material: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite),
			tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorYellow).Bold(true),
		).
		SetOptions([]string{"Libro", "Tesis", "Revista"}, func(option string, optionIndex int) {
			mfv.actualizarEtiquetasExtra(option)
		}).
		SetCurrentOption(0)

	mfv.tituloField = tview.NewInputField().
		SetLabel("Título: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(30)

	mfv.anioField = tview.NewInputField().
		SetLabel("Año de Publicación: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(10)

	mfv.disponField = tview.NewDropDown().
		SetLabel("Estado: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite),
			tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorYellow).Bold(true),
		).
		SetOptions([]string{"Disponible", "Prestado"}, nil).
		SetCurrentOption(0)

	mfv.extra1Field = tview.NewInputField().
		SetLabel("ISBN: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(25)

	mfv.extra2Field = tview.NewInputField().
		SetLabel("Editorial: ").
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(25)

	mfv.errorLabel = tview.NewTextView().
		SetTextColor(tcell.ColorRed).
		SetTextAlign(tview.AlignCenter).
		SetDynamicColors(true)

	mfv.form = tview.NewForm().
		AddFormItem(mfv.tipoDropDown).
		AddFormItem(mfv.tituloField).
		AddFormItem(mfv.anioField).
		AddFormItem(mfv.disponField).
		AddFormItem(mfv.extra1Field).
		AddFormItem(mfv.extra2Field).
		AddButton(" Guardar ", func() {
			if mfv.onGuardar != nil {
				id, tipo, titulo, anio, disp, e1, e2, errStr := mfv.ObtenerDatos()
				if errStr != "" {
					mfv.MostrarError(errStr)
					return
				}
				mfv.onGuardar(id, tipo, titulo, anio, disp, e1, e2)
			}
		}).
		AddButton(" Cancelar ", func() {
			if mfv.onCancelar != nil {
				mfv.onCancelar()
			}
		})

	mfv.form.SetButtonBackgroundColor(tcell.ColorBlue)
	mfv.form.SetButtonTextColor(tcell.ColorYellow)
	mfv.form.SetFieldBackgroundColor(tcell.ColorDarkBlue)
	mfv.form.SetFieldTextColor(tcell.ColorWhite)
	mfv.form.SetLabelColor(tcell.ColorYellow)

	mfv.form.SetBorder(true)
	mfv.form.SetBorderColor(tcell.ColorTeal)
	mfv.form.SetTitle(" 📝 Formulario de Material Bibliotecario ")
	mfv.form.SetTitleColor(tcell.ColorYellow)
	mfv.form.SetTitleAlign(tview.AlignCenter)

	mfv.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(mfv.form, 18, 1, true).
				AddItem(mfv.errorLabel, 2, 1, false), 65, 1, true).
			AddItem(nil, 0, 1, false), 20, 1, true).
		AddItem(nil, 0, 1, false)

	return mfv
}

func (mfv *MaterialFormView) actualizarEtiquetasExtra(tipo string) {
	if mfv.extra1Field == nil || mfv.extra2Field == nil {
		return
	}
	switch tipo {
	case "Libro":
		mfv.extra1Field.SetLabel("ISBN: ")
		mfv.extra2Field.SetLabel("Editorial: ")
	case "Tesis":
		mfv.extra1Field.SetLabel("Carrera: ")
		mfv.extra2Field.SetLabel("Asesor: ")
	case "Revista":
		mfv.extra1Field.SetLabel("ISSN: ")
		mfv.extra2Field.SetLabel("Edición #: ")
	}
}

// CargarDatosModoCrear limpia los campos para registrar un nuevo material.
func (mfv *MaterialFormView) CargarDatosModoCrear() {
	mfv.materialID = 0
	mfv.form.SetTitle(" ➕ Crear Nuevo Material ")
	mfv.tipoDropDown.SetCurrentOption(0)
	mfv.tituloField.SetText("")
	mfv.anioField.SetText("2024")
	mfv.disponField.SetCurrentOption(0)
	mfv.extra1Field.SetText("")
	mfv.extra2Field.SetText("")
	mfv.errorLabel.SetText("")
	mfv.actualizarEtiquetasExtra("Libro")
	mfv.form.SetFocus(0)
}

// CargarDatosModoEditar carga los datos de un material existente.
func (mfv *MaterialFormView) CargarDatosModoEditar(m model.Prestable) {
	mfv.materialID = m.GetID()
	mfv.form.SetTitle(" ✏️ Editar Material ")
	mfv.tituloField.SetText(m.GetTitulo())
	mfv.anioField.SetText(strconv.Itoa(m.GetAnio()))

	if m.EstaDisponible() {
		mfv.disponField.SetCurrentOption(0)
	} else {
		mfv.disponField.SetCurrentOption(1)
	}

	switch obj := m.(type) {
	case *model.Libro:
		mfv.tipoDropDown.SetCurrentOption(0)
		mfv.actualizarEtiquetasExtra("Libro")
		mfv.extra1Field.SetText(obj.ISBN)
		mfv.extra2Field.SetText(obj.Editorial)
	case *model.Tesis:
		mfv.tipoDropDown.SetCurrentOption(1)
		mfv.actualizarEtiquetasExtra("Tesis")
		mfv.extra1Field.SetText(obj.Carrera)
		mfv.extra2Field.SetText(obj.Asesor)
	case *model.Revista:
		mfv.tipoDropDown.SetCurrentOption(2)
		mfv.actualizarEtiquetasExtra("Revista")
		mfv.extra1Field.SetText(obj.ISSN)
		mfv.extra2Field.SetText(strconv.Itoa(obj.Edicion))
	}

	mfv.errorLabel.SetText("")
	mfv.form.SetFocus(0)
}

// ObtenerDatos extrae y valida los datos del formulario.
func (mfv *MaterialFormView) ObtenerDatos() (int, string, string, int, bool, string, string, string) {
	_, tipo := mfv.tipoDropDown.GetCurrentOption()
	titulo := mfv.tituloField.GetText()
	anioStr := mfv.anioField.GetText()
	_, dispStr := mfv.disponField.GetCurrentOption()
	extra1 := mfv.extra1Field.GetText()
	extra2 := mfv.extra2Field.GetText()

	if titulo == "" {
		return 0, "", "", 0, false, "", "", "El título es obligatorio"
	}

	anio, err := strconv.Atoi(anioStr)
	if err != nil || anio <= 0 {
		return 0, "", "", 0, false, "", "", "Año de publicación inválido"
	}

	if extra1 == "" {
		return 0, "", "", 0, false, "", "", fmt.Sprintf("El campo %s es obligatorio", mfv.extra1Field.GetLabel())
	}
	if extra2 == "" {
		return 0, "", "", 0, false, "", "", fmt.Sprintf("El campo %s es obligatorio", mfv.extra2Field.GetLabel())
	}

	if tipo == "Revista" {
		_, err := strconv.Atoi(extra2)
		if err != nil {
			return 0, "", "", 0, false, "", "", "El número de edición debe ser un entero válido"
		}
	}

	disponible := (dispStr == "Disponible")
	return mfv.materialID, tipo, titulo, anio, disponible, extra1, extra2, ""
}

// MostrarError muestra un aviso en el pie del formulario.
func (mfv *MaterialFormView) MostrarError(msg string) {
	if msg != "" {
		mfv.errorLabel.SetText("[red:][bold]❌ " + msg)
	} else {
		mfv.errorLabel.SetText("")
	}
}

// Setup handlers
func (mfv *MaterialFormView) SetOnGuardar(fn func(id int, tipo, titulo string, anio int, disponible bool, extra1, extra2 string)) {
	mfv.onGuardar = fn
}

func (mfv *MaterialFormView) SetOnCancelar(fn func()) {
	mfv.onCancelar = fn
}

// GetPrimitive retorna el elemento raíz de tview.
func (mfv *MaterialFormView) GetPrimitive() tview.Primitive {
	return mfv.layout
}
