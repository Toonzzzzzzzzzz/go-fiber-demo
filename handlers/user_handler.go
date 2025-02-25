package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/Toonzzzzzzzzzz/go-fiber-demo/database"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "test",
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []map[string]any
	cursor, err := database.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "ไม่สามารถดึงข้อมูลได้"})
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "ไม่สามารถดึงข้อมูลได้"})
	}
	if len(users) == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ไม่พบข้อมูล"})
	}
	return c.JSON(fiber.Map{
		"message": "ดึงข้อมูลสําเร็จ",
		"data":    users,
	})
}

func CreateUser(c *fiber.Ctx) error {
	var user map[string]any
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	if user["username"] == nil || user["password"] == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ป้อน Username and password"})
	}
	var checkUser map[string]any
	err := database.UserCollection.FindOne(context.TODO(), bson.M{"username": user["username"]}).Decode(&checkUser)
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Username ซ้ำ"})
	}

	user["created_at"] = time.Now()
	user["updated_at"] = time.Now()
	result, err := database.UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "ไม่สามารถเพิ่มข้อมูลได้"})
	}
	return c.JSON(fiber.Map{
		"message":  "เพิ่มข้อมูลสำเร็จ",
		"id":       result.InsertedID,
		"username": user["username"],
	})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "รหัส ID ไม่ถูกต้อง"})
	}
	result, err := database.UserCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "ไม่สามารถลบข้อมูลได้"})
	}
	if result.DeletedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ไม่พบข้อมูล"})
	}
	return c.JSON(fiber.Map{
		"message": "ลบข้อมูลสําเร็จ",
	})
}

func GetUsersById(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "รหัส ID ไม่ถูกต้อง"})
	}
	var user map[string]any
	err = database.UserCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ไม่พบข้อมูล"})
	}
	return c.JSON(fiber.Map{
		"message": "ดึงข้อมูลสําเร็จ",
		"data":    user,
	})
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "รหัส ID ไม่ถูกต้อง"})
	}

	var user map[string]any
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ข้อมูลไม่ถูกต้อง"})
	}
	if len(user) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "ไม่มีข้อมูลสำหรับอัปเดต"})
	}

	updateFields := bson.M{}
	if username, exists := user["username"]; exists {
		updateFields["username"] = username
	}
	if password, exists := user["password"]; exists {
		updateFields["password"] = password
	}

	updateFields["updated_at"] = time.Now()
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updateFields}

	result, err := database.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "ไม่สามารถอัปเดตข้อมูลได้"})
	}
	if result.MatchedCount == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "ไม่พบผู้ใช้ที่ต้องการอัปเดต"})
	}

	return c.JSON(fiber.Map{
		"message":        "แก้ไขข้อมูลสำเร็จ",
		"updated_fields": updateFields,
	})
}
