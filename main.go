package main

import (
	barang "blog/Barang"
	"blog/artikel"
	"blog/auth"
	"blog/berita"
	endpointcount "blog/endpointCount"
	"blog/handler"
	"blog/helper"
	"blog/home"
	"blog/karyakmpf"
	"blog/merch"
	"blog/phototalk"
	"blog/shortvideo"
	"blog/user"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {

	if _, exists := os.LookupEnv("RAILWAY_ENVIRONMENT"); exists == false {
		if err := godotenv.Load(); err != nil {
			log.Fatal("error loading .env file:", err)
		}
	}

	dbUsername := os.Getenv("MYSQLUSER")
	dbPassword := os.Getenv("MYSQLPASSWORD")
	dbHost := os.Getenv("MYSQLHOST")
	dbPort := os.Getenv("MYSQLPORT")
	dbName := os.Getenv("MYSQLDATABASE")
	secretKey := os.Getenv("SECRET_KEY")

	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}

	err = db.AutoMigrate(&home.TagLineHome{} ,&user.User{}, &endpointcount.Statistics{}, &berita.Berita{}, &barang.Barang{}, &phototalk.PhotoTalk{}, &karyakmpf.KMPF{}, &merch.Merch{}, &barang.Category{}, &berita.TagsBerita{}, &shortvideo.ShortVideo{}, &berita.KaryaBerita{}, &artikel.Artikel{}, &berita.BeritaImage{})
	if err != nil {
		log.Fatal("eror migration")
	}
	//auth
	authService := auth.NewService()
	authService.SetSecretKey(secretKey)
	statisticsRepository := endpointcount.NewStatisticsRepository(db)
	// Inisialisasi service
	statisticsService := endpointcount.NewStatisticsService(statisticsRepository)
	// Inisialisasi handler
	statisticsHandler := handler.NewStatisticsHandler(statisticsService)

	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	//berita
	beritaRepository := berita.NewRepository(db)
	beritaService := berita.NewService(beritaRepository)
	beritaHandler := handler.NewBeritaHandler(beritaService, statisticsService)

	//barang
	barangRepository := barang.NewRepository(db)
	barangService := barang.NewService(barangRepository)
	barangHandler := handler.NewBarangHandler(barangService, statisticsService)

	//photoTalk
	photoTalkRepository := phototalk.NewRepository(db)
	photoTalkService := phototalk.NewService(photoTalkRepository)
	photoTalkHandler := handler.NewPhotoTalkHandler(photoTalkService, statisticsService)

	//KMPF
	karyakmpfRepository := karyakmpf.NewRepository(db)
	karyakmpfService := karyakmpf.NewService(karyakmpfRepository)
	karyakmpfHandler := handler.NewKMPFHandler(karyakmpfService)

	//merch
	merchRepository := merch.NewRepository(db)
	merchService := merch.NewService(merchRepository)
	merchHandler := handler.NewMerchHandler(merchService, statisticsService)

	//ShortVideo 
	shortVideoRepository := shortvideo.NewRepository(db)
	shortVideoService := shortvideo.NewService(shortVideoRepository)
	shortVideoHandler := handler.NewShortVideoHandler(shortVideoService, statisticsService)

	//Artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := handler.NewArtikelHandler(artikelService, statisticsService)

	homeRepository := home.NewRepository(db)
	homeService := home.NewService(homeRepository)
	homeHandler := handler.NewHomeHandler(homeService, statisticsService)

  router := gin.Default()
  router.Use(cors.New(cors.Config{
	AllowAllOrigins: true,
	AllowHeaders: []string{"Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Access-Control-Allow-Origin , Origin , Accept , X-Requested-With , Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, Authorization"},
	AllowMethods: []string{"POST, OPTIONS, GET, PUT, DELETE"},
  }))

	//user
	api := router.Group("/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/", userHandler.GetAllUser)
	api.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), userHandler.DeletedUser)
	api.POST("/checkemail", userHandler.CheckEmailAvailabilty)

	//berita
	apiBerita := router.Group("/berita")
	apiBerita.POST("/", authMiddleware(authService, userService), authRole(authService, userService),  beritaHandler.CreateBerita)
	apiBerita.GET("/", beritaHandler.GetAllBerita)
	apiBerita.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), beritaHandler.DeleteBerita)
	apiBerita.PUT("update/:id", authMiddleware(authService, userService), authRole(authService, userService), beritaHandler.UpdateBerita)
	apiBerita.GET("/:id", beritaHandler.GetOneBerita)
	apiBerita.GET("/tags/:id", beritaHandler.GetByTags)
	apiBerita.GET("/karya", beritaHandler.GetByKarya)

	
	//barang
	apiBarang := router.Group("/barang")
	apiBarang.POST("/", authMiddleware(authService, userService), authRole(authService, userService), barangHandler.CreateBarang)
	apiBarang.GET("/", barangHandler.GetAllBarang)
	apiBarang.DELETE("/delete/:id", authMiddleware(authService, userService) , authRole(authService, userService), barangHandler.DeleteBarang)
	apiBarang.PUT("update/:id", authMiddleware(authService, userService), authRole(authService, userService), barangHandler.UpdateBarang)
	apiBarang.GET("/:id", barangHandler.GetOneBarang)
	apiBarang.GET("/category/:id",barangHandler.GetCategoryBarang)


	//photoTalk
	apiPhotoTalk := router.Group("/phototalk")
	apiPhotoTalk.POST("/", authMiddleware(authService, userService), authRole(authService, userService), photoTalkHandler.CreatePhotoTalk)
	apiPhotoTalk.GET("/", photoTalkHandler.GetAllPhotoTalk)
	apiPhotoTalk.DELETE("/delete/:id", authMiddleware(authService, userService) , authRole(authService, userService), photoTalkHandler.DeletePhotoTalk)
	apiPhotoTalk.GET("/:id", photoTalkHandler.GetOnePhotoTalk)

	//karyakmpf
	apikaryakmpf := router.Group("/karyakmpf")
	apikaryakmpf.POST("/", authMiddleware(authService, userService), authRole(authService, userService), karyakmpfHandler.CreateKMPF)
	apikaryakmpf.GET("/", karyakmpfHandler.GetAllKMPF)
	apikaryakmpf.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), karyakmpfHandler.DeleteKMPF)
	apikaryakmpf.GET("/:id", karyakmpfHandler.GetOneKMPF)

	//merch
	apimerch := router.Group("/merch")
	apimerch.POST("/", authMiddleware(authService, userService), authRole(authService, userService), merchHandler.CreateMerch)
	apimerch.GET("/", merchHandler.GetAllMerch)
	apimerch.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), merchHandler.DeleteMerch)
	apimerch.GET("/:id",  merchHandler.GetOneMerch)

	//ShortVideo
	apiShortVideo := router.Group("/short-video")
	apiShortVideo.POST("/", authMiddleware(authService, userService), authRole(authService, userService), shortVideoHandler.CreateShortVideo)
	apiShortVideo.GET("/", shortVideoHandler.GetAllShortVideo)
	apiShortVideo.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), shortVideoHandler.DeleteShortVideo)
	apiShortVideo.GET("/:id", shortVideoHandler.GetOneShortVideo)

	//artikel
	apiArtikel := router.Group("/artikel")
	apiArtikel.POST("/", authMiddleware(authService, userService), authRole(authService, userService), artikelHandler.CreateArtikel)
	apiArtikel.GET("/", artikelHandler.GetAllArtikel)
	apiArtikel.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), artikelHandler.DeleteArtikel)
	apiArtikel.GET("/:id", artikelHandler.GetOneArtikel)

	apiHome := router.Group("/tagLine")
	apiHome.POST("/", authMiddleware(authService, userService), authRole(authService, userService), homeHandler.CreateTagHome)
	apiHome.GET("/", homeHandler.GetAllTagHome)
	apiHome.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), homeHandler.DeleteHome)
	apiHome.PUT("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), homeHandler.UpdateTagHome)

	// apiHome.GET("/:id", homeHandler.GetAllTagHome)


	// statistics
	router.GET("/statistics", statisticsHandler.GetStatisticsHandler)
	router.GET("/total-unique-user-agents", statisticsHandler.GetTotalUniqueUserAgentsHandler)


	router.Run(":8080")

}


func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)
 		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			//nah ini kalau emang ada dua key nya dan sesuai, maka tokenString tadi masuk ke arrtoken index ke1
			tokenString = arrToken[1]
		}
		token, err := authService.ValidasiToken(tokenString)
		// fmt.Println(token, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println(claim, ok)
		if !ok || !token.Valid {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByid(userID)
		// fmt.Println(user, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}




func authRole(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// fmt.Println(authHeader)
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := ""
		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) == 2 {
			//nah ini kalau emang ada dua key nya dan sesuai, maka tokenString tadi masuk ke arrtoken index ke1
			tokenString = arrToken[1]
		}
		token, err := authService.ValidasiToken(tokenString)
		// fmt.Println(token, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		// fmt.Println(claim, ok)
		if !ok || !token.Valid {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		userID := int(claim["user_id"].(float64))

		if int(claim["role"].(float64)) != 1 {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user, err := userService.GetUserByid(userID)
		// fmt.Println(user, err)
		if err != nil {
			response := helper.APIresponse(http.StatusUnauthorized, nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		c.Set("currentUser", user)
	}
}
