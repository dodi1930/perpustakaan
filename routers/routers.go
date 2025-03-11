package routers

import (
	"perpustakaan-x-cgpt/controllers"
	"perpustakaan-x-cgpt/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter mengatur semua rute API
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rute tanpa autentikasi
	r.POST("/login/admin", controllers.AdminLogin)
	r.POST("/login", controllers.UserLogin)

	// Grup API dengan autentikasi Admin
	admin := r.Group("/")
	admin.Use(middlewares.AuthMiddleware("admin")) // Middleware hanya untuk admin
	{
		admin.GET("/admin/users", controllers.GetUsers)
		admin.POST("/admin/add-user", controllers.AddUser)
		admin.PUT("/admin/users/:id", controllers.EditUser)
		admin.DELETE("/admin/users/:id", controllers.DeleteUser)

		admin.POST("/admin/books", controllers.AddBookHandler)
		admin.PUT("/admin/books/:id", controllers.UpdateBook)
		admin.DELETE("/admin/books/:id", controllers.DeleteBook)
		admin.GET("/admin/books/search", controllers.SearchBooks)

		admin.GET("/admin/users/:id/fines", controllers.GetUserFineHandler)
		admin.GET("/admin/statistic", controllers.GetLoanStatistics)
	}

	// Grup API untuk User biasa
	siswa := r.Group("/")
	siswa.Use(middlewares.AuthMiddleware("siswa")) // Middleware hanya untuk user
	{
		siswa.POST("/siswa/loans", controllers.GetLoanNotifications)
		siswa.POST("/siswa/searchbook", controllers.SearchBooks)

		//user.PUT("/loans/:id/return", controllers.ReturnBook)
	}
	// Grup API untuk pustakawan
	pustakawan := r.Group("/")
	pustakawan.Use(middlewares.AuthMiddleware("pustakawan")) // Middleware hanya untuk user
	{
		pustakawan.GET("/pustakawan/:id/fines", controllers.GetUserFineHandler)
		pustakawan.GET("/pustakawan/books/search", controllers.SearchBooks)
		pustakawan.POST("/pustakawan/books", controllers.AddBookHandler)
		pustakawan.PUT("/pustakawan/books/:id", controllers.UpdateBook)
		pustakawan.DELETE("/pustakawan/books/:id", controllers.DeleteBook)
	}
	// SetupDashboardRoutes mengatur rute untuk dashboard
	//func SetupDashboardRoutes(r *gin.Engine) {
	r.GET("/adm", controllers.DashboardHandler)

	return r

}
