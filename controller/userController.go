package controller

import (
	appContext "context"
	"edu_app_backend/dto"
	"edu_app_backend/service"
	"edu_app_backend/types"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

var MongodDbClient = service.ConnectDb()

var DataBase = MongodDbClient.Database("EduCoApp")

type CreateUserResponse struct {
	Success bool        `json:"success"`
	Reason  string      `json:"reason"`
	UserId  interface{} `json:"userId"`
}

func creatPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func respondError(errorMessage string, context *fiber.Ctx) {
	responseData := CreateUserResponse{
		Success: false,
		Reason:  errorMessage,
	}
	context.JSON(fiber.Map{
		"success": responseData.Success,
		"reason":  responseData.Reason,
	})
}

func CreateUser(context *fiber.Ctx) error {
	context.Accepts("application/json")
	user := new(dto.User)
	err := context.BodyParser(user)
	if err != nil {
		fmt.Println("An Error Occured")
		responseData := CreateUserResponse{
			Success: false,
			Reason:  "An Error Occured",
		}
		context.JSON(fiber.Map{
			"success": responseData.Success,
			"reason":  err.Error(),
		})
		return err
	} else {

		filter := bson.M{"EmailAddress": user.EmailAddress}

		result, err := DataBase.Collection("User").Find(appContext.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}

		var usersFiltered []bson.M
		if err = result.All(appContext.Background(), &usersFiltered); err != nil {
			log.Fatal(err)
		}

		if len(usersFiltered) == 0 {
			hashedPassword, err := creatPasswordHash(user.Password)
			if err == nil {
				var emptyStringArray = make([]string, 0)
				var emptySocialMediaLinkArray = make([]types.UserSocialLink, 0)
				doc := bson.D{{Key: "UserName", Value: user.UserName}, {Key: "PassHash", Value: hashedPassword}, {Key: "EmailAddress", Value: user.EmailAddress}, {Key: "Courses", Value: emptyStringArray}, {Key: "Followers", Value: emptyStringArray}, {Key: "Following", Value: emptyStringArray}, {Key: "ProfileDescription", Value: ""}, {Key: "SocialMediaLinks", Value: emptySocialMediaLinkArray}}
				result, err := DataBase.Collection("User").InsertOne(appContext.TODO(), doc)
				if err == nil {
					responseData := CreateUserResponse{
						Success: true,
						UserId:  result.InsertedID,
					}
					context.JSON(fiber.Map{
						"success": responseData.Success,
						"UserId":  responseData.UserId,
					})

				} else {
					respondError(err.Error(), context)
				}

			} else {
				context.Status(301)
				panic(err)
			}
		} else {
			respondError("User Already Exists", context)
		}
	}

	return err
}
