// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: copyfrom.go

package database

import (
	"context"
)

// iteratorForAddIngredientsToRecipe implements pgx.CopyFromSource.
type iteratorForAddIngredientsToRecipe struct {
	rows                 []AddIngredientsToRecipeParams
	skippedFirstNextCall bool
}

func (r *iteratorForAddIngredientsToRecipe) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForAddIngredientsToRecipe) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Amount,
		r.rows[0].PrepNote,
		r.rows[0].CreatedAt,
		r.rows[0].UpdatedAt,
		r.rows[0].IngredientID,
		r.rows[0].RecipeID,
		r.rows[0].Index,
	}, nil
}

func (r iteratorForAddIngredientsToRecipe) Err() error {
	return nil
}

func (q *Queries) AddIngredientsToRecipe(ctx context.Context, arg []AddIngredientsToRecipeParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"recipe_ingredient"}, []string{"amount", "prep_note", "created_at", "updated_at", "ingredient_id", "recipe_id", "index"}, &iteratorForAddIngredientsToRecipe{rows: arg})
}
