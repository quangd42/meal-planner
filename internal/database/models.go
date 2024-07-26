// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Cuisine struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Name      string     `json:"name"`
	ParentID  *uuid.UUID `json:"parent_id"`
}

type Ingredient struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

type Instruction struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	StepNo      int32     `json:"step_no"`
	Instruction string    `json:"instruction"`
	RecipeID    uuid.UUID `json:"recipe_id"`
}

type Recipe struct {
	ID                uuid.UUID `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ExternalUrl       *string   `json:"external_url"`
	Name              string    `json:"name"`
	UserID            uuid.UUID `json:"user_id"`
	Servings          int32     `json:"servings"`
	Yield             *string   `json:"yield"`
	CookTimeInMinutes int32     `json:"cook_time_in_minutes"`
	Notes             *string   `json:"notes"`
}

type RecipeCuisine struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CuisineID uuid.UUID `json:"cuisine_id"`
	RecipeID  uuid.UUID `json:"recipe_id"`
}

type RecipeIngredient struct {
	Index        int32     `json:"index"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Amount       string    `json:"amount"`
	PrepNote     *string   `json:"prep_note"`
	IngredientID uuid.UUID `json:"ingredient_id"`
	RecipeID     uuid.UUID `json:"recipe_id"`
}

type Token struct {
	Value     string    `json:"value"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
	IsRevoked bool      `json:"is_revoked"`
	UserID    uuid.UUID `json:"user_id"`
}

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Hash      string    `json:"hash"`
}
