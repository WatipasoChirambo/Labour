package controllers

import (
	"context"
	"github.com/WatipasoChirambo/go-lang-roado/configs"
	"github.com/WatipasoChirambo/go-lang-roado/models"
	"github.com/WatipasoChirambo/go-lang-roado/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var labourerCollection *mongo.Collection = configs.GetCollection(configs.DB, "labourers")
var validate = validator.New()

func CreateLabourer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var labourer models.Labourer
		defer cancel()

		//validate request body
		if err := c.BindJSON(&labourer); err != nil {
			c.JSON(http.StatusBadRequest, responses.LabourerResponse{Status: http.StatusBadRequest, Message: "Erro", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//validate required fields
		if validationError := validate.Struct(&labourer); validationError != nil {
			c.JSON(http.StatusBadRequest, responses.LabourerResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": validationError.Error()}})
			return
		}

		newLabourer := models.Labourer{
			Id:                  primitive.NewObjectID(),
			RegNo:               labourer.RegNo,
			FirstName:           labourer.FirstName,
			LastName:            labourer.LastName,
			Gender:              labourer.Gender,
			DistrictOfBirth:     labourer.DistrictOfBirth,
			Address:             labourer.Address,
			HomeAddress:         labourer.HomeAddress,
			Village:             labourer.Village,
			Ta:                  labourer.Ta,
			District:            labourer.District,
			MaritalStatus:       labourer.MaritalStatus,
			PreviousResidence:   labourer.PreviousResidence,
			CurrentResidence:    labourer.CurrentResidence,
			Education:           labourer.Education,
			Institution:         labourer.Institution,
			Profession:          labourer.Profession,
			JobWanted:           labourer.JobWanted,
			PreviousJob:         labourer.PreviousJob,
			YearsUnemployed:     labourer.YearsUnemployed,
			ExperienceYears:     labourer.ExperienceYears,
			Training:            labourer.Training,
			RegisteredElseWhere: labourer.RegisteredElseWhere,
			RegisteredNumber:    labourer.RegisteredNumber,
		}

		result, err := labourerCollection.InsertOne(ctx, newLabourer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LabourerResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.LabourerResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetLabourer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		labourerId := c.Param("id")
		var labourer models.Labourer
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(labourerId)

		err := labourerCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&labourer)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LabourerResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.LabourerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": labourer}})
	}
}

func UpdateLabourer() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		labourerId := c.Param("id")
		var labourer models.Labourer
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(labourerId)

		if err := c.BindJSON(&labourer); err != nil {
			c.JSON(http.StatusBadRequest, responses.LabourerResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&labourer); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.LabourerResponse{Status: http.StatusBadRequest, Message: "Error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"RegNo": labourer.RegNo, "FirstName": labourer.FirstName, "LastName": labourer.LastName, "Gender": labourer.Gender, "DistrictOfBirth": labourer.DistrictOfBirth, "Address": labourer.Address, "HomeAddress": labourer.HomeAddress, "Village": labourer.Village,"Ta":labourer.Ta,"District":labourer.District,"MaritalStatus":labourer.MaritalStatus,"PreviousResidence":labourer.PreviousResidence,"CurrentResidence":labourer.CurrentResidence,"Education":labourer.Education,"Institution":labourer.Institution,"Profession":labourer.Profession,"JobWanted":labourer.JobWanted,"PreviousJob":labourer.PreviousJob,"YearsUnemployed":labourer.YearsUnemployed,"ExperienceYears":labourer.ExperienceYears,"Training":labourer.Training,"RegisteredElseWhere":labourer.RegisteredElseWhere,"RegisteredNumber":labourer.RegisteredNumber}
		result, err := labourerCollection.UpdateOne(ctx, bson.M{"id":objId}, bson.M{"$set":update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.LabourerResponse{Status: http.StatusInternalServerError, Message:"Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedLabourer models.Labourer
		if result.MatchedCount == 1{
			err := labourerCollection.FindOne(ctx, bson.M{"id":objId}).Decode(&updatedLabourer)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.LabourerResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.LabourerResponse{Status: http.StatusOK, Message:"success", Data: map[string]interface{}{"data":updatedLabourer}})
	}
}

func GetAllLabourers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var labourers []models.Labourer
        defer cancel()
		result, err := labourerCollection.Find(ctx, bson.M{})

        if err!= nil {
			c.JSON(http.StatusInternalServerError, responses.LabourerResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		defer result.Close(ctx)
		for result.Next(ctx) {
			var singleLabourer models.Labourer
			if err= result.Decode(&singleLabourer); err != nil {
				c.JSON(http.StatusInternalServerError, responses.LabourerResponse{Status: http.StatusInternalServerError, Message: "Error", Data: map[string]interface{}{"data": err.Error()}})
			}
			labourers = append(labourers, singleLabourer)
		}
		c.JSON(http.StatusOK, responses.LabourerResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": labourers}})
	}
}
