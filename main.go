package main

import (
	barang "blog/Barang"
	"blog/artikel"
	"blog/auth"
	"blog/berita"
	"blog/handler"
	"blog/helper"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	// dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")

	dsn := dbUsername + ":" + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection Error")
	}
	// 1.User (Login, Register dll)
	// 2.Admin CRUD berita
	// 3.Admin CRUD barang
	// 4.Admin CRUD PhotoTalk
	// 5.Admin CRUD Merch
	// 6.Admin CRUD Karya kmpf unj

	err = db.AutoMigrate(&user.User{}, &berita.Berita{}, &barang.Barang{}, &phototalk.PhotoTalk{}, &karyakmpf.KMPF{}, &merch.Merch{}, &barang.Category{}, &berita.TagsBerita{}, &shortvideo.ShortVideo{}, &berita.KaryaBerita{}, &artikel.Artikel{})
	if err != nil {
		log.Fatal("eror migration")
	}
	//auth
	authService := auth.NewService()
	authService.SetSecretKey(secretKey)

	//user
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	//berita
	beritaRepository := berita.NewRepository(db)
	beritaService := berita.NewService(beritaRepository)
	beritaHandler := handler.NewBeritaHandler(beritaService)

	//barang
	barangRepository := barang.NewRepository(db)
	barangService := barang.NewService(barangRepository)
	barangHandler := handler.NewBarangHandler(barangService)

	//photoTalk
	photoTalkRepository := phototalk.NewRepository(db)
	photoTalkService := phototalk.NewService(photoTalkRepository)
	photoTalkHandler := handler.NewPhotoTalkHandler(photoTalkService)

	//KMPF
	karyakmpfRepository := karyakmpf.NewRepository(db)
	karyakmpfService := karyakmpf.NewService(karyakmpfRepository)
	karyakmpfHandler := handler.NewKMPFHandler(karyakmpfService)

	//merch
	merchRepository := merch.NewRepository(db)
	merchService := merch.NewService(merchRepository)
	merchHandler := handler.NewMerchHandler(merchService)

	//ShortVideo 
	shortVideoRepository := shortvideo.NewRepository(db)
	shortVideoService := shortvideo.NewService(shortVideoRepository)
	shortVideoHandler := handler.NewShortVideoHandler(shortVideoService)

	//Artikel
	artikelRepository := artikel.NewRepository(db)
	artikelService := artikel.NewService(artikelRepository)
	artikelHandler := handler.NewArtikelHandler(artikelService)

	router := gin.Default()
	router.Use(cors.Default())


	//user
	api := router.Group("/users")
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.GET("/get", authMiddleware(authService, userService), authRole(authService, userService), userHandler.GetAllUser)
	api.DELETE("/delete/:id", authRole(authService, userService), userHandler.DeletedUser)
	api.POST("/checkemail", userHandler.CheckEmailAvailabilty)

	//berita
	apiBerita := router.Group("/berita")
	apiBerita.POST("/", authMiddleware(authService, userService), authRole(authService, userService),  beritaHandler.CreateBerita)
	apiBerita.GET("/", beritaHandler.GetAllBerita)
	apiBerita.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), beritaHandler.DeleteBerita)
	apiBerita.GET("/:id", authMiddleware(authService, userService), authRole(authService, userService), beritaHandler.GetOneBerita)
	apiBerita.GET("/tags/:id", beritaHandler.GetByTags)
	apiBerita.GET("/karya/:id", beritaHandler.GetByKarya)

	
	//barang
	apiBarang := router.Group("/barang")
	apiBarang.POST("/", authMiddleware(authService, userService), authRole(authService, userService), barangHandler.CreateBarang)
	apiBarang.GET("/", authMiddleware(authService, userService), authRole(authService, userService), barangHandler.GetAllBarang)
	apiBarang.DELETE("/delete/:id", authMiddleware(authService, userService) , authRole(authService, userService), barangHandler.DeleteBarang)
	apiBarang.GET("/:id", authMiddleware(authService, userService) , authRole(authService, userService), barangHandler.GetOneBarang)
	apiBarang.GET("/category/:id",barangHandler.GetCategoryBarang)

	//photoTalk
	apiPhotoTalk := router.Group("/phototalk")
	apiPhotoTalk.POST("/", authMiddleware(authService, userService), authRole(authService, userService), photoTalkHandler.CreatePhotoTalk)
	apiPhotoTalk.GET("/", authMiddleware(authService, userService), authRole(authService, userService), photoTalkHandler.GetAllPhotoTalk)
	apiPhotoTalk.GET("/:id", authMiddleware(authService, userService) , authRole(authService, userService), photoTalkHandler.GetOnePhotoTalk)

	//karyakmpf
	apikaryakmpf := router.Group("/karyakmpf")
	apikaryakmpf.POST("/", authMiddleware(authService, userService), authRole(authService, userService), karyakmpfHandler.CreateKMPF)
	apikaryakmpf.GET("/", authMiddleware(authService, userService), authRole(authService, userService), karyakmpfHandler.GetAllKMPF)
	apikaryakmpf.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), karyakmpfHandler.DeleteKMPF)
	apikaryakmpf.GET("/:id", authMiddleware(authService, userService), authRole(authService, userService), karyakmpfHandler.GetOneKMPF)

	//merch
	apimerch := router.Group("/merch")
	apimerch.POST("/", authMiddleware(authService, userService), authRole(authService, userService), merchHandler.CreateMerch)
	apimerch.GET("/", authMiddleware(authService, userService), authRole(authService, userService), merchHandler.GetAllMerch)
	apimerch.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), merchHandler.DeleteMerch)
	apimerch.GET("/:id", authMiddleware(authService, userService), authRole(authService, userService), merchHandler.GetOneMerch)

	//ShortVideo
	apiShortVideo := router.Group("/short-video")
	apiShortVideo.POST("/", authMiddleware(authService, userService), authRole(authService, userService), shortVideoHandler.CreateShortVideo)
	apiShortVideo.GET("/", authMiddleware(authService, userService), authRole(authService, userService), shortVideoHandler.GetAllShortVideo)
	apiShortVideo.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), shortVideoHandler.DeleteShortVideo)
	apiShortVideo.GET("/:id", authMiddleware(authService, userService), authRole(authService, userService), shortVideoHandler.GetOneShortVideo)

	//artikel
	apiArtikel := router.Group("/artikel")
	apiArtikel.POST("/", authMiddleware(authService, userService), authRole(authService, userService), artikelHandler.CreateArtikel)
	apiArtikel.GET("/", authMiddleware(authService, userService), authRole(authService, userService), artikelHandler.GetAllArtikel)
	apiArtikel.DELETE("/delete/:id", authMiddleware(authService, userService), authRole(authService, userService), artikelHandler.DeleteArtikel)
	apiArtikel.GET("/:id", authMiddleware(authService, userService), authRole(authService, userService), artikelHandler.GetOneArtikel)


	router.Run(":8080")

}

// func CustomHeaderAPI(c *gin.Context) {
//     // Add CORS headers
//     c.Header("Access-Control-Allow-Origin", "http://example.com")
//     c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")

//     // Prepare response
   
// }

// func CORSMiddleware() gin.HandlerFunc {
//     return func(c *gin.Context) {
//         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
//         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
//         c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
//         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

//         if c.Request.Method == "OPTIONS" {
//             c.AbortWithStatus(204)
//             return
//         }

//         c.Next()
//     }
// }

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
