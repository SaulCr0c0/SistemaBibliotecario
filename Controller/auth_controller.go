package controller

import (
	"biblioteca/Model"
	"biblioteca/View"
)

// AuthController conecta el modelo de gestión de usuarios con las vistas de login y menú principal.
type AuthController struct {
	userMgr  *model.UserManager
	view     *view.LoginView
	router   *view.AppRouter
	mainMenu *view.MainMenuView
}

// NewAuthController crea un nuevo controlador de autenticación e instala los listeners de eventos.
func NewAuthController(userMgr *model.UserManager, loginView *view.LoginView, router *view.AppRouter, mainMenu *view.MainMenuView) *AuthController {
	ctrl := &AuthController{
		userMgr:  userMgr,
		view:     loginView,
		router:   router,
		mainMenu: mainMenu,
	}

	// Escuchar evento de intento de login en la vista
	ctrl.view.SetOnLogin(func(u, p string) {
		ctrl.SubmitLogin(u, p)
	})

	// Escuchar evento de salir en la vista de login
	ctrl.view.SetOnExit(func() {
		ctrl.Salir()
	})

	// Escuchar evento de selección en el menú principal
	if ctrl.mainMenu != nil {
		ctrl.mainMenu.SetOnSelect(func(opc string) {
			if opc == "Cerrar Sesión" {
				ctrl.CerrarSesion()
			} else {
				ctrl.router.MostrarModalError("El módulo [" + opc + "] está en desarrollo.")
			}
		})
	}

	return ctrl
}

// SubmitLogin valida las credenciales recibidas y actualiza la navegación.
func (ac *AuthController) SubmitLogin(u, p string) {
	if u == "" || p == "" {
		ac.view.MostrarError("Por favor ingrese usuario y contraseña.")
		return
	}

	user, err := ac.userMgr.Autenticar(u, p)
	if err != nil {
		ac.view.MostrarError(err.Error())
		return
	}

	// Limpiar mensaje de error si fue exitoso
	ac.view.MostrarError("")

	// Actualizar menú con información del usuario autenticado y conmutar pantalla
	if ac.mainMenu != nil {
		ac.mainMenu.ActualizarInfoUsuario(user.Username, user.Rol)
	}

	ac.router.CambiarPantalla("main_menu")
}

// CerrarSesion quita la sesión activa en el modelo y redirige al login.
func (ac *AuthController) CerrarSesion() {
	ac.userMgr.CerrarSesion()
	ac.router.CambiarPantalla("login")
}

// Salir finaliza la ejecución de la aplicación.
func (ac *AuthController) Salir() {
	ac.router.Salir()
}
