package controller

import (
	"strconv"

	model "biblioteca/Model"
	view "biblioteca/View"
)

// MaterialController coordina las acciones entre el repositorio de materiales y la vista tview.
type MaterialController struct {
	repo      *model.MaterialRepo
	tableView *view.MaterialTableView
	formView  *view.MaterialFormView
	router    *view.AppRouter
}

// NewMaterialController crea una instancia del controlador de materiales y enlaza los eventos de vista.
func NewMaterialController(
	repo *model.MaterialRepo,
	tableView *view.MaterialTableView,
	formView *view.MaterialFormView,
	router *view.AppRouter,
) *MaterialController {
	ctrl := &MaterialController{
		repo:      repo,
		tableView: tableView,
		formView:  formView,
		router:    router,
	}

	ctrl.setupListeners()
	return ctrl
}

func (mc *MaterialController) setupListeners() {
	if mc.tableView != nil {
		mc.tableView.SetOnNuevo(func() {
			mc.AbrirFormularioCrear()
		})

		mc.tableView.SetOnEditar(func(id int) {
			mc.AbrirFormularioEditar(id)
		})

		mc.tableView.SetOnEliminar(func(id int) {
			mc.EliminarMaterial(id)
		})

		mc.tableView.SetOnVolver(func() {
			mc.router.CambiarPantalla("main_menu")
		})
	}

	if mc.formView != nil {
		mc.formView.SetOnGuardar(func(id int, tipo, titulo string, anio int, disponible bool, extra1, extra2 string) {
			mc.GuardarMaterial(id, tipo, titulo, anio, disponible, extra1, extra2)
		})

		mc.formView.SetOnCancelar(func() {
			mc.MostrarInventario()
		})
	}
}

// MostrarInventario refresca la tabla de materiales y conmuta a la vista del inventario.
func (mc *MaterialController) MostrarInventario() {
	materiales := mc.repo.ListarTodos()
	if mc.tableView != nil {
		mc.tableView.CargarMateriales(materiales)
	}
	if mc.router != nil {
		mc.router.CambiarPantalla("material_table")
	}
}

// AbrirFormularioCrear limpia la vista del formulario y conmuta a la vista de creación.
func (mc *MaterialController) AbrirFormularioCrear() {
	if mc.formView != nil {
		mc.formView.CargarDatosModoCrear()
	}
	if mc.router != nil {
		mc.router.CambiarPantalla("material_form")
	}
}

// AbrirFormularioEditar puebla el formulario con el material a editar y cambia de pantalla.
func (mc *MaterialController) AbrirFormularioEditar(id int) {
	material, err := mc.repo.BuscarPorID(id)
	if err != nil {
		if mc.router != nil {
			mc.router.MostrarModalError(err.Error())
		}
		return
	}
	if mc.formView != nil {
		mc.formView.CargarDatosModoEditar(material)
	}
	if mc.router != nil {
		mc.router.CambiarPantalla("material_form")
	}
}

// GuardarMaterial procesa la creación o actualización de un material en el repositorio.
func (mc *MaterialController) GuardarMaterial(id int, tipo, titulo string, anio int, disponible bool, extra1, extra2 string) {
	var nuevoMaterial model.Prestable

	if id == 0 {
		nuevoID := 1
		todos := mc.repo.ListarTodos()
		for _, item := range todos {
			if item.GetID() >= nuevoID {
				nuevoID = item.GetID() + 1
			}
		}
		id = nuevoID
	}

	switch tipo {
	case "Libro":
		l := model.NewLibro(id, titulo, anio, extra1, extra2)
		l.SetDisponible(disponible)
		nuevoMaterial = l
	case "Tesis":
		t := model.NewTesis(id, titulo, anio, extra1, extra2)
		t.SetDisponible(disponible)
		nuevoMaterial = t
	case "Revista":
		edicion, _ := strconv.Atoi(extra2)
		r := model.NewRevista(id, titulo, anio, extra1, edicion)
		r.SetDisponible(disponible)
		nuevoMaterial = r
	default:
		if mc.formView != nil {
			mc.formView.MostrarError("Tipo de material no válido")
		}
		return
	}

	var err error
	_, errFind := mc.repo.BuscarPorID(id)
	if errFind == nil {
		err = mc.repo.Actualizar(nuevoMaterial)
	} else {
		mc.repo.Agregar(nuevoMaterial)
	}

	if err != nil {
		if mc.formView != nil {
			mc.formView.MostrarError(err.Error())
		}
		return
	}

	mc.MostrarInventario()
}

// EliminarMaterial remueve un material por su ID.
func (mc *MaterialController) EliminarMaterial(id int) {
	err := mc.repo.Eliminar(id)
	if err != nil {
		if mc.router != nil {
			mc.router.MostrarModalError(err.Error())
		}
		return
	}
	mc.MostrarInventario()
}
