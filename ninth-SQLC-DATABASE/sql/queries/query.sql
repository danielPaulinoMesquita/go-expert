-- name: ListCategories :many
SELECT * FROM categories;

-- name: GetCategory :one
SELECT * FROM categories
where id = ?;

-- name: CreateCategory :exec
INSERT INTO categories (id, name, description)
values (?,?,?);

-- name: UpdateCategory :exec
UPDATE categories SET name = ?, description = ?
WHERE id = ?;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = ?;