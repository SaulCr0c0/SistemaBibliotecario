package model

import (
	"testing"
)

func TestUserManager_Autenticar(t *testing.T) {
	mgr := NewUserManager()

	admin := NewAdministrador(1, "admin", "admin123", "Administrador", 1)
	biblio := NewBibliotecario(2, "biblio", "biblio123", "Bibliotecario", "Mañana")

	mgr.AgregarUsuario(&admin.Usuario)
	mgr.AgregarUsuario(&biblio.Usuario)

	// Test 1: Successful Auth
	u, err := mgr.Autenticar("admin", "admin123")
	if err != nil {
		t.Fatalf("Esperaba autenticación exitosa, obtuvo error: %v", err)
	}
	if u.Username != "admin" {
		t.Errorf("Esperaba usuario 'admin', obtuvo '%s'", u.Username)
	}
	if mgr.GetSesion() == nil || mgr.GetSesion().Username != "admin" {
		t.Errorf("Sesión activa no fue asignada correctamente")
	}

	// Test 2: Invalid Password
	_, err = mgr.Autenticar("admin", "wrongpass")
	if err == nil {
		t.Errorf("Esperaba error por contraseña incorrecta")
	}

	// Test 3: Non-existent User
	_, err = mgr.Autenticar("nonexistent", "admin123")
	if err == nil {
		t.Errorf("Esperaba error por usuario no existente")
	}

	// Test 4: Logout
	mgr.CerrarSesion()
	if mgr.GetSesion() != nil {
		t.Errorf("Esperaba sesión nula tras cerrar sesión")
	}
}
