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

func TestUserManager_CRUD(t *testing.T) {
	mgr := NewUserManager()

	// 1. Crear Usuario
	u1, err := mgr.CrearUsuario("user1", "pass123", "Bibliotecario", 0, "Tarde")
	if err != nil {
		t.Fatalf("Error creando usuario: %v", err)
	}
	if u1.ID != 1 || len(mgr.ListarUsuarios()) != 1 {
		t.Errorf("Esperaba ID 1 y total 1 usuario")
	}

	// 2. Buscar por ID
	found, err := mgr.BuscarPorID(1)
	if err != nil || found.Username != "user1" {
		t.Errorf("Esperaba encontrar usuario 'user1'")
	}

	// 3. Actualizar Usuario
	err = mgr.ActualizarUsuario(1, "user1_mod", "newpass", "Administrador", 5, "")
	if err != nil {
		t.Fatalf("Error actualizando usuario: %v", err)
	}

	updated, _ := mgr.BuscarPorID(1)
	if updated.Username != "user1_mod" || updated.Rol != "Administrador" || updated.NivelAcceso != 5 {
		t.Errorf("Usuario no fue actualizado correctamente")
	}

	// 4. Eliminar Usuario
	err = mgr.EliminarUsuario(1)
	if err != nil {
		t.Fatalf("Error eliminando usuario: %v", err)
	}

	if len(mgr.ListarUsuarios()) != 0 {
		t.Errorf("Esperaba lista vacía tras eliminar usuario")
	}
}
