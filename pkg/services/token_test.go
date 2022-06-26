package services

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/abelgalef/course-reg/pkg/models"
)

func TestTokenService(t *testing.T) {
	tokenService := NewJWTService()

	// SINCE THERE ARE NO EDGE CASES FOR THIS SERVICE, WE CAN JUST TEST IT WITH A SIMPLE STRUCT
	user := models.User{FirstName: "Abel", LastName: "Galef", Email: "blah_blah"}

	// GENERATE THE TOKEN
	token, err := tokenService.GenerateToken(user)
	if err != nil {
		t.Fatalf("Error generating token: %s", err)
	}

	fmt.Println(token)

	// VALIDATE THE GENERATED TOKEN
	data, validated := tokenService.ValidateToken(token)
	if !validated {
		t.Fatal("Token validation failed")
	}

	// THE JWT SERVICE ALWAYS RETURNS A MAP, AND GO CURRENTLY DOSEN'T SUPPORT CONVERTING MAPS TO STRUCTS SO WE NEED TO MARSHAL IT TO A JSON AND THEN UNMARSHAL THAT TO A STRUCT
	// AN UNNECESSARY BUT ESSENTIAL IMPLEMENTATION DUE TO GO'S SHORT COMINGS
	jsonBody, err := json.Marshal(&data)
	if err != nil {
		t.Fatal("Could not marshal the employee map to a json")
	}
	var empModel models.User
	if err := json.Unmarshal(jsonBody, &empModel); err != nil {
		t.Fatal("Could not unmarshal the json to an employee model")
	}

	if empModel.FirstName != "Abel" {
		t.Fatalf("Token data is incorrect or corrupted \nexcpected %v got %v", user, empModel)
	}
}
