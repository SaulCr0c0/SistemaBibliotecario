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

	// Cargar usuarios simulados en listas dinámicas
	admin := model.NewAdministrador(1, "admin", "admin123", "Administrador", 1)
	biblio := model.NewBibliotecario(2, "biblio", "biblio123", "Bibliotecario", "Mañana")

	userMgr.AgregarUsuario(&admin.Usuario)
	userMgr.AgregarUsuario(&biblio.Usuario)

	// Cargar inventario de materiales simulados
	materialRepo := model.NewMaterialRepo()
	libroSample := model.NewLibro(1, "Cien Años de Soledad", 1967, "978-0307474728", "Editorial Sudamericana")
	tesisSample := model.NewTesis(2, "Optimización de Consultas Graph", 2024, "Ingeniería Informática", "Dr. Roberto Carlos")
	revistaSample := model.NewRevista(3, "ACM Computing Surveys", 2025, "0360-0300", 56)

	materialRepo.Agregar(libroSample)
	materialRepo.Agregar(tesisSample)
	materialRepo.Agregar(revistaSample)

	// 2. Inicializar Capa Vista y Enrutador
	router := view.NewAppRouter()
	loginView := view.NewLoginView()
	mainMenuView := view.NewMainMenuView()
	userTableView := view.NewUserTableView()
	userFormView := view.NewUserFormView()
	materialTableView := view.NewMaterialTableView()
	materialFormView := view.NewMaterialFormView()

	// Registrar pantallas en el AppRouter
	router.AgregarPantalla("login", loginView.GetPrimitive())
	router.AgregarPantalla("main_menu", mainMenuView.GetPrimitive())
	router.AgregarPantalla("user_table", userTableView.GetPrimitive())
	router.AgregarPantalla("user_form", userFormView.GetPrimitive())
	router.AgregarPantalla("material_table", materialTableView.GetPrimitive())
	router.AgregarPantalla("material_form", materialFormView.GetPrimitive())

	// 3. Inicializar Capa Controlador
	authCtrl := controller.NewAuthController(userMgr)
	userCtrl := controller.NewUserController(userMgr, userTableView, userFormView, router)
	materialCtrl := controller.NewMaterialController(materialRepo, materialTableView, materialFormView, router)

	libCtrl := controller.NewLibraryController(router, authCtrl, userCtrl, loginView, mainMenuView)
	libCtrl.SetMaterialController(materialCtrl)

	// 4. Ejecutar aplicación tview iniciando en la pantalla de login
	if err := router.Run("login"); err != nil {
		log.Fatalf("Error al ejecutar la aplicación: %v", err)
	}
}
