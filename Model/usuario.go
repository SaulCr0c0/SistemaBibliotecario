package model

import (
	"errors"
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
	if u.Rol == "Administrador" {
		return fmt.Sprintf("Nivel Acceso: %d", u.NivelAcceso)
	} else if u.Rol == "Bibliotecario" {
		return fmt.Sprintf("Turno: %s", u.Turno)
	}
	return "-"
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

// UserManager gestiona los usuarios en memoria y la sesión activa.
type UserManager struct {
	usuarios     []*Usuario
	sesionActiva *Usuario
	siguienteID  int
}

// NewUserManager inicializa un UserManager con lista dinámica en memoria.
func NewUserManager() *UserManager {
	return &UserManager{
		usuarios:     make([]*Usuario, 0),
		sesionActiva: nil,
		siguienteID:  1,
	}
}

// AgregarUsuario añade un usuario a la lista dinámica en memoria.
func (m *UserManager) AgregarUsuario(u *Usuario) {
	if u.ID >= m.siguienteID {
		m.siguienteID = u.ID + 1
	}
	m.usuarios = append(m.usuarios, u)
}

// ListarUsuarios retorna la lista completa de usuarios registrados.
func (m *UserManager) ListarUsuarios() []*Usuario {
	return m.usuarios
}

// BuscarPorID retorna el usuario correspondiente al ID especificado.
func (m *UserManager) BuscarPorID(id int) (*Usuario, error) {
	for _, u := range m.usuarios {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("Usuario no encontrado")
}

// CrearUsuario registra un nuevo usuario asignándole un ID autogenerado.
func (m *UserManager) CrearUsuario(username, password, rol string, nivelAcceso int, turno string) (*Usuario, error) {
	if username == "" || password == "" {
		return nil, errors.New("El nombre de usuario y la contraseña no pueden estar vacíos")
	}

	for _, u := range m.usuarios {
		if u.Username == username {
			return nil, errors.New("El nombre de usuario ya se encuentra registrado")
		}
	}

	nuevoUsuario := &Usuario{
		ID:          m.siguienteID,
		Username:    username,
		password:    password,
		Rol:         rol,
		NivelAcceso: nivelAcceso,
		Turno:       turno,
	}
	m.siguienteID++
	m.usuarios = append(m.usuarios, nuevoUsuario)

	return nuevoUsuario, nil
}

// ActualizarUsuario modifica los datos de un usuario existente.
func (m *UserManager) ActualizarUsuario(id int, username, password, rol string, nivelAcceso int, turno string) error {
	u, err := m.BuscarPorID(id)
	if err != nil {
		return err
	}

	for _, existing := range m.usuarios {
		if existing.ID != id && existing.Username == username {
			return errors.New("El nombre de usuario ya pertenece a otra cuenta")
		}
	}

	u.Username = username
	if password != "" {
		u.password = password
	}
	u.Rol = rol
	u.NivelAcceso = nivelAcceso
	u.Turno = turno

	return nil
}

// EliminarUsuario remueve un usuario de la lista por su ID.
func (m *UserManager) EliminarUsuario(id int) error {
	for i, u := range m.usuarios {
		if u.ID == id {
			if m.sesionActiva != nil && m.sesionActiva.ID == id {
				return errors.New("No es posible eliminar al usuario que tiene la sesión activa actualmente")
			}
			m.usuarios = append(m.usuarios[:i], m.usuarios[i+1:]...)
			return nil
		}
	}
	return errors.New("Usuario no encontrado")
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
