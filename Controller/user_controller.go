package controller

import (
	model "biblioteca/Model"
	view "biblioteca/View"
)

// UserController gestiona las acciones de la tabla y formulario de usuarios.
type UserController struct {
	userMgr   *model.UserManager
	tableView *view.UserTableView
	formView  *view.UserFormView
	router    *view.AppRouter
}

// NewUserController crea una nueva instancia del controlador de usuarios.
func NewUserController(
	userMgr *model.UserManager,
	tableView *view.UserTableView,
	formView *view.UserFormView,
	router *view.AppRouter,
) *UserController {
	ctrl := &UserController{
		userMgr:   userMgr,
		tableView: tableView,
		formView:  formView,
		router:    router,
	}

	ctrl.setupListeners()
	return ctrl
}

func (uc *UserController) setupListeners() {
	if uc.tableView != nil {
		uc.tableView.SetOnNuevo(func() {
			uc.AbrirFormularioCrear()
		})

		uc.tableView.SetOnEditar(func(id int) {
			uc.AbrirFormularioEditar(id)
		})

		uc.tableView.SetOnEliminar(func(id int) {
			uc.EliminarUsuario(id)
		})

		uc.tableView.SetOnVolver(func() {
			uc.router.CambiarPantalla("main_menu")
		})
	}

	if uc.formView != nil {
		uc.formView.SetOnGuardar(func(id int, username, password, rol string, nivelAcceso int, turno string) {
			uc.GuardarUsuario(id, username, password, rol, nivelAcceso, turno)
		})

		uc.formView.SetOnCancelar(func() {
			uc.MostrarUsuarios()
		})
	}
}

// MostrarUsuarios refresca los datos de la tabla y conmuta a la pantalla de lista de usuarios.
func (uc *UserController) MostrarUsuarios() {
	usuarios := uc.userMgr.ListarUsuarios()
	uc.tableView.CargarUsuarios(usuarios)
	uc.router.CambiarPantalla("user_table")
}

// AbrirFormularioCrear limpia el formulario y conmuta a la pantalla de creación.
func (uc *UserController) AbrirFormularioCrear() {
	uc.formView.CargarDatosModoCrear()
	uc.router.CambiarPantalla("user_form")
}

// AbrirFormularioEditar puebla el formulario con el usuario seleccionado y conmuta a la pantalla de edición.
func (uc *UserController) AbrirFormularioEditar(id int) {
	user, err := uc.userMgr.BuscarPorID(id)
	if err != nil {
		uc.router.MostrarModalError(err.Error())
		return
	}
	uc.formView.CargarDatosModoEditar(user)
	uc.router.CambiarPantalla("user_form")
}

// GuardarUsuario procesa el guardado (creación o actualización) de un usuario.
func (uc *UserController) GuardarUsuario(id int, username, password, rol string, nivelAcceso int, turno string) {
	var err error
	if id == 0 {
		_, err = uc.userMgr.CrearUsuario(username, password, rol, nivelAcceso, turno)
	} else {
		err = uc.userMgr.ActualizarUsuario(id, username, password, rol, nivelAcceso, turno)
	}

	if err != nil {
		uc.formView.MostrarError(err.Error())
		return
	}

	uc.MostrarUsuarios()
}

// EliminarUsuario procesa el borrado del usuario seleccionado tras validación.
func (uc *UserController) EliminarUsuario(id int) {
	err := uc.userMgr.EliminarUsuario(id)
	if err != nil {
		uc.router.MostrarModalError(err.Error())
		return
	}
	uc.MostrarUsuarios()
}
