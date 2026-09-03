package controller

import (
	view "biblioteca/View"
)

// LibraryController es el controlador principal que orquesta la transición entre vistas y eventos.
type LibraryController struct {
	router    *view.AppRouter
	authCtrl  *AuthController
	loginView *view.LoginView
	mainMenu  *view.MainMenuView
}

// NewLibraryController crea el controlador principal e instala los escuchadores de navegación.
func NewLibraryController(
	router *view.AppRouter,
	authCtrl *AuthController,
	loginView *view.LoginView,
	mainMenu *view.MainMenuView,
) *LibraryController {
	ctrl := &LibraryController{
		router:    router,
		authCtrl:  authCtrl,
		loginView: loginView,
		mainMenu:  mainMenu,
	}
	ctrl.setupListeners()
	return ctrl
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
			if opc == "Cerrar Sesión" {
				lc.handleLogout()
			} else {
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
	// Éxito: limpiar error, actualizar información en la vista del menú y conmutar pantalla
	lc.loginView.MostrarError("")
	if lc.mainMenu != nil {
		lc.mainMenu.ActualizarInfoUsuario(user.Username, user.Rol)
	}
	lc.router.CambiarPantalla("main_menu")
}

// handleLogout procesa el cierre de sesión, limpia la pantalla de login y regresa al acceso.
func (lc *LibraryController) handleLogout() {
	lc.authCtrl.CerrarSesion()
	if lc.loginView != nil {
		lc.loginView.LimpiarCampos()
	}
	lc.router.CambiarPantalla("login")
}
