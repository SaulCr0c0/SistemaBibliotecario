package view

import (
	"strconv"

	model "biblioteca/Model"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// MaterialTableView es el componente visual para listar y gestionar materiales bibliotecarios.
type MaterialTableView struct {
	layout     *tview.Flex
	table      *tview.Table
	buttons    *tview.Flex
	onNuevo    func()
	onEditar   func(id int)
	onEliminar func(id int)
	onVolver   func()
}

// NewMaterialTableView instancia la vista de tabla de inventario de materiales en alto contraste.
func NewMaterialTableView() *MaterialTableView {
	mtv := &MaterialTableView{}

	mtv.table = CrearTablaEstilizada(" 📖 Inventario de Materiales Bibliotecarios ")

	var inputCapture func(event *tcell.EventKey) *tcell.EventKey
	mtv.buttons, inputCapture = CrearBarraAccionesTabla(
		func() {
			if mtv.onNuevo != nil {
				mtv.onNuevo()
			}
		},
		func(id int) {
			if mtv.onEditar != nil {
				mtv.onEditar(id)
			}
		},
		func(id int) {
			if mtv.onEliminar != nil {
				mtv.onEliminar(id)
			}
		},
		func() {
			if mtv.onVolver != nil {
				mtv.onVolver()
			}
		},
		mtv.ObtenerIDSeleccionado,
	)

	mtv.layout = tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(mtv.table, 0, 1, true).
		AddItem(mtv.buttons, 3, 1, false)

	mtv.layout.SetInputCapture(inputCapture)

	return mtv
}

// CargarMateriales proyecta la lista de materiales prestables en la tabla.
func (mtv *MaterialTableView) CargarMateriales(materiales []model.Prestable) {
	mtv.table.Clear()

	RenderizarEncabezadosTabla(mtv.table, []string{"ID", "Tipo", "Título", "Año", "Estado", "Máx Días", "Detalle Específico"})

	for i, m := range materiales {
		row := i + 1

		cellID := tview.NewTableCell(strconv.Itoa(m.GetID())).SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite)
		cellTipo := tview.NewTableCell(m.GetTipo()).SetAlign(tview.AlignCenter)
		cellTitulo := tview.NewTableCell(m.GetTitulo()).SetAlign(tview.AlignLeft).SetTextColor(tcell.ColorWhite)
		cellAnio := tview.NewTableCell(strconv.Itoa(m.GetAnio())).SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorWhite)

		// Formato visual por tipo
		switch m.GetTipo() {
		case "Libro":
			cellTipo.SetTextColor(tcell.ColorGreen).SetAttributes(tcell.AttrBold)
		case "Tesis":
			cellTipo.SetTextColor(tcell.ColorTeal).SetAttributes(tcell.AttrBold)
		case "Revista":
			cellTipo.SetTextColor(tcell.ColorLightCyan).SetAttributes(tcell.AttrBold)
		default:
			cellTipo.SetTextColor(tcell.ColorWhite)
		}

		// Estado de disponibilidad
		cellEstado := tview.NewTableCell("").SetAlign(tview.AlignCenter)
		if m.EstaDisponible() {
			cellEstado.SetText("Disponible").SetTextColor(tcell.ColorGreen).SetAttributes(tcell.AttrBold)
		} else {
			cellEstado.SetText("Prestado").SetTextColor(tcell.ColorRed).SetAttributes(tcell.AttrBold)
		}

		cellMaxDias := tview.NewTableCell(strconv.Itoa(m.MaxDiasPrestamo()) + " días").SetAlign(tview.AlignCenter).SetTextColor(tcell.ColorYellow)
		cellDetalle := tview.NewTableCell(m.ObtenerDetalle()).SetAlign(tview.AlignLeft).SetTextColor(tcell.ColorLightCyan)

		mtv.table.SetCell(row, 0, cellID)
		mtv.table.SetCell(row, 1, cellTipo)
		mtv.table.SetCell(row, 2, cellTitulo)
		mtv.table.SetCell(row, 3, cellAnio)
		mtv.table.SetCell(row, 4, cellEstado)
		mtv.table.SetCell(row, 5, cellMaxDias)
		mtv.table.SetCell(row, 6, cellDetalle)
	}

	if len(materiales) > 0 {
		mtv.table.Select(1, 0)
	}
}

// ObtenerIDSeleccionado obtiene el ID del material marcado en la fila activa.
func (mtv *MaterialTableView) ObtenerIDSeleccionado() int {
	row, _ := mtv.table.GetSelection()
	if row <= 0 {
		return 0
	}
	cell := mtv.table.GetCell(row, 0)
	if cell == nil {
		return 0
	}
	id, err := strconv.Atoi(cell.Text)
	if err != nil {
		return 0
	}
	return id
}

// Configuración de callbacks
func (mtv *MaterialTableView) SetOnNuevo(fn func())           { mtv.onNuevo = fn }
func (mtv *MaterialTableView) SetOnEditar(fn func(id int))     { mtv.onEditar = fn }
func (mtv *MaterialTableView) SetOnEliminar(fn func(id int))   { mtv.onEliminar = fn }
func (mtv *MaterialTableView) SetOnVolver(fn func())          { mtv.onVolver = fn }

// GetPrimitive retorna el elemento tview principal.
func (mtv *MaterialTableView) GetPrimitive() tview.Primitive {
	return mtv.layout
}
