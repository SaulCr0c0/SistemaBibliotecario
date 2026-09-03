package controller

import (
	"testing"
	"biblioteca/Model"
)

func TestAuthController_Autenticar(t *testing.T) {
	userMgr := model.NewUserManager()
	admin := model.NewAdministrador(1, "admin", "admin123", "Administrador", 1)
	userMgr.AgregarUsuario(&admin.Usuario)

	authCtrl := NewAuthController(userMgr)

	// Auth success
	user, err := authCtrl.Autenticar("admin", "admin123")
	if err != nil {
		t.Fatalf("Error inesperado en autenticación: %v", err)
	}
	if user.Username != "admin" {
		t.Errorf("Esperaba usuario 'admin', obtuvo '%s'", user.Username)
	}

	// Auth failure
	_, err = authCtrl.Autenticar("admin", "wrong")
	if err == nil {
		t.Errorf("Esperaba error con contraseña incorrecta")
	}

	// Session logout
	authCtrl.CerrarSesion()
	if authCtrl.GetSesion() != nil {
		t.Errorf("Esperaba sesión nula tras cerrar sesión")
	}
}
