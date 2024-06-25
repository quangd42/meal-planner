// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: recipe_ingredient.sql

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

const updateIngredientInRecipe = `-- name: UpdateIngredientInRecipe :exec
UPDATE recipe_ingredient
SET
  amount = $1,
  instruction = $2,
  updated_at = $3
WHERE
  ingredient_id = $4 AND recipe_id = $5
`

type UpdateIngredientInRecipeParams struct {
	Amount       string      `json:"amount"`
	Instruction  pgtype.Text `json:"instruction"`
	UpdatedAt    time.Time   `json:"updated_at"`
	IngredientID pgtype.UUID `json:"ingredient_id"`
	RecipeID     pgtype.UUID `json:"recipe_id"`
}

func (q *Queries) UpdateIngredientInRecipe(ctx context.Context, arg UpdateIngredientInRecipeParams) error {
	_, err := q.db.Exec(ctx, updateIngredientInRecipe,
		arg.Amount,
		arg.Instruction,
		arg.UpdatedAt,
		arg.IngredientID,
		arg.RecipeID,
	)
	return err
}
