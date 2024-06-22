// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: ingredients.sql

package database

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type AddIngredientsToRecipeParams struct {
	Amount       string      `json:"amount"`
	Instruction  pgtype.Text `json:"instruction"`
	CreatedAt    time.Time   `json:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at"`
	IngredientID pgtype.UUID `json:"ingredient_id"`
	RecipeID     pgtype.UUID `json:"recipe_id"`
}

const createIngredient = `-- name: CreateIngredient :one
INSERT INTO ingredients (id, created_at, updated_at, name, parent_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, name, parent_id
`

type CreateIngredientParams struct {
	ID        pgtype.UUID `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Name      string      `json:"name"`
	ParentID  pgtype.UUID `json:"parent_id"`
}

func (q *Queries) CreateIngredient(ctx context.Context, arg CreateIngredientParams) (Ingredient, error) {
	row := q.db.QueryRow(ctx, createIngredient,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.ParentID,
	)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ParentID,
	)
	return i, err
}

const getIngredientByID = `-- name: GetIngredientByID :one
SELECT id, created_at, updated_at, name, parent_id
FROM ingredients
WHERE id = $1
`

func (q *Queries) GetIngredientByID(ctx context.Context, id pgtype.UUID) (Ingredient, error) {
	row := q.db.QueryRow(ctx, getIngredientByID, id)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ParentID,
	)
	return i, err
}

const listIngredientsByRecipeID = `-- name: ListIngredientsByRecipeID :many
SELECT
  id,
  name,
  amount,
  instruction,
  recipe_id
FROM ingredients
JOIN recipe_ingredient ON id = ingredient_id
WHERE recipe_id = $1
`

type ListIngredientsByRecipeIDRow struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Amount      string      `json:"amount"`
	Instruction pgtype.Text `json:"instruction"`
	RecipeID    pgtype.UUID `json:"recipe_id"`
}

func (q *Queries) ListIngredientsByRecipeID(ctx context.Context, recipeID pgtype.UUID) ([]ListIngredientsByRecipeIDRow, error) {
	rows, err := q.db.Query(ctx, listIngredientsByRecipeID, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListIngredientsByRecipeIDRow
	for rows.Next() {
		var i ListIngredientsByRecipeIDRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Amount,
			&i.Instruction,
			&i.RecipeID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateIngredientByID = `-- name: UpdateIngredientByID :one
UPDATE ingredients
SET
  name = $2,
  parent_id = $3,
  updated_at = $4
WHERE id = $1
RETURNING id, created_at, updated_at, name, parent_id
`

type UpdateIngredientByIDParams struct {
	ID        pgtype.UUID `json:"id"`
	Name      string      `json:"name"`
	ParentID  pgtype.UUID `json:"parent_id"`
	UpdatedAt time.Time   `json:"updated_at"`
}

func (q *Queries) UpdateIngredientByID(ctx context.Context, arg UpdateIngredientByIDParams) (Ingredient, error) {
	row := q.db.QueryRow(ctx, updateIngredientByID,
		arg.ID,
		arg.Name,
		arg.ParentID,
		arg.UpdatedAt,
	)
	var i Ingredient
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ParentID,
	)
	return i, err
}
