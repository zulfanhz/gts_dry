package router

import (
	"database/sql"
	"gts-dry/handler"
	"gts-dry/middleware"
	"gts-dry/repository"
	"gts-dry/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupRouter(db *sql.DB) *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders: "Content-Type, Authorization, Origin, Accept",
	}))

	restoRepo := repository.NewRestoRepository(db)
	userRepo := repository.NewUserRepository(db)
	barangRepo := repository.NewBarangRepository(db)
	rakRepo := repository.NewRakRepository(db)
	adjustmentRepo := repository.NewAdjustmentRepository(db)
	outgoingRepo := repository.NewOutgoingRepository(db)
	incomingRepo := repository.NewIncomingRepository(db)
	mutasiRepo := repository.NewMutasiRepository(db)

	restoService := service.NewRestoService(restoRepo)
	userService := service.NewUserService(userRepo)
	rakService := service.NewRakService(rakRepo)
	barangService := service.NewBarangService(barangRepo, rakRepo)
	adjustmentService := service.NewAdjustmentService(adjustmentRepo, rakRepo, barangRepo)
	outgoingService := service.NewOutgoingService(outgoingRepo, rakRepo, restoRepo, barangRepo)
	incomingService := service.NewIncomingService(incomingRepo, rakRepo, barangRepo)
	mutasiService := service.NewMutasiService(mutasiRepo, rakRepo, barangRepo)

	restoHandler := handler.NewRestoHandler(restoService)
	userHandler := handler.NewUserHandler(userService)
	barangHandler := handler.NewBarangHandler(barangService)
	rakHandler := handler.NewRakHandler(rakService, userService)
	adjustmentHandler := handler.NewAdjustmentHandler(adjustmentService)
	outgoingHandler := handler.NewOutgoingHandler(outgoingService)
	incominghandler := handler.NewIncomingHandler(incomingService)
	mutasiHandler := handler.NewMutasiHandler(mutasiService)

	api := app.Group("/gtsdry/api")

	// api.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "TEEESSSSSSS"})
	// })

	// fmt.Println(rakRepo.CekRakisAvailable("100051", "GCK2-R011A3", "2024-05-30"))

	// resto
	api.Get("/resto", restoHandler.GetRestoAll)
	api.Get("/resto/:kategori", restoHandler.GetRestoByKategori)
	api.Get("/resto/kode/:kode", restoHandler.GetRestoByKode)

	//user
	api.Post("/login", userHandler.LoginUser)
	api.Get("/user", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), userHandler.GetUserByToken)
	api.Post("/changepassword", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), userHandler.ChangePassword)

	//barang
	api.Get("/barang", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.GetBarangAll)
	api.Get("/barang/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.GetBarangByKode)
	api.Get("/barang/kategori/:kategori", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.GetBarangByKategori)

	//simpan barang dengan satuan
	api.Post("/barang/", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.AddBarang)
	//update barang
	api.Put("/barang/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.UpdateBarang)
	//hapus barang sekaligus satuan barang
	api.Delete("/barang/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.DeleteBarang)
	//tambah satuan barang
	api.Post("/barang/satuan", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.AddSatuanBarang)
	//ubah satuan flag
	api.Put("/barang/satuan/:kode/:satuan", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.UpdateSatuanBarang)
	//hapus satuan dan carikan flag baru yaitu yg level paling kecil
	api.Delete("barang/satuan/:kode/:satuan/:ishitung", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), barangHandler.DeleteSatuanBarang)

	//rak
	api.Get("/rak", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.GetRakAll)
	api.Get("/rak/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.GetRakByKodeRak)
	api.Get("/rak/type/:type", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.GetRakByType)
	api.Get("/rak/jenis/:jenis", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.GetRakByJenis)
	api.Get("/rak/available/:kode/:productCategory/:exp", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.CekRakListAvailableIncoming)
	api.Post("/rak", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.AddRak)
	api.Put("/rak/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.UpdateRak)
	api.Delete("/rak/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.DeleteRak)

	//lihat barang di rak mana saja
	api.Get("/rak/product/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.GetRakIsiByProductCode)
	api.Get("/rak/product/:productCode/:rakCode/:exp", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), rakHandler.GetRakIsiByProductCodedanRak)

	//adjusment
	api.Get("/adjustment", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), adjustmentHandler.GetAdjusmentAll)
	api.Get("/adjustment/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), adjustmentHandler.GetAdjusmentByKode)
	api.Post("/adjustment", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), adjustmentHandler.AddAdjustment)
	api.Put("/adjustment/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), adjustmentHandler.UpdateAdjustment)
	api.Delete("/adjustment/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), adjustmentHandler.DeleteAdjustment)

	//incoming
	api.Get("/incoming", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.GetIncomingAll)
	api.Get("/incoming/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.GetIncomingByKode)
	api.Get("/incoming/sj/:noSJ", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.GetIncomingBySJ)
	api.Get("/incoming/po/:po", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.GetIncomingByKode)
	api.Get("/incoming/po/:po/product/:codeProduct", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.GetIncomingByPOdanProduct)
	api.Post("/incoming", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.AddIncoming)
	api.Put("/incoming/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.UpdateIncoming)
	api.Delete("/incoming/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), incominghandler.DeleteIncoming)

	// mutasi
	api.Get("/mutasi", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), mutasiHandler.GetMutasiAll)
	api.Post("/mutasi", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), mutasiHandler.AddMutasi)

	//outgoing
	api.Get("/outgoing", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), outgoingHandler.GetOutgoingAll)
	api.Get("/outgoing/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), outgoingHandler.GetOutgoingByKode)
	api.Get("/outgoing/sj/:sj", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), outgoingHandler.GetOutgoingBySJ)
	api.Post("/outgoing", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), outgoingHandler.AddOutgoing)
	api.Put("/outgoing/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), outgoingHandler.UpdateOutgoing)
	api.Delete("/outgoing/:kode", middleware.JWTAuthMiddleware(os.Getenv("JWT_SECRET")), outgoingHandler.DeleteOutgoing)

	return app
}
