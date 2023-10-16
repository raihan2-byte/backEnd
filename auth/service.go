package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID int, role int) (string, error)
	ValidasiToken(token string) (*jwt.Token, error)
	// Refresh(refreshPayload string) (string, error)
	// PhotoAuthorization() gin.HandlerFunc
	// CommentAuthorization() gin.HandlerFunc
	// SosmedAuthorization() gin.HandlerFunc
}

type jwtService struct {
}

var SecretKey []byte

// pengertiannya kaya service di sebelah
func NewService() *jwtService {
	return &jwtService{}

}

func (s *jwtService) SetSecretKey(key string) {
	SecretKey = []byte(key)
}

func (s *jwtService) GenerateToken(userID int, role int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["role"] = role
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SecretKey)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidasiToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SecretKey), nil
	})
	if err != nil {
		return token, err
	}

	return token, nil
}

// func (u *jwtService) Refresh(refreshPayload string) (string, error) {
// 	// err := helpers.ValidateStruct(refreshPayload)

// 	// if err != nil {
// 	// 	return nil, err
// 	// }
	
// 	user := user.User{}


// 	_,errToken := user.ParseToken(refreshPayload.Token)

// 	if errToken != nil {
// 		if errToken.Error() == "token expired" {
// 			err := u.authRepo.RefreshTokenExists(refreshPayload.RefreshToken)
// 			if err != nil {
// 				return nil, err
// 			}
// 				_, errRefreshToken := user.ParseToken(refreshPayload.RefreshToken)
// 				if errRefreshToken != nil {
// 					return nil, errs.NewNotAuthenticated(err.Error())
// 				}

			
// 			errToken = user.ParseTokenUnverified(refreshPayload.Token)

// 			if errToken != nil {
			
// 				return nil, errs.NewInternalServerErrorr("something went wrong")
// 			}

// 			userData, err := u.userRepo.GetUserByIdAndEmail(user)

// 			if err != nil {
// 				return nil, errs.NewNotAuthenticated("invalid token")
// 			}

// 			newToken := &dto.RefreshTokenResponse{
// 				Token: userData.GenerateToken(),
// 			}

// 			return newToken, nil
// 		}

// 		return nil, errs.NewNotAuthenticated("invalid token")
// 	}


// 	return nil, errs.NewNotAuthenticated("token is still valid")
// }

// func (s *jwtService) PhotoAuthorization() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var getParam campaign.GetPhotoDetailInput
// 		err := ctx.ShouldBindUri(&getParam)
// 		if err != nil {
// 			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
// 			ctx.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}

// 		currentUser := ctx.MustGet("currentUser").(user.User)
// 		// input.User.ID = currentUser.ID
// 		// user := ctx.MustGet("userData").(entity.User)

// 		// movieId, err := helper.GetParamId(getParam, "movieId")

// 		getOnePhoto, err := s.photoRepository.FindById(getParam.ID)

// 		if err != nil {
// 			ctx.AbortWithStatusJSON(404, err)
// 			return
// 		}

// 		if getOnePhoto.UserId != currentUser.ID {
// 			unauthorizedErr := errors.New("you are not authorized to modify the movie data")
// 			ctx.AbortWithStatusJSON(403, unauthorizedErr)
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

// func (s *jwtService) CommentAuthorization() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var getParam comment.GetCommentInput
// 		err := ctx.ShouldBindUri(&getParam)
// 		if err != nil {
// 			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
// 			ctx.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}

// 		currentUser := ctx.MustGet("currentUser").(user.User)
// 		// input.User.ID = currentUser.ID
// 		// user := ctx.MustGet("userData").(entity.User)

// 		// movieId, err := helper.GetParamId(getParam, "movieId")

// 		getOnePhoto, err := s.commentRepository.FindById(getParam.ID)

// 		if err != nil {
// 			ctx.AbortWithStatusJSON(404, err)
// 			return
// 		}

// 		if getOnePhoto.UserId != currentUser.ID {
// 			unauthorizedErr := errors.New("you are not authorized to modify the movie data")
// 			ctx.AbortWithStatusJSON(403, unauthorizedErr)
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

// func (s *jwtService) SosmedAuthorization() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		var getParam sosialMedia.GetSosmedInput
// 		err := ctx.ShouldBindUri(&getParam)
// 		if err != nil {
// 			response := helper.APIresponse(http.StatusUnprocessableEntity, err)
// 			ctx.JSON(http.StatusUnprocessableEntity, response)
// 			return
// 		}

// 		currentUser := ctx.MustGet("currentUser").(user.User)
// 		// input.User.ID = currentUser.ID
// 		// user := ctx.MustGet("userData").(entity.User)

// 		// movieId, err := helper.GetParamId(getParam, "movieId")

// 		getOnePhoto, err := s.commentRepository.FindById(getParam.ID)

// 		if err != nil {
// 			ctx.AbortWithStatusJSON(404, err)
// 			return
// 		}

// 		if getOnePhoto.UserId != currentUser.ID {
// 			unauthorizedErr := errors.New("you are not authorized to modify the movie data")
// 			ctx.AbortWithStatusJSON(403, unauthorizedErr)
// 			return
// 		}

// 		ctx.Next()
// 	}
// }

