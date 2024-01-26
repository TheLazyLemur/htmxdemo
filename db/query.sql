-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = ? LIMIT 1;

-- name: ListTransactions :many
SELECT * FROM transactions;

-- name: FlagTransaction :exec
UPDATE transactions
set flagged = true
WHERE id = ?;

-- name: UnflagTransaction :exec
UPDATE transactions
set flagged = false
WHERE id = ?;

-- name: InsertTransaction :exec
INSERT INTO transactions
(from_account, to_account, amount, status, flagged, created_at, updated_at)
VALUES
(?, ?, ?, ?, ?, ?, ?);
