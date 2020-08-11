package routes

import (
	"docApp/controllers"
	"github.com/labstack/echo"
)

func InitRoute() *echo.Echo {
	e := echo.New()
	var ctrlUser controllers.UserController
	var ctrlAdmin controllers.AdminController
	var ctrlVariety controllers.VarietyController
	var ctrlCategory controllers.CategoryController
	var ctrlProduct controllers.ProductController
	var ctrlDoctor controllers.DoctorController
	var ctrlSchedule controllers.ScheduleController
	var ctrlService controllers.ServiceController
	var ctrlOrder controllers.OrderController
	var ctrlReservation controllers.ReservationController

	v1 := e.Group("/api")
	{
		vet := v1.Group("/veterinary")
		{
			user := vet.Group("/user")
			{
				user.GET("", ctrlUser.FindAll)
				user.GET("/:id", ctrlUser.FindById)
				user.POST("", ctrlUser.Create)
				user.PUT("", ctrlUser.Update)
				user.DELETE("/:id", ctrlUser.Delete)
			}

			admin := vet.Group("/admin")
			{
				admin.GET("", ctrlAdmin.FindAll)
				admin.GET("/:id", ctrlAdmin.FindById)
				admin.POST("", ctrlAdmin.Create)
				admin.PUT("", ctrlAdmin.Update)
				admin.DELETE("/:id", ctrlAdmin.Delete)
			}

			variety := vet.Group("/variety")
			{
				variety.GET("", ctrlVariety.FindAll)
				variety.GET("/:id", ctrlVariety.FindById)
				variety.POST("", ctrlVariety.Create)
				variety.PUT("", ctrlVariety.Update)
				variety.DELETE("/:id", ctrlVariety.Delete)
			}

			category := vet.Group("/category")
			{
				category.GET("", ctrlCategory.FindAll)
				category.GET("/:id", ctrlCategory.FindById)
				category.POST("", ctrlCategory.Create)
				category.PUT("", ctrlCategory.Update)
				category.DELETE("/:id", ctrlCategory.Delete)
			}

			product := vet.Group("/product")
			{
				product.GET("", ctrlProduct.FindAll)
				product.GET("/:id", ctrlProduct.FindById)
				product.POST("", ctrlProduct.Create)
				product.PUT("", ctrlProduct.Update)
				product.DELETE("/:id", ctrlProduct.Delete)
			}

			doctor := vet.Group("/doctor")
			{
				doctor.GET("", ctrlDoctor.FindAll)
				doctor.GET("/:id", ctrlDoctor.FindById)
				doctor.POST("", ctrlDoctor.Create)
				doctor.PUT("", ctrlDoctor.Update)
				doctor.DELETE("/:id", ctrlDoctor.Delete)
			}

			schedule := vet.Group("/schedule")
			{
				schedule.GET("", ctrlSchedule.FindAll)
				schedule.GET("/:id", ctrlSchedule.FindById)
				schedule.POST("", ctrlSchedule.Create)
				schedule.PUT("", ctrlSchedule.Update)
				schedule.DELETE("/:id", ctrlSchedule.Delete)
			}

			service := vet.Group("/service")
			{
				service.GET("", ctrlService.FindAll)
				service.GET("/id/:id", ctrlService.FindById)
				service.GET("/search/:name", ctrlService.FindByName)
				service.POST("", ctrlService.Create)
				service.PUT("", ctrlService.Update)
				service.DELETE("/:id", ctrlService.Delete)
			}

			order := vet.Group("/order")
			{
				order.GET("", ctrlOrder.FindAll)
				order.GET("/:id", ctrlOrder.FindById)
				order.POST("", ctrlOrder.Create)
				order.PUT("", ctrlOrder.Update)
				order.DELETE("/:id", ctrlOrder.Delete)
			}

			reservation := vet.Group("/reservation")
			{
				reservation.GET("", ctrlReservation.FindAll)
				reservation.GET("/:id", ctrlReservation.FindById)
				reservation.POST("", ctrlReservation.Create)
				reservation.PUT("", ctrlReservation.Update)
				reservation.DELETE("/:id", ctrlReservation.Delete)
			}
		}
	}
	return e
}
