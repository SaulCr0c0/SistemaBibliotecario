package controller

import (
	"biblioteca/Model"
)

// AuthController se encarga exclusivamente de la lógica de autenticación y sesión de usuario.
type AuthController struct {
	userMgr *model.UserManager
}

// NewAuthController inicializa el controlador de autenticación con el gestor de usuarios.
func NewAuthController(userMgr *model.UserManager) *AuthController {
	return &AuthController{
		userMgr: userMgr,
	}
}

// Autenticar valida el usuario y contraseña delegando al UserManager.
func (ac *AuthController) Autenticar(username, password string) (*model.Usuario, error) {
	return ac.userMgr.Autenticar(username, password)
}

// CerrarSesion quita la referencia del usuario de la sesión activa.
func (ac *AuthController) CerrarSesion() {
	ac.userMgr.CerrarSesion()
}

// GetSesion obtiene el usuario actualmente autenticado en el sistema.
func (ac *AuthController) GetSesion() *model.Usuario {
	return ac.userMgr.GetSesion()
}
