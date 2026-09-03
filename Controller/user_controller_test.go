package controller

import (
	"testing"

	"biblioteca/Model"
	"biblioteca/View"
)

func TestUserController_Flow(t *testing.T) {
	userMgr := model.NewUserManager()
	router := view.NewAppRouter()
	tableView := view.NewUserTableView()
	formView := view.NewUserFormView()

	router.AgregarPantalla("user_table", tableView.GetPrimitive())
	router.AgregarPantalla("user_form", formView.GetPrimitive())

	userCtrl := NewUserController(userMgr, tableView, formView, router)

	// Crear usuario vía controller
	userCtrl.GuardarUsuario(0, "testadmin", "pass123", "Administrador", 3, "")

	usuarios := userMgr.ListarUsuarios()
	if len(usuarios) != 1 {
		t.Fatalf("Esperaba 1 usuario creado, obtuvo %d", len(usuarios))
	}

	u := usuarios[0]
	if u.Username != "testadmin" || u.Rol != "Administrador" {
		t.Errorf("Datos de usuario no coinciden")
	}

	// Modificar usuario vía controller
	userCtrl.GuardarUsuario(u.ID, "testadmin_edit", "pass123", "Administrador", 8, "")

	if usuarios[0].Username != "testadmin_edit" || usuarios[0].NivelAcceso != 8 {
		t.Errorf("Usuario no actualizado correctamente vía controller")
	}

	// Eliminar usuario vía controller
	userCtrl.EliminarUsuario(u.ID)
	if len(userMgr.ListarUsuarios()) != 0 {
		t.Errorf("Esperaba 0 usuarios tras eliminar")
	}
}
