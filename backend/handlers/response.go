package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/quangd42/meal-planner/backend/internal/database"
)

func respondJSON[T any](w http.ResponseWriter, code int, v T) {
	data, err := json.Marshal(v)
	if err != nil {
		log.Printf("error decoding JSON: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Printf("Responding with 5xx error: %s\n", msg)
	}
	type response struct {
		Error string `json:"error"`
	}
	respondJSON(w, code, response{
		Error: msg,
	})
}

func respondInternalServerError(w http.ResponseWriter) {
	respondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func respondUniqueValueError(w http.ResponseWriter, err error, msg string) {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		respondError(w, http.StatusBadRequest, "unique value constraint violated: "+msg)
		return
	}
	respondError(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}

func respondMalformedRequestError(w http.ResponseWriter) {
	respondError(w, http.StatusBadRequest, "malformed request body")
}

type User struct {
	ID           uuid.UUID `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Hash         string    `json:"-"`
	Token        string    `json:"token,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
}

func createUserResponse(u database.User) User {
	return User{
		ID:        u.ID.Bytes,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		Username:  u.Username,
	}
}

func createUserResponseWithToken(u database.User, token, refreshToken string) User {
	user := createUserResponse(u)
	user.Token = token
	user.RefreshToken = refreshToken
	return user
}

type Recipe struct {
	ID                uuid.UUID             `json:"id"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	Name              string                `json:"name"`
	ExternalUrl       string                `json:"external_url"`
	UserID            uuid.UUID             `json:"user_id"`
	Servings          int                   `json:"servings"`
	Yield             string                `json:"yield"`
	CookTimeInMinutes int                   `json:"cook_time_in_minutes"`
	Notes             string                `json:"notes"`
	Ingredients       []IngredientInRecipe  `json:"ingredients"`
	Instructions      []InstructionInRecipe `json:"instructions"`
}

type IngredientInRecipe struct {
	ID       uuid.UUID `json:"id"`
	Amount   string    `json:"amount"`
	PrepNote string    `json:"prep_note"`
	Name     string    `json:"name"`
}

type InstructionInRecipe struct {
	StepNo      int    `json:"step_no"`
	Instruction string `json:"instruction"`
}

func createRecipeResponse(recipe database.Recipe, dbIngredients []database.ListIngredientsByRecipeIDRow, DBInstructions []database.Instruction) Recipe {
	ingredients := []IngredientInRecipe{}
	for _, di := range dbIngredients {
		ingredients = append(ingredients, IngredientInRecipe{
			ID:       di.ID.Bytes,
			Amount:   di.Amount,
			PrepNote: di.PrepNote.String,
			Name:     di.Name,
		})
	}

	instructions := []InstructionInRecipe{}
	for _, di := range DBInstructions {
		instructions = append(instructions, InstructionInRecipe{
			StepNo:      int(di.StepNo),
			Instruction: di.Instruction,
		})
	}

	return Recipe{
		ID:                recipe.ID.Bytes,
		CreatedAt:         recipe.CreatedAt,
		UpdatedAt:         recipe.UpdatedAt,
		Name:              recipe.Name,
		ExternalUrl:       recipe.ExternalUrl,
		UserID:            recipe.UserID.Bytes,
		Servings:          int(recipe.Servings),
		Yield:             recipe.Yield.String,
		CookTimeInMinutes: int(recipe.CookTimeInMinutes),
		Notes:             recipe.Notes.String,
		Ingredients:       ingredients,
		Instructions:      instructions,
	}
}

type Ingredient struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Name      string     `json:"name"`
	ParentID  *uuid.UUID `json:"parent_id"`
}

func createIngredientResponse(i database.Ingredient) Ingredient {
	res := Ingredient{
		ID:        i.ID.Bytes,
		Name:      i.Name,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		ParentID:  &uuid.UUID{},
	}
	if i.ParentID.Valid {
		*res.ParentID = i.ParentID.Bytes
	} else {
		res.ParentID = nil
	}
	return res
}
