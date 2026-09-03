package controller

import (
	"biblioteca/View"
)

// LibraryController es el controlador principal que orquesta la transición entre vistas y eventos.
type LibraryController struct {
	router       *view.AppRouter
	authCtrl     *AuthController
	userCtrl     *UserController
	materialCtrl *MaterialController
	loginView    *view.LoginView
	mainMenu     *view.MainMenuView
}

// NewLibraryController crea el controlador principal e instala los escuchadores de navegación.
func NewLibraryController(
	router *view.AppRouter,
	authCtrl *AuthController,
	userCtrl *UserController,
	loginView *view.LoginView,
	mainMenu *view.MainMenuView,
) *LibraryController {
	ctrl := &LibraryController{
		router:    router,
		authCtrl:  authCtrl,
		userCtrl:  userCtrl,
		loginView: loginView,
		mainMenu:  mainMenu,
	}

	ctrl.setupListeners()
	return ctrl
}

// SetUserController asigna la referencia al controlador de gestión de usuarios.
func (lc *LibraryController) SetUserController(userCtrl *UserController) {
	lc.userCtrl = userCtrl
}

// SetMaterialController asigna la referencia al controlador de gestión de materiales.
func (lc *LibraryController) SetMaterialController(materialCtrl *MaterialController) {
	lc.materialCtrl = materialCtrl
}

// setupListeners configura los callbacks de eventos de las distintas vistas.
func (lc *LibraryController) setupListeners() {
	if lc.loginView != nil {
		lc.loginView.SetOnLogin(func(u, p string) {
			lc.handleLogin(u, p)
		})

		lc.loginView.SetOnExit(func() {
			lc.router.Salir()
		})
	}

	if lc.mainMenu != nil {
		lc.mainMenu.SetOnSelect(func(opc string) {
			switch opc {
			case "Cerrar Sesión":
				lc.handleLogout()
			case "Gestionar Usuarios":
				lc.handleGestionUsuarios()
			case "Gestionar Materiales":
				lc.handleGestionMateriales()
			default:
				lc.router.MostrarModalError("El módulo [" + opc + "] está en desarrollo.")
			}
		})
	}
}

// handleLogin procesa el formulario de acceso y ejecuta la transición al menú principal.
func (lc *LibraryController) handleLogin(u, p string) {
	if u == "" || p == "" {
		lc.loginView.MostrarError("Por favor ingrese usuario y contraseña.")
		return
	}

	user, err := lc.authCtrl.Autenticar(u, p)
	if err != nil {
		lc.loginView.MostrarError(err.Error())
		return
	}

	// Éxito: limpiar error, actualizar vista del menú según el rol y conmutar pantalla
	lc.loginView.MostrarError("")
	if lc.mainMenu != nil {
		lc.mainMenu.ActualizarInfoUsuario(user.Username, user.Rol)
	}

	lc.router.CambiarPantalla("main_menu")
}

// handleGestionUsuarios valida los permisos del rol antes de dar acceso al módulo de usuarios.
func (lc *LibraryController) handleGestionUsuarios() {
	user := lc.authCtrl.GetSesion()
	if user != nil && user.Rol == "Administrador" {
		if lc.userCtrl != nil {
			lc.userCtrl.MostrarUsuarios()
		} else {
			lc.router.MostrarModalError("Módulo de usuarios no inicializado.")
		}
	} else {
		lc.router.MostrarModalError("Acceso Denegado: Únicamente el Administrador puede gestionar usuarios.")
	}
}

// handleGestionMateriales permite a los usuarios con rol adecuado ingresar a la gestión de inventario.
func (lc *LibraryController) handleGestionMateriales() {
	if lc.materialCtrl != nil {
		lc.materialCtrl.MostrarInventario()
	} else {
		lc.router.MostrarModalError("Módulo de materiales no inicializado.")
	}
}

// handleLogout procesa el cierre de sesión, limpia la pantalla de login y regresa al acceso.
func (lc *LibraryController) handleLogout() {
	lc.authCtrl.CerrarSesion()
	if lc.loginView != nil {
		lc.loginView.LimpiarCampos()
	}
	lc.router.CambiarPantalla("login")
}
