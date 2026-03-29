-- name: CreateExpense :one
INSERT INTO expenses (title, amount)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateExpense :one 
UPDATE expenses
SET 
  title = $2,
  amount = $3
WHERE 
  id = $1
RETURNING *;

-- name: GetExpenseByID :one 
SELECT * 
FROM expenses
WHERE id = $1;

-- name: GetAllExpenses :many
SELECT * 
FROM expenses
ORDER BY id DESC;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = $1;
