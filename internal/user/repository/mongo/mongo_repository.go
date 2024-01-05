package mongo

import (
	"Human_Resources_Managament_System/models/employee"
	"Human_Resources_Managament_System/pkg/mongodb"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepo struct {
	db mongodb.MongoDBManager
}

func NewUserRepo(db mongodb.MongoDBManager) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetAllEmployeesData(ctx context.Context) ([]employee.EmployeeData, error) {
	query := bson.D{{}}

	cursor, err := u.db.Database.Collection("employees").Find(ctx, query)
	if err != nil {
		return []employee.EmployeeData{}, err
	}

	var employees []employee.EmployeeData
	if err := cursor.All(ctx, &employees); err != nil {
		return []employee.EmployeeData{}, err
	}

	return employees, nil
}

func (u *UserRepo) CreateEmployee(ctx context.Context, info employee.Employee) (employee.EmployeeData, error) {
	employeeData := info.ToStorage()
	collection := u.db.Database.Collection("employees")

	employeeData.ID = ""

	insertionResult, err := collection.InsertOne(ctx, employeeData)
	if err != nil {
		return employee.EmployeeData{}, err
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(ctx, filter)

	var createdEmployee employee.EmployeeData
	if err := createdRecord.Decode(&createdEmployee); err != nil {
		return employee.EmployeeData{}, err
	}

	return createdEmployee, nil
}
func (u *UserRepo) UpdateEmployeeByID(ctx context.Context, id string, info employee.Employee) error {
	employeeData := info.ToStorage()
	collection := u.db.Database.Collection("employees")

	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employeeData.Name},
				{Key: "age", Value: employeeData.Age},
				{Key: "salary", Value: employeeData.Salary},
			},
		},
	}

	err = collection.FindOneAndUpdate(ctx, query, update).Err()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) DeleteEmployeeByID(ctx context.Context, id string) error {
	employeeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	query := bson.D{{Key: "_id", Value: employeeID}}
	collection := u.db.Database.Collection("employees")

	result, err := collection.DeleteOne(ctx, query)
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return errors.New("Not Found")
	}

	return nil
}
