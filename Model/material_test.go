package model

import (
	"testing"
)

func TestMaterialPolimorfismoYMaxDias(t *testing.T) {
	libro := NewLibro(1, "Cien Años de Soledad", 1967, "978-0307474728", "Editorial Sudamericana")
	tesis := NewTesis(2, "Optimización de Algoritmos R-Tree", 2024, "Ingeniería Informática", "Dr. Pérez")
	revista := NewRevista(3, "IEEE Computer", 2025, "0018-9162", 102)

	materiales := []Prestable{libro, tesis, revista}

	if materiales[0].MaxDiasPrestamo() != 14 {
		t.Errorf("Esperado 14 días para libro, obtenido %d", materiales[0].MaxDiasPrestamo())
	}
	if materiales[1].MaxDiasPrestamo() != 7 {
		t.Errorf("Esperado 7 días para tesis, obtenido %d", materiales[1].MaxDiasPrestamo())
	}
	if materiales[2].MaxDiasPrestamo() != 3 {
		t.Errorf("Esperado 3 días para revista, obtenido %d", materiales[2].MaxDiasPrestamo())
	}

	if materiales[0].GetTipo() != "Libro" || materiales[1].GetTipo() != "Tesis" || materiales[2].GetTipo() != "Revista" {
		t.Errorf("Los tipos retornados por GetTipo() no coinciden")
	}
}

func TestMaterialRepoCRUD(t *testing.T) {
	repo := NewMaterialRepo()

	libro := NewLibro(1, "Go en Acción", 2023, "123-456", "Manning")
	tesis := NewTesis(2, "IA en Medicina", 2024, "Medicina", "Dra. Gomez")

	repo.Agregar(libro)
	repo.Agregar(tesis)

	if len(repo.ListarTodos()) != 2 {
		t.Fatalf("Esperados 2 elementos, obtenidos %d", len(repo.ListarTodos()))
	}

	mat, err := repo.BuscarPorID(1)
	if err != nil || mat.GetTitulo() != "Go en Acción" {
		t.Errorf("Error al buscar material por ID: %v", err)
	}

	// Test Actualizar Estado
	err = repo.ActualizarEstado(1, false)
	if err != nil || repo.ListarTodos()[0].EstaDisponible() {
		t.Errorf("Error al actualizar estado de disponibilidad")
	}

	// Test Eliminar
	err = repo.Eliminar(2)
	if err != nil || len(repo.ListarTodos()) != 1 {
		t.Errorf("Error al eliminar material")
	}
}
