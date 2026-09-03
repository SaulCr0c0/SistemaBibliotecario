package model

import (
	"errors"
)

// Usuario representa la entidad base de un usuario en el sistema.
type Usuario struct {
	ID       int
	Username string
	password string
	Rol      string
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

// Administrador representa un usuario con privilegios administrativos.
type Administrador struct {
	Usuario
	NivelAcceso int
}

// NewAdministrador crea un Administrador con sus campos embebidos.
func NewAdministrador(id int, username, password, rol string, nivelAcceso int) Administrador {
	return Administrador{
		Usuario:     NewUsuario(id, username, password, rol),
		NivelAcceso: nivelAcceso,
	}
}

// Bibliotecario representa un usuario del personal de biblioteca.
type Bibliotecario struct {
	Usuario
	Turno string
}

// NewBibliotecario crea un Bibliotecario con sus campos embebidos.
func NewBibliotecario(id int, username, password, rol string, turno string) Bibliotecario {
	return Bibliotecario{
		Usuario: NewUsuario(id, username, password, rol),
		Turno:   turno,
	}
}

// UserManager gestiona los usuarios en memoria y la sesión activa.
type UserManager struct {
	usuarios     []*Usuario
	sesionActiva *Usuario
}

// NewUserManager inicializa un UserManager con lista dinámica en memoria.
func NewUserManager() *UserManager {
	return &UserManager{
		usuarios:     make([]*Usuario, 0),
		sesionActiva: nil,
	}
}

// AgregarUsuario añade un usuario a la lista dinámica en memoria.
func (m *UserManager) AgregarUsuario(u *Usuario) {
	m.usuarios = append(m.usuarios, u)
}

// Autenticar busca el usuario por Username y valida la contraseña.
func (m *UserManager) Autenticar(u string, p string) (*Usuario, error) {
	for _, user := range m.usuarios {
		if user.Username == u {
			if user.ValidarPassword(p) {
				m.sesionActiva = user
				return user, nil
			}
			return nil, errors.New("Contraseña incorrecta")
		}
	}
	return nil, errors.New("Usuario no encontrado")
}

// CerrarSesion quita la referencia del usuario autenticado actualmente.
func (m *UserManager) CerrarSesion() {
	m.sesionActiva = nil
}

// GetSesion retorna el usuario de la sesión activa actual.
func (m *UserManager) GetSesion() *Usuario {
	return m.sesionActiva
}
