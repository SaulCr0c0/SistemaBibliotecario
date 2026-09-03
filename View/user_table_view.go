package view

import (
	"strconv"

	model "biblioteca/Model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// UserTableView es el componente visual para listar y gestionar usuarios en una tabla con alto contraste.
type UserTableView struct {
	layout     *tview.Flex
	table      *tview.Table
	buttons    *tview.Flex
	onNuevo    func()
	onEditar   func(id int)
	onEliminar func(id int)
	onVolver   func()
}

// NewUserTableView instancia la tabla de gestión de usuarios con colores de alto contraste.
func NewUserTableView() *UserTableView {
	utv := &UserTableView{}

	utv.table = CrearTablaEstilizada(" 👤 Gestión de Usuarios - Administrador ")

	var inputCapture func(event *tcell.EventKey) *tcell.EventKey
	utv.buttons, inputCapture = CrearBarraAccionesTabla(
		func() {
			if utv.onNuevo != nil {
				utv.onNuevo()
			}
		},
		func(id int) {
			if utv.onEditar != nil {
				utv.onEditar(id)
			}
		},
		func(id int) {
			if utv.onEliminar != nil {
				utv.onEliminar(id)
			}
		},
		func() {
			if utv.onVolver != nil {
				utv.onVolver()
			}
		},
		utv.ObtenerIDSeleccionado,
	)

	utv.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(utv.table, 0, 1, true).
		AddItem(utv.buttons, 3, 1, false)

	utv.layout.SetInputCapture(inputCapture)

	return utv
}

// CargarUsuarios renderiza la lista de usuarios recibida en la tabla.
func (utv *UserTableView) CargarUsuarios(usuarios []*model.Usuario) {
	utv.table.Clear()

	RenderizarEncabezadosTabla(utv.table, []string{"ID", "Nombre de Usuario", "Rol", "Detalle Específico"})

	for i, u := range usuarios {
		row := i + 1

		cellID := tview.NewTableCell(strconv.Itoa(u.ID)).SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite)
		cellUsername := tview.NewTableCell(u.Username).SetAlign(tview.AlignLeft).SetTextColor(tcell.ColorWhite)
		cellRol := tview.NewTableCell(u.Rol).SetAlign(tview.AlignCenter)
		cellDetalle := tview.NewTableCell(u.ObtenerDetalle()).SetAlign(tview.AlignLeft).SetTextColor(tcell.ColorLightCyan)

		if u.Rol == "Administrador" {
			cellRol.SetTextColor(tcell.ColorTeal).SetAttributes(tcell.AttrBold)
		} else {
			cellRol.SetTextColor(tcell.ColorGreen).SetAttributes(tcell.AttrBold)
		}

		utv.table.SetCell(row, 0, cellID)
		utv.table.SetCell(row, 1, cellUsername)
		utv.table.SetCell(row, 2, cellRol)
		utv.table.SetCell(row, 3, cellDetalle)
	}

	if len(usuarios) > 0 {
		utv.table.Select(1, 0)
	}
}

// ObtenerIDSeleccionado retorna el ID del usuario correspondiente a la fila seleccionada.
func (utv *UserTableView) ObtenerIDSeleccionado() int {
	row, _ := utv.table.GetSelection()
	if row <= 0 {
		return 0
	}
	cell := utv.table.GetCell(row, 0)
	if cell == nil {
		return 0
	}
	id, err := strconv.Atoi(cell.Text)
	if err != nil {
		return 0
	}
	return id
}

// SetOnNuevo asigna el callback para crear un usuario.
func (utv *UserTableView) SetOnNuevo(fn func()) {
	utv.onNuevo = fn
}

// SetOnEditar asigna el callback para editar el usuario seleccionado.
func (utv *UserTableView) SetOnEditar(fn func(id int)) {
	utv.onEditar = fn
}

// SetOnEliminar asigna el callback para eliminar el usuario seleccionado.
func (utv *UserTableView) SetOnEliminar(fn func(id int)) {
	utv.onEliminar = fn
}

// SetOnVolver asigna el callback para regresar al menú principal.
func (utv *UserTableView) SetOnVolver(fn func()) {
	utv.onVolver = fn
}

// GetPrimitive retorna el elemento raíz de la vista de tabla.
func (utv *UserTableView) GetPrimitive() tview.Primitive {
	return utv.layout
}
