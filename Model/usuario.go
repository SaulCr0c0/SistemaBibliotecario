package model

import (
	"fmt"
)

// Usuario representa la entidad base de un usuario en el sistema.
type Usuario struct {
	ID          int
	Username    string
	password    string
	Rol         string
	NivelAcceso int    // Aplicable si Rol == "Administrador"
	Turno       string // Aplicable si Rol == "Bibliotecario"

}

// NewUsuario crea y retorna una nueva instancia de Usuario.
func NewUsuario(id int, username, password, rol string) Usuario {
	return Usuario{
		ID:       id,
		Username: username,
		password: password,
		Rol:      rol,
	}
}

// ValidarPassword verifica si la contraseña ingresada coincide con la almacenada.
func (u *Usuario) ValidarPassword(p string) bool {
	return u.password == p
}

// SetPassword asigna una nueva contraseña.
func (u *Usuario) SetPassword(p string) {
	u.password = p
}

// GetPassword retorna la contraseña almacenada.
func (u *Usuario) GetPassword() string {
	return u.password
}

// ObtenerDetalle retorna una representación formateada de los atributos específicos del rol.
func (u *Usuario) ObtenerDetalle() string {
	switch u.Rol {
	case "Administrador":
		return fmt.Sprintf("Nivel Acceso: %d", u.NivelAcceso)
	case "Bibliotecario":
		return fmt.Sprintf("Turno: %s", u.Turno)
	default:
		return "-"
	}
}

// Administrador representa un usuario con privilegios administrativos.
type Administrador struct {
	Usuario
}

// NewAdministrador crea un Administrador con sus campos embebidos.
func NewAdministrador(id int, username, password, rol string, nivelAcceso int) Administrador {
	u := NewUsuario(id, username, password, rol)
	u.NivelAcceso = nivelAcceso
	return Administrador{
		Usuario: u,
	}
}

// Bibliotecario representa un usuario del personal de biblioteca.
type Bibliotecario struct {
	Usuario
}

// NewBibliotecario crea un Bibliotecario con sus campos embebidos.
func NewBibliotecario(id int, username, password, rol string, turno string) Bibliotecario {
	u := NewUsuario(id, username, password, rol)
	u.Turno = turno
	return Bibliotecario{
		Usuario: u,
	}
}
