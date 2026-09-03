package view

import (
	"testing"
)

func TestMainMenuView_ConfigurarMenuPorRol(t *testing.T) {
	menuView := NewMainMenuView()

	// Case 1: Administrador role
	menuView.ConfigurarMenuPorRol("Administrador")
	if menuView.menu.GetItemCount() != 6 {
		t.Errorf("Esperaba 6 opciones en el menú para Administrador, obtuvo: %d", menuView.menu.GetItemCount())
	}

	mainText, _ := menuView.menu.GetItemText(0)
	if mainText != "Gestionar Usuarios" {
		t.Errorf("Esperaba la primera opción 'Gestionar Usuarios' para Administrador, obtuvo: %s", mainText)
	}

	// Case 2: Bibliotecario role
	menuView.ConfigurarMenuPorRol("Bibliotecario")
	if menuView.menu.GetItemCount() != 5 {
		t.Errorf("Esperaba 5 opciones en el menú para Bibliotecario, obtuvo: %d", menuView.menu.GetItemCount())
	}

	mainTextBiblio, _ := menuView.menu.GetItemText(0)
	if mainTextBiblio == "Gestionar Usuarios" {
		t.Errorf("Bibliotecario NO debe tener la opción 'Gestionar Usuarios'")
	}
}
