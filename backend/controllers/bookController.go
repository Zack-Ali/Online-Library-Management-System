package controllers

import (
	"context"
	"log"
	"net/http"
	"online_library/models"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection *mongo.Collection

func SetCollection(c *mongo.Collection) {
	collection = c
}
func GetAllBooks(c *gin.Context) {
	books := []models.Book{}
	reslut, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error getting Books"})
		return
	}
	defer reslut.Close(context.Background())

	bookChanel := make(chan models.Book)
	errorChannel := make(chan error)

	go func() {
		for reslut.Next(context.Background()) {
			book := models.Book{}
			if err := reslut.Decode(&book); err != nil {
				errorChannel <- err
				return
			}
			bookChanel <- book
		}
		close(bookChanel)
	}()
	go func() {
		for err := range errorChannel {
			log.Fatal("Error decoding Books", err)
		}
	}()
	for book := range bookChanel {
		books = append(books, book)
	}
	c.JSON(http.StatusOK, books)
}
func GetBook(c *gin.Context) {
	book := models.Book{}
	id := c.Params.ByName("id")
	object_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": object_id}
	err = collection.FindOne(context.Background(), filter).Decode(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error fetching book"})
		return
	}
	c.JSON(http.StatusOK, book)
}
func CreateBook(c *gin.Context) {
	book := new(models.Book)
	if err := c.ShouldBindJSON(book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error parsing Book"})
		return
	}
	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter title of book.."})
		return
	}
	if book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter author of book.."})
		return
	}
	if book.PublishedYear == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Please enter published year of book.."})
		return
	}

	insertData, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error inserted Book"})
		return
	}
	book.ID = insertData.InsertedID.(primitive.ObjectID)
	c.JSON(http.StatusOK, gin.H{"msg": "Inseted successfully.."})

}
func UpdateBook(c *gin.Context) {
	book := new(models.Book)
	id := c.Params.ByName("id")
	updated_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Error parsing book"})
		return
	}
	update_fields := bson.M{}
	if book.Title != "" {
		update_fields["title"] = book.Title
	}
	if book.Author != "" {
		update_fields["author"] = book.Author
	}
	if book.PublishedYear != 0 {
		update_fields["publishedYear"] = book.PublishedYear
	}

	filter := bson.M{"_id": updated_id}
	update := bson.M{"$set": update_fields}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error updating book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Updated successfully.."})
}
func DeleteBook(c *gin.Context) {
	id := c.Params.ByName("id")
	deleted_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": "Invalid Id, please try again?"})
		return
	}
	filter := bson.M{"_id": deleted_id}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "Error deleting book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Deleted successfully.."})
}
