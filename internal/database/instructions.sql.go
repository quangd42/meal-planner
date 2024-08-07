// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: instructions.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addInstructionToRecipe = `-- name: AddInstructionToRecipe :exec
INSERT INTO instructions (
  created_at, updated_at, instruction, step_no, recipe_id
) VALUES ($1, $2, $3, $4, $5)
`

type AddInstructionToRecipeParams struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Instruction string    `json:"instruction"`
	StepNo      int32     `json:"step_no"`
	RecipeID    uuid.UUID `json:"recipe_id"`
}

func (q *Queries) AddInstructionToRecipe(ctx context.Context, arg AddInstructionToRecipeParams) error {
	_, err := q.db.Exec(ctx, addInstructionToRecipe,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Instruction,
		arg.StepNo,
		arg.RecipeID,
	)
	return err
}

const deleteInstructionByID = `-- name: DeleteInstructionByID :exec
DELETE FROM instructions
WHERE step_no = $1 AND recipe_id = $2
`

type DeleteInstructionByIDParams struct {
	StepNo   int32     `json:"step_no"`
	RecipeID uuid.UUID `json:"recipe_id"`
}

func (q *Queries) DeleteInstructionByID(ctx context.Context, arg DeleteInstructionByIDParams) error {
	_, err := q.db.Exec(ctx, deleteInstructionByID, arg.StepNo, arg.RecipeID)
	return err
}

const listInstructionsByRecipeID = `-- name: ListInstructionsByRecipeID :many
SELECT created_at, updated_at, step_no, instruction, recipe_id
FROM instructions
WHERE recipe_id = $1
ORDER BY step_no ASC
`

func (q *Queries) ListInstructionsByRecipeID(ctx context.Context, recipeID uuid.UUID) ([]Instruction, error) {
	rows, err := q.db.Query(ctx, listInstructionsByRecipeID, recipeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Instruction
	for rows.Next() {
		var i Instruction
		if err := rows.Scan(
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.StepNo,
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

const updateInstructionByID = `-- name: UpdateInstructionByID :exec
UPDATE instructions
SET updated_at = $3, instruction = $4
WHERE step_no = $1 AND recipe_id = $2
`

type UpdateInstructionByIDParams struct {
	StepNo      int32     `json:"step_no"`
	RecipeID    uuid.UUID `json:"recipe_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	Instruction string    `json:"instruction"`
}

func (q *Queries) UpdateInstructionByID(ctx context.Context, arg UpdateInstructionByIDParams) error {
	_, err := q.db.Exec(ctx, updateInstructionByID,
		arg.StepNo,
		arg.RecipeID,
		arg.UpdatedAt,
		arg.Instruction,
	)
	return err
}
