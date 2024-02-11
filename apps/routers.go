package apps

import (
	"healthcare-capt-america/middlewares"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) InitRouter(r *gin.Engine) {
	r.Use(middlewares.ErrorMiddleware())
	r.Use(middlewares.CorsMiddleware())
	auth := r.Group("/auth")
	auth.POST("register", h.AuthHandler.Register)
	auth.POST("verify", h.AuthHandler.Verify)
	auth.POST("login", h.AuthHandler.Login)
	auth.POST("forgot-password", h.ForgotPasswordHandler.GetToken)
	auth.POST("reset-password", h.ForgotPasswordHandler.ResetPassword)

	adminPharmacy := r.Group("/admin-pharmacy")
	adminPharmacy.POST("", h.AdminPharmacy.CreateAdminPharmacy)
	adminPharmacy.GET("", h.AdminPharmacy.GetAllAdmin)
	adminPharmacy.DELETE("/:id", h.AdminPharmacy.DeleteAdminPharmacy)
	adminPharmacy.GET("/:id", h.AdminPharmacy.GetDetailAdmin)
	adminPharmacy.PUT("/:id", h.AdminPharmacy.UpdateAdmin)
	adminPharmacy.GET("/pharmacies", middlewares.AuthMiddleware, h.PharmacyHandler.GetPharmaciesByLoginAdminPharmacy)
	adminPharmacy.GET("/products", middlewares.AuthMiddleware, h.PharmacyProductHandler.GetAllProductsByAdminPharmacy)
	adminPharmacy.PUT("/products/{id}", middlewares.AuthMiddleware, h.PharmacyProductHandler.EditProductInPharmacy)

	adminPharmacyOrder := adminPharmacy.Group("/orders")
	adminPharmacyOrder.Use(middlewares.AuthMiddleware)
	adminPharmacyOrder.GET("", h.OrderHandler.GetAllOrders)
	adminPharmacyOrder.GET("/:id", h.OrderHandler.GetOrderById)
	adminPharmacyOrder.PUT("/:id", h.OrderHandler.AdminUpdateOrder)

	doctors := r.Group("/doctors")
	doctors.GET("", h.DoctorHandler.GetAllDoctor)
	doctors.GET("/:id", h.DoctorHandler.GetDetailDoctor)
	doctors.PUT("/:id", h.DoctorHandler.UpdateStatusDoctor)

	doctor := r.Group("/doctor")
	doctor.GET("", middlewares.AuthMiddleware, h.DoctorHandler.GetCurrentDoctorDetail)
	doctor.PUT("", middlewares.AuthMiddleware, h.DoctorHandler.UpdateProfile)
	doctor.GET("telemedicines", middlewares.AuthMiddleware, h.TelemedicineHandler.GetAllTelemedicines)
	doctor.GET("telemedicines/:id", middlewares.AuthMiddleware, h.TelemedicineHandler.GetAllChatByRoomId)
	doctor.PUT("telemedicines/:id", middlewares.AuthMiddleware, h.TelemedicineHandler.Update)

	chat := r.Group("/chat")
	chat.GET("/:room_id", h.ChatHandler.WebSocketHandler)

	users := r.Group("/users")
	users.GET("", h.UserHandler.GetAllUsers)
	users.GET("/:id", h.UserHandler.GetUserDetail)

	user := r.Group("/user")
	user.Use(middlewares.AuthMiddleware)
	user.GET("", h.UserHandler.GetCurrentUserDetail)
	user.PUT("", h.UserHandler.UpdateProfile)
	user.GET("/products", h.ProductHandler.FindAllWithDist)

	userAddress := user.Group("/addresses")
	userAddress.Use(middlewares.AuthMiddleware)
	userAddress.POST("", h.UserHandler.AddAddress)
	userAddress.GET("", h.UserHandler.FindAllUserAddress)
	userAddress.POST("/default", h.UserHandler.SetDefaultAddress)
	userAddress.DELETE("", h.UserHandler.DeleteAddress)
	userAddress.PUT("", h.UserHandler.UpdateAddress)

	userOrder := user.Group("/orders")
	userOrder.Use(middlewares.AuthMiddleware)
	userOrder.GET("", h.OrderHandler.GetUserOrders)
	userOrder.GET("/:id", h.OrderHandler.GetUserDetailOrder)
	userOrder.PUT("/:id", h.OrderHandler.UpdateUserOrder)

	userTele := user.Group("/telemedicines")
	userTele.Use(middlewares.AuthMiddleware)
	userTele.GET("", h.TelemedicineHandler.GetAllTelemedicines)
	userTele.POST("", h.TelemedicineHandler.CreateRoom)

	productRouter := r.Group("/products")
	productRouter.GET("", h.ProductHandler.FindAllDrug)
	productRouter.GET("/:id", h.ProductHandler.FindByID)
	productRouter.POST("", h.ProductHandler.CreateDrug)
	productRouter.PUT("/:id", h.ProductHandler.EditDrug)

	productCategoryRouter := r.Group("/categories")
	productCategoryRouter.POST("", h.ProductCategoryHandler.CreateCategory)
	productCategoryRouter.GET("", h.ProductCategoryHandler.FindAllCategory)
	productCategoryRouter.GET("/:id", h.ProductCategoryHandler.FindByID)
	productCategoryRouter.PUT("/:id", h.ProductCategoryHandler.EditCategory)
	productCategoryRouter.DELETE("/:id", h.ProductCategoryHandler.DeleteCategory)

	pharmacyRouter := r.Group("pharmacies")
	pharmacyRouter.POST("", middlewares.AuthMiddleware, h.PharmacyHandler.AddPharmacy)
	pharmacyRouter.GET("", h.PharmacyHandler.GetPharmacies)
	pharmacyRouter.GET("/:id", h.PharmacyHandler.GetPharmacyById)
	operationalRouter := pharmacyRouter.Group("operationals")
	operationalRouter.POST("", h.OperationalHandler.AddOperationalDays)
	operationalRouter.DELETE("/:id", h.OperationalHandler.DeletePharmacyOperationalDay)
	pharmacyRouter.PUT("/:id", middlewares.AuthMiddleware, h.PharmacyHandler.UpdatePharmacy)
	pharmacyRouter.DELETE("/:id", middlewares.AuthMiddleware, h.PharmacyHandler.DeletePharmacy)

	pharmacyProductRouter := pharmacyRouter.Group("/:id/products")
	pharmacyProductRouter.GET("", h.PharmacyProductHandler.FindAllByPharmacyID)
	pharmacyProductRouter.GET("/:drug_id", h.PharmacyProductHandler.FindByID)
	pharmacyProductRouter.POST("", h.PharmacyProductHandler.CreatePharmacyDrug)
	pharmacyProductRouter.PUT("/:drug_id", h.PharmacyProductHandler.EditPharmacyDrug)
	pharmacyProductRouter.DELETE("/:drug_id", h.PharmacyProductHandler.DeletePharmacyDrug)

	cartRouter := r.Group("carts")
	cartRouter.Use(middlewares.AuthMiddleware)
	cartRouter.POST("", h.CartHandler.AddCart)
	cartRouter.GET("", h.CartHandler.GetCartsByLoginUser)
	cartRouter.DELETE("/:id", h.CartHandler.DeleteCartById)
	cartRouter.DELETE("", h.CartHandler.DeleteAllDrugs)

	journalRouter := r.Group("journals")
	journalRouter.GET("", h.JournalHandler.FindAllJournal)

	stockMutationRouter := r.Group("stock-mutations")
	stockMutationRouter.Use(middlewares.AuthMiddleware)
	stockMutationRouter.GET("", h.StockMutation.FindAll)
	stockMutationRouter.GET("/:id/", h.StockMutation.FindById)
	stockMutationRouter.POST("", h.StockMutation.CreateRequestStockMutation)
	stockMutationRouter.DELETE("/:id/", h.StockMutation.Delete)
	stockMutationRouter.PUT("/:id/", h.StockMutation.Update)

	orderRouter := r.Group("orders")
	orderRouter.Use(middlewares.AuthMiddleware)
	orderRouter.GET("", h.OrderHandler.GetAllOrders)
	orderRouter.DELETE("/:id", h.OrderHandler.DeleteOrderById)
	orderRouter.GET("/:id", h.OrderHandler.GetOrderById)
	orderRouter.POST("", middlewares.AuthMiddleware, h.OrderHandler.CreateOrder)

	paymentRouter := r.Group("payments")
	paymentRouter.POST("", middlewares.AuthMiddleware, h.PaymentHandler.UploadPayment)
	paymentRouter.PUT("/:id", h.PaymentHandler.UpdatePaymentStatus)

	provinceRouter := r.Group("provinces")
	provinceRouter.GET("", h.ProvinceHandler.GetAllProvinces)
	provinceRouter.GET("/:id", h.ProvinceHandler.GetProvinceById)

	manufactureRouter := r.Group("manufactures")
	manufactureRouter.GET("", h.ManufactureHandler.GetAllManufactures)

	formRouter := r.Group("drug-forms")
	formRouter.GET("", h.FormHandler.GetAllForms)

	categoryRouter := r.Group("product-categories")
	categoryRouter.GET("", h.ProductCategoryHandler.FindAllCategories)

	checkoutRouter := r.Group("/checkout")
	checkoutRouter.POST("", middlewares.AuthMiddleware, h.CheckoutHandler.Checkout)
	shipment := r.Group("shipment")
	shipment.GET("", h.ShipmentHandler.FindAll)
	shipment.POST("/fee", h.ShipmentHandler.CalculateDistance)

	specializationRouter := r.Group("/specializations")
	specializationRouter.GET("", h.SpecializationHandler.GetAllSpecializations)

	telemedicinesRouter := r.Group("/telemedicines")
	//telemedicinesRouter.Use(middlewares.AuthMiddleware)
	telemedicinesRouter.GET("/:id", h.ChatHandler.GetChatHistoriesRoom)
}
