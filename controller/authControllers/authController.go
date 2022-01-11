package authController

import (
	appContext "context"
	"edu_app_backend/dto"
	"edu_app_backend/models"
	"edu_app_backend/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var userCollection = service.ConnectToCollection("User")

func comparePasswords(passWordInDb string, passWordFromClient string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passWordInDb), []byte(passWordFromClient))
	return err == nil
}

func responseBadRequest(context *fiber.Ctx, errorMessage string) {
	responseData := dto.LoginUserResponse{
		Success:      false,
		Reason:       errorMessage,
		UserLoggedIn: false,
	}

	context.JSON(fiber.Map{
		"success":      responseData.Success,
		"reason":       responseData.Reason,
		"userLoggedIn": responseData.UserLoggedIn,
	})

}

func LoginUser(context *fiber.Ctx) error {
	context.Accepts("application/json")
	loginRequest := new(dto.LoginRequest)
	err := context.BodyParser(loginRequest)
	filter := bson.M{
		"EmailAddress": loginRequest.EmailAddress,
	}

	if err != nil {
		responseBadRequest(context, "An Error Occured ")
	}

	result, error := userCollection.Find(appContext.Background(), filter)

	if error != nil {
		responseBadRequest(context, "Email / Password is Incorrect")
	}

	var userDataArr []models.User

	if err := result.All(appContext.TODO(), &userDataArr); err != nil {
		responseBadRequest(context, "An Error Occured")
	} else {
		log.Println(userDataArr)
		passWordFromClient := loginRequest.Password
		mainUserData := userDataArr[0]
		userPasswordHash := mainUserData.PassHash
		isPasswordCorrect := comparePasswords(userPasswordHash, passWordFromClient)
		if isPasswordCorrect {
			responseObject := dto.LoginUserResponse{
				Success:      true,
				Reason:       "Logged In Successfully",
				UserId:       mainUserData.ID_,
				UserLoggedIn: true,
			}
			context.JSON(fiber.Map{
				"success":      responseObject.Success,
				"reason":       responseObject.Reason,
				"userId":       responseObject.UserId,
				"userLoggedIn": responseObject.UserLoggedIn,
			})
		}
	}

	return err
}
