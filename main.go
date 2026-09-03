package main

import (
	"log"

	"biblioteca/Controller"
	"biblioteca/Model"
	"biblioteca/View"
)

func main() {
	// 1. Inicializar Capa Modelo
	userMgr := model.NewUserManager()

	// Cargar datos simulados en listas dinámicas en tiempo de ejecución
	admin := model.NewAdministrador(1, "admin", "admin123", "Administrador", 1)
	biblio := model.NewBibliotecario(2, "biblio", "biblio123", "Bibliotecario", "Mañana")

	userMgr.AgregarUsuario(&admin.Usuario)
	userMgr.AgregarUsuario(&biblio.Usuario)

	// 2. Inicializar Capa Vista y Enrutador
	router := view.NewAppRouter()
	loginView := view.NewLoginView()
	mainMenuView := view.NewMainMenuView()
	userTableView := view.NewUserTableView()
	userFormView := view.NewUserFormView()

	// Registrar pantallas en el AppRouter
	router.AgregarPantalla("login", loginView.GetPrimitive())
	router.AgregarPantalla("main_menu", mainMenuView.GetPrimitive())
	router.AgregarPantalla("user_table", userTableView.GetPrimitive())
	router.AgregarPantalla("user_form", userFormView.GetPrimitive())

	// 3. Inicializar Capa Controlador
	authCtrl := controller.NewAuthController(userMgr)
	userCtrl := controller.NewUserController(userMgr, userTableView, userFormView, router)
	_ = controller.NewLibraryController(router, authCtrl, userCtrl, loginView, mainMenuView)

	// 4. Ejecutar aplicación tview iniciando en la pantalla de login
	if err := router.Run("login"); err != nil {
		log.Fatalf("Error al ejecutar la aplicación: %v", err)
	}
}
