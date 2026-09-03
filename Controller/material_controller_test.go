package controller

import (
	"testing"

	model "biblioteca/Model"
)

func TestMaterialController_GuardarYEliminar(t *testing.T) {
	repo := model.NewMaterialRepo()
	ctrl := NewMaterialController(repo, nil, nil, nil)

	// Crear Libro
	ctrl.GuardarMaterial(0, "Libro", "Don Quijote de la Mancha", 1605, true, "978-8424115456", "Francisco de Robles")
	if len(repo.ListarTodos()) != 1 {
		t.Fatalf("Esperado 1 material guardado, obtenidos %d", len(repo.ListarTodos()))
	}

	mat, err := repo.BuscarPorID(1)
	if err != nil || mat.GetTipo() != "Libro" || mat.GetTitulo() != "Don Quijote de la Mancha" {
		t.Errorf("Error al verificar creación de libro: %v", err)
	}

	// Crear Tesis
	ctrl.GuardarMaterial(0, "Tesis", "Sistemas Distribuidos", 2024, true, "Sistemas", "Ing. Garcia")
	if len(repo.ListarTodos()) != 2 {
		t.Fatalf("Esperados 2 materiales guardados, obtenidos %d", len(repo.ListarTodos()))
	}

	// Editar Libro
	ctrl.GuardarMaterial(1, "Libro", "Don Quijote (Edición Revisada)", 1605, false, "978-8424115456", "Editorial Planeta")
	matEdit, err := repo.BuscarPorID(1)
	if err != nil || matEdit.GetTitulo() != "Don Quijote (Edición Revisada)" || matEdit.EstaDisponible() {
		t.Errorf("Error al verificar edición de libro")
	}

	// Eliminar Tesis
	ctrl.EliminarMaterial(2)
	if len(repo.ListarTodos()) != 1 {
		t.Errorf("Esperado 1 material tras eliminación, obtenidos %d", len(repo.ListarTodos()))
	}
}
