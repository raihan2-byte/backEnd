package handler

import (
	barang "blog/Barang"
	"blog/artikel"
	"blog/auth"
	"blog/berita"
	"blog/db"
	endpointcount "blog/endpointCount"
	"blog/home"
	"blog/karyakmpf"
	"blog/merch"
	"blog/middleware"
	"blog/phototalk"
	"blog/shortvideo"
	"blog/user"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApp() {
	db, err := db.InitDb()
	if err != nil {
		log.Fatal("Eror Db Connection", err)
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin , Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
		AllowMethods:    []string{"POST, OPTIONS, GET, PUT, DELETE"},
	}))

	//auth
	secretKey := os.Getenv("SECRET_KEY")
	authService := auth.NewService()
	authService.SetSecretKey(secretKey)
	statisticsRepository := endpointcount.NewStatisticsRepository(db)
	// Inisialisasi service
	statisticsService := endpointcount.NewStatisticsService(statisticsRepository)
	// Inisialisasi handler
	statisticsHandler := NewStatisticsHandler(statisticsService)

	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := NewUserHandler(userService, authService)

	//berita
	beritaRepository := berita.NewRepository(db)
	beritaService := berita.NewService(beritaRepository)
	beritaHandler := NewBeritaHandler(beritaService, statisticsService)

	//barang
	barangRepository := barang.NewRepository(db)
	barangService := barang.NewService(barangRepository)
	barangHandler := NewBarangHandler(barangService, statisticsService)

	//photoTalk
	photoTalkRepository := phototalk.NewRepository(db)
	photoTalkService := phototalk.NewService(photoTalkRepository)
	photoTalkHandler := NewPhotoTalkHandler(photoTalkService, statisticsService)

	//KMPF
	karyakmpfRepository := karyakmpf.NewRepository(db)
	karyakmpfService := karyakmpf.NewService(karyakmpfRepository)
	karyakmpfHandler := NewKMPFHandler(karyakmpfService)

	//merch
	merchRepository := merch.NewRepository(db)
	merchService := merch.NewService(merchRepository)
	merchHandler := NewMerchHandler(merchService, statisticsService)

	//ShortVideo
	shortVideoRepository := shortvideo.NewRepository(db)
	shortVideoService := shortvideo.NewService(shortVideoRepository)
	shortVideoHandler := NewShortVideoHandler(shortVideoService, statisticsService)

	//Artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := NewArtikelHandler(artikelService, statisticsService)

	homeRepository := home.NewRepository(db)
	homeService := home.NewService(homeRepository)
	homeHandler := NewHomeHandler(homeService, statisticsService)

	//user
	api := router.Group("/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/", userHandler.GetAllUser)
	api.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), userHandler.DeletedUser)
	api.POST("/checkemail", userHandler.CheckEmailAvailabilty)

	//berita
	apiBerita := router.Group("/berita")
	apiBerita.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), beritaHandler.CreateBerita)
	apiBerita.GET("/", beritaHandler.GetAllBerita)
	apiBerita.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), beritaHandler.DeleteBerita)
	apiBerita.PUT("update/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), beritaHandler.UpdateBerita)
	apiBerita.GET("/:slug", beritaHandler.GetOneBerita)
	apiBerita.GET("/tags/:id", beritaHandler.GetByTags)
	apiBerita.GET("/karya", beritaHandler.GetByKarya)

	//barang
	apiBarang := router.Group("/barang")
	apiBarang.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), barangHandler.CreateBarang)
	apiBarang.GET("/", barangHandler.GetAllBarang)
	apiBarang.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), barangHandler.DeleteBarang)
	apiBarang.PUT("update/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), barangHandler.UpdateBarang)
	apiBarang.GET("/:id", barangHandler.GetOneBarang)
	apiBarang.GET("/category/:id", barangHandler.GetCategoryBarang)

	//photoTalk
	apiPhotoTalk := router.Group("/phototalk")
	apiPhotoTalk.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), photoTalkHandler.CreatePhotoTalk)
	apiPhotoTalk.GET("/", photoTalkHandler.GetAllPhotoTalk)
	apiPhotoTalk.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), photoTalkHandler.DeletePhotoTalk)
	apiPhotoTalk.GET("/:id", photoTalkHandler.GetOnePhotoTalk)

	//karyakmpf
	apikaryakmpf := router.Group("/karyakmpf")
	apikaryakmpf.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), karyakmpfHandler.CreateKMPF)
	apikaryakmpf.GET("/", karyakmpfHandler.GetAllKMPF)
	apikaryakmpf.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), karyakmpfHandler.DeleteKMPF)
	apikaryakmpf.GET("/:id", karyakmpfHandler.GetOneKMPF)

	//merch
	apimerch := router.Group("/merch")
	apimerch.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), merchHandler.CreateMerch)
	apimerch.GET("/", merchHandler.GetAllMerch)
	apimerch.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), merchHandler.DeleteMerch)
	apimerch.GET("/:id", merchHandler.GetOneMerch)

	//ShortVideo
	apiShortVideo := router.Group("/short-video")
	apiShortVideo.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), shortVideoHandler.CreateShortVideo)
	apiShortVideo.GET("/", shortVideoHandler.GetAllShortVideo)
	apiShortVideo.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), shortVideoHandler.DeleteShortVideo)
	apiShortVideo.GET("/:id", shortVideoHandler.GetOneShortVideo)

	//artikel
	apiArtikel := router.Group("/artikel")
	apiArtikel.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.CreateArtikel)
	apiArtikel.GET("/", artikelHandler.GetAllArtikel)
	apiArtikel.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), artikelHandler.DeleteArtikel)
	apiArtikel.GET("/:id", artikelHandler.GetOneArtikel)

	apiHome := router.Group("/tagLine")
	apiHome.POST("/", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), homeHandler.CreateTagHome)
	apiHome.GET("/", homeHandler.GetAllTagHome)
	apiHome.DELETE("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), homeHandler.DeleteHome)
	apiHome.PUT("/delete/:id", middleware.AuthMiddleware(authService, userService), middleware.AuthRole(authService, userService), homeHandler.UpdateTagHome)

	// apiHome.GET("/:id", homeHandler.GetAllTagHome)

	// statistics
	router.GET("/statistics", statisticsHandler.GetStatisticsHandler)
	router.GET("/statistics/:endpoint", statisticsHandler.GetTotalCountForEndpointHandler)

	router.Run(":8080")

}
