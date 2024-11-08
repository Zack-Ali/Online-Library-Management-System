package controllers

import (
	"context"
	"net/http"
	"online_library/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection

func SetUserCollection(c *mongo.Collection) {
	userCollection = c
}
func GetAllUsers(c *gin.Context) {
	user := []models.User{}
	reslut, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error getting users"})
		return
	}
	defer reslut.Close(context.Background())
	for reslut.Next(context.Background()) {
		sp_user := models.User{}
		if err := reslut.Decode(&sp_user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "Error decoding  users"})
			return
		}
		user = append(user, sp_user)

	}

	c.JSON(http.StatusOK, user)
}
func CreateUser(c *gin.Context) {
	user := new(models.User)
	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error parsing user"})
		return
	}
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter you username.."})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter you email.."})
		return
	}
	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter password .."})
		return
	}
	// Check if user already exists
	var userExist models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&userExist)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"msg": "User already exists!.."})
		return
	}
	insertData, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error inserted User"})
		return
	}
	user.ID = insertData.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusOK, gin.H{"msg": "User creating successfully.."})

}
func UserLogin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid data"})
		return
	}

	// Check if user exists
	var existingUser models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": user.Email}).Decode(&existingUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "User not found"})
		return
	}

	// Compare passwords (add password hashing for security)
	if existingUser.Password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid password, please enter a valid password."})
		return
	}

	token := "sas-sa-sas"

	c.JSON(http.StatusOK, gin.H{"msg": "Login successful", "token": token, "data": existingUser})
}

func UpdateUser(c *gin.Context) {
	// updateUser := new(models.User)
	// id := c.Params.ByName("id")
	// updated_id, err := primitive.ObjectIDFromHex(id)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid Id, please try again?"})
	// 	return
	// }
	// if err := c.ShouldBindJSON(&calorieTrackerEntire); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"msg": "Error parsing Calorie Tracker"})
	// 	return
	// }
	// update_fields := bson.M{}
	// if calorieTrackerEntire.Dish != "" {
	// 	update_fields["dish"] = calorieTrackerEntire.Dish
	// }
	// if calorieTrackerEntire.Fat != 0 {
	// 	update_fields["fat"] = calorieTrackerEntire.Fat
	// }
	// if calorieTrackerEntire.Ingredients != "" {
	// 	update_fields["ingredients"] = calorieTrackerEntire.Ingredients
	// }
	// if calorieTrackerEntire.Calories != 0 {
	// 	update_fields["calories"] = calorieTrackerEntire.Calories
	// }
	// filter := bson.M{"_id": updated_id}
	// update := bson.M{"$set": update_fields}
	// _, err = collection.UpdateOne(context.Background(), filter, update)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error updating calorie tracking"})
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{"msg": "Updated successfully.."})
}
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	deleted_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": deleted_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error deleting user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully.."})
}
