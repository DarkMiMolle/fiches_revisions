package controller

import (
	"context"
	gerrors "errors"
	"fmt"
	"github.com/DarkMiMolle/Fiche/backend/env"
	"github.com/DarkMiMolle/Fiche/backend/errors"
	"github.com/DarkMiMolle/Fiche/backend/models"
	"github.com/DarkMiMolle/Fiche/backend/utils"
	"github.com/DarkMiMolle/GTL/optional"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

func SignUp(c *gin.Context) {
	// get user from context
	var body struct {
		Pseudo   string                 `json:"pseudo"`
		Password string                 `json:"password"`
		Name     optional.Value[string] `json:"name"`
	}
	if err := c.Bind(&body); err != nil {
		panic(errors.BadRequest(err))
	}

	if len(body.Password) < 5 {
		panic(errors.NewBadRequest("invalid password, requires at least 5 character"))
		return
	}

	user := models.User{
		Pseudo: body.Pseudo,
		Name:   body.Name.ValueOr(fmt.Sprintf("Unnamed_%v", time.Now().Unix())),
	}

	// check if user email exists in db
	db := utils.ExtractValues(c)
	collection := db.Collection(os.Getenv(env.UserCollection))
	result := collection.FindOne(context.Background(), user.MongoIDFilter())
	if result.Err() == nil {
		panic(errors.NewBadRequest("user with same pseudo (%v) already exists", user.Pseudo))
	}
	if !gerrors.Is(result.Err(), mongo.ErrNoDocuments) {
		panic(errors.Internal(result.Err()))
	}

	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if gerrors.Is(err, bcrypt.ErrPasswordTooLong) {
		panic(errors.BadRequest(err))
	}
	if err != nil {
		panic(errors.Internal(err))
	}

	// register the user
	newUser := models.AuthUser{
		User:     user,
		Password: models.Password(hashPassword),
	}
	_, err = collection.InsertOne(context.Background(), newUser)
	if err != nil {
		panic(errors.Internal(err))
	}

	c.JSON(utils.Success(user))
}

func Login(c *gin.Context) {
	// get login information from body
	var body struct {
		Pseudo   string `json:"pseudo"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		panic(errors.BadRequest(err))
	}
	fmt.Printf("%+v\n", body)
	// check if there is a match in db
	db := utils.ExtractValues(c)
	userCollection := db.Collection(os.Getenv(env.UserCollection))

	result := userCollection.FindOne(c, models.User{Pseudo: body.Pseudo}.MongoIDFilter())
	if result.Err() != nil && gerrors.Is(result.Err(), mongo.ErrNoDocuments) {
		panic(errors.BadRequest(result.Err()))
	}
	if result.Err() != nil {
		panic(errors.Internal(result.Err()))
	}
	var user models.AuthUser
	if err := result.Decode(&user); err != nil {
		panic(errors.Internal(err))
	}
	if !user.Password.Match(body.Password) {
		panic(errors.NewBadRequest("invalid password"))
	}

	// create and return token
	token, err := utils.GenerateJwt(&user)
	if err != nil {
		panic(errors.Internal(err))
	}
	c.JSON(utils.Success(gin.H{"jwt": token}))
}
