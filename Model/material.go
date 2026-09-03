package model

import (
	"errors"
	"fmt"
)

// Prestable define la interfaz común para todos los materiales bibliotecarios prestables.
type Prestable interface {
	GetID() int
	GetTitulo() string
	GetAnio() int
	EstaDisponible() bool
	SetDisponible(d bool)
	GetTipo() string
	MaxDiasPrestamo() int
	ObtenerDetalle() string
}

// MaterialBase es la estructura base con campos comunes para todo material.
type MaterialBase struct {
	ID         int
	Titulo     string
	Anio       int
	Disponible bool
}

// GetID retorna el identificador único del material.
func (m *MaterialBase) GetID() int {
	return m.ID
}

// GetTitulo retorna el título del material.
func (m *MaterialBase) GetTitulo() string {
	return m.Titulo
}

// GetAnio retorna el año de publicación del material.
func (m *MaterialBase) GetAnio() int {
	return m.Anio
}

// EstaDisponible indica si el material está libre para préstamo.
func (m *MaterialBase) EstaDisponible() bool {
	return m.Disponible
}

// SetDisponible modifica la disponibilidad del material.
func (m *MaterialBase) SetDisponible(d bool) {
	m.Disponible = d
}

// Libro representa un libro físico en la biblioteca.
type Libro struct {
	MaterialBase
	ISBN      string
	Editorial string
}

// NewLibro crea una instancia de Libro con sus datos y disponible por defecto.
func NewLibro(id int, titulo string, anio int, isbn, editorial string) *Libro {
	return &Libro{
		MaterialBase: MaterialBase{
			ID:         id,
			Titulo:     titulo,
			Anio:       anio,
			Disponible: true,
		},
		ISBN:      isbn,
		Editorial: editorial,
	}
}

// GetTipo retorna el tipo "Libro".
func (l *Libro) GetTipo() string {
	return "Libro"
}

// MaxDiasPrestamo retorna el límite máximo de días de préstamo para un libro (14 días).
func (l *Libro) MaxDiasPrestamo() int {
	return 14
}

// ObtenerDetalle retorna una descripción con ISBN y Editorial.
func (l *Libro) ObtenerDetalle() string {
	return fmt.Sprintf("ISBN: %s | Editorial: %s", l.ISBN, l.Editorial)
}

// Tesis representa una tesis académica de grado o posgrado.
type Tesis struct {
	MaterialBase
	Carrera string
	Asesor  string
}

// NewTesis crea una instancia de Tesis con sus datos y disponible por defecto.
func NewTesis(id int, titulo string, anio int, carrera, asesor string) *Tesis {
	return &Tesis{
		MaterialBase: MaterialBase{
			ID:         id,
			Titulo:     titulo,
			Anio:       anio,
			Disponible: true,
		},
		Carrera: carrera,
		Asesor:  asesor,
	}
}

// GetTipo retorna el tipo "Tesis".
func (t *Tesis) GetTipo() string {
	return "Tesis"
}

// MaxDiasPrestamo retorna el límite máximo de días de préstamo para una tesis (7 días).
func (t *Tesis) MaxDiasPrestamo() int {
	return 7
}

// ObtenerDetalle retorna una descripción con Carrera y Asesor.
func (t *Tesis) ObtenerDetalle() string {
	return fmt.Sprintf("Carrera: %s | Asesor: %s", t.Carrera, t.Asesor)
}

// Revista representa una publicación periódica o científica.
type Revista struct {
	MaterialBase
	ISSN    string
	Edicion int
}

// NewRevista crea una instancia de Revista con sus datos y disponible por defecto.
func NewRevista(id int, titulo string, anio int, issn string, edicion int) *Revista {
	return &Revista{
		MaterialBase: MaterialBase{
			ID:         id,
			Titulo:     titulo,
			Anio:       anio,
			Disponible: true,
		},
		ISSN:    issn,
		Edicion: edicion,
	}
}

// GetTipo retorna el tipo "Revista".
func (r *Revista) GetTipo() string {
	return "Revista"
}

// MaxDiasPrestamo retorna el límite máximo de días de préstamo para una revista (3 días).
func (r *Revista) MaxDiasPrestamo() int {
	return 3
}

// ObtenerDetalle retorna una descripción con ISSN y Edición.
func (r *Revista) ObtenerDetalle() string {
	return fmt.Sprintf("ISSN: %s | Edición: %d", r.ISSN, r.Edicion)
}

// MaterialRepo gestiona la colección de materiales bibliotecarios prestables en memoria.
type MaterialRepo struct {
	items []Prestable
}

// NewMaterialRepo inicializa un repositorio de materiales vacío.
func NewMaterialRepo() *MaterialRepo {
	return &MaterialRepo{
		items: make([]Prestable, 0),
	}
}

// Agregar almacena un nuevo material en el repositorio.
func (r *MaterialRepo) Agregar(m Prestable) {
	r.items = append(r.items, m)
}

// BuscarPorID retorna un material existente por su ID.
func (r *MaterialRepo) BuscarPorID(id int) (Prestable, error) {
	for _, item := range r.items {
		if item.GetID() == id {
			return item, nil
		}
	}
	return nil, errors.New("material no encontrado")
}

// ListarTodos retorna la lista de todos los materiales almacenados.
func (r *MaterialRepo) ListarTodos() []Prestable {
	return r.items
}

// Actualizar reemplaza la instancia de un material guardado por el mismo ID.
func (r *MaterialRepo) Actualizar(m Prestable) error {
	for i, item := range r.items {
		if item.GetID() == m.GetID() {
			r.items[i] = m
			return nil
		}
	}
	return errors.New("material no encontrado para actualizar")
}

// Eliminar remueve un material por su identificador único.
func (r *MaterialRepo) Eliminar(id int) error {
	for i, item := range r.items {
		if item.GetID() == id {
			r.items = append(r.items[:i], r.items[i+1:]...)
			return nil
		}
	}
	return errors.New("material no encontrado para eliminar")
}

// ActualizarEstado modifica el indicador de disponibilidad de un material.
func (r *MaterialRepo) ActualizarEstado(id int, disp bool) error {
	item, err := r.BuscarPorID(id)
	if err != nil {
		return err
	}
	item.SetDisponible(disp)
	return nil
}
