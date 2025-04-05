-- name: CreateOrder :exec
INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?);

-- name: ListOrders :many
SELECT * FROM orders;

