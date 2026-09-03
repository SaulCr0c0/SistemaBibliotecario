package view

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// CrearCampoTexto genera un InputField estilizado con colores de alto contraste.
func CrearCampoTexto(label string, width int) *tview.InputField {
	return tview.NewInputField().
		SetLabel(label).
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetFieldWidth(width)
}

// CrearCampoPassword genera un InputField para contraseñas con máscara '*'.
func CrearCampoPassword(label string, width int) *tview.InputField {
	field := CrearCampoTexto(label, width)
	field.SetMaskCharacter('*')
	return field
}

// CrearDesplegable genera un DropDown estilizado con colores de alto contraste.
func CrearDesplegable(label string, opciones []string, onSelect func(option string, index int)) *tview.DropDown {
	return tview.NewDropDown().
		SetLabel(label).
		SetLabelColor(tcell.ColorYellow).
		SetFieldBackgroundColor(tcell.ColorDarkBlue).
		SetFieldTextColor(tcell.ColorWhite).
		SetListStyles(
			tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorWhite),
			tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorYellow).Bold(true),
		).
		SetOptions(opciones, onSelect)
}

// EstilarFormulario aplica el tema de color, bordes y título a un Form.
func EstilarFormulario(form *tview.Form, titulo string) {
	form.SetButtonBackgroundColor(tcell.ColorBlue)
	form.SetButtonTextColor(tcell.ColorYellow)
	form.SetFieldBackgroundColor(tcell.ColorDarkBlue)
	form.SetFieldTextColor(tcell.ColorWhite)
	form.SetLabelColor(tcell.ColorYellow)

	form.SetBorder(true)
	form.SetBorderColor(tcell.ColorTeal)
	form.SetTitle(titulo)
	form.SetTitleColor(tcell.ColorYellow)
	form.SetTitleAlign(tview.AlignCenter)
}

// CrearLayoutCentrado envuelve un Form y un TextView de error en un Flex centrado.
func CrearLayoutCentrado(form *tview.Form, errorLabel *tview.TextView, formWidth, formHeight int) *tview.Flex {
	return tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexColumn).
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(form, formHeight, 1, true).
				AddItem(errorLabel, 2, 1, false), formWidth, 1, true).
			AddItem(nil, 0, 1, false), formHeight+2, 1, true).
		AddItem(nil, 0, 1, false)
}

// CrearTablaEstilizada genera una tabla configurada con bordes y título en alto contraste.
func CrearTablaEstilizada(titulo string) *tview.Table {
	table := tview.NewTable().
		SetBorders(true).
		SetBordersColor(tcell.ColorTeal).
		SetSelectable(true, false).
		SetSelectedStyle(tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorYellow).Bold(true))

	table.SetBorder(true).
		SetBorderColor(tcell.ColorTeal).
		SetTitle(titulo).
		SetTitleColor(tcell.ColorYellow).
		SetTitleAlign(tview.AlignCenter)

	return table
}

// RenderizarEncabezadosTabla coloca las celdas de encabezado amarillas y en negrita en la fila 0.
func RenderizarEncabezadosTabla(table *tview.Table, headers []string) {
	for col, h := range headers {
		cell := tview.NewTableCell(h).
			SetTextColor(tcell.ColorYellow).
			SetAttributes(tcell.AttrBold).
			SetSelectable(false).
			SetAlign(tview.AlignCenter)
		table.SetCell(0, col, cell)
	}
}

// CrearBarraAccionesTabla instancia los botones [A] Agregar, [E] Editar, [D] Eliminar y [V] Volver
// y configura la captura de teclado para shortcuts a, e, d, v.
func CrearBarraAccionesTabla(
	onNuevo func(),
	onEditar func(id int),
	onEliminar func(id int),
	onVolver func(),
	getIDSeleccionado func() int,
) (*tview.Flex, func(event *tcell.EventKey) *tcell.EventKey) {
	btnNuevo := tview.NewButton(" [A] Agregar ").SetSelectedFunc(func() {
		if onNuevo != nil {
			onNuevo()
		}
	})

	btnEditar := tview.NewButton(" [E] Editar ").SetSelectedFunc(func() {
		if onEditar != nil {
			id := getIDSeleccionado()
			if id > 0 {
				onEditar(id)
			}
		}
	})

	btnEliminar := tview.NewButton(" [D] Eliminar ").SetSelectedFunc(func() {
		if onEliminar != nil {
			id := getIDSeleccionado()
			if id > 0 {
				onEliminar(id)
			}
		}
	})

	btnVolver := tview.NewButton(" [V] Volver ").SetSelectedFunc(func() {
		if onVolver != nil {
			onVolver()
		}
	})

	for _, btn := range []*tview.Button{btnNuevo, btnEditar, btnEliminar, btnVolver} {
		btn.SetStyle(tcell.StyleDefault.Background(tcell.ColorDarkBlue).Foreground(tcell.ColorYellow).Bold(true)).
			SetActivatedStyle(tcell.StyleDefault.Background(tcell.ColorYellow).Foreground(tcell.ColorBlack).Bold(true))
	}

	buttons := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(btnNuevo, 0, 1, false).
		AddItem(btnEditar, 0, 1, false).
		AddItem(btnEliminar, 0, 1, false).
		AddItem(btnVolver, 0, 1, false)

	inputCapture := func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a', 'A':
			if onNuevo != nil {
				onNuevo()
				return nil
			}
		case 'e', 'E':
			if onEditar != nil {
				id := getIDSeleccionado()
				if id > 0 {
					onEditar(id)
				}
				return nil
			}
		case 'd', 'D':
			if onEliminar != nil {
				id := getIDSeleccionado()
				if id > 0 {
					onEliminar(id)
				}
				return nil
			}
		case 'v', 'V':
			if onVolver != nil {
				onVolver()
				return nil
			}
		}
		return event
	}

	return buttons, inputCapture
}
