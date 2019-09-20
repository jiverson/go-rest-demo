package db

import (
	"database/sql"
	"log"

	"github.com/jiverson/go-rest-demo/model"
)

func (db *Database) GetTodoById(id uint) (*model.Todo, error) {
	var todo model.Todo

	sqlStatement := `
		SELECT *
		FROM todos
		WHERE (id = $1)
		AND deleted_at IS NULL
		ORDER BY id
		ASC LIMIT 1`
	if err := db.QueryRow(sqlStatement, id).Scan(&todo.ID, &todo.Name, &todo.Done, &todo.UserID, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &todo, nil
}

func (db *Database) GetTodosByUserId(userId uint) ([]*model.Todo, error) {
	var todos []*model.Todo

	sqlStatement := `
		SELECT id, name, done
		FROM todos
		WHERE (user_id = $1)
		AND deleted_at IS NULL`
	rows, err := db.Query(sqlStatement, userId)

	if err != nil {
		return nil, err
	}
	//defer rows.Close()
	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Done)
		if err != nil {
			// TODO return error??
			log.Fatal(err)
		}
		todos = append(todos, &todo)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// return empty array
	if len(todos) == 0 {
		todos = []*model.Todo{}
	}

	return todos, nil
}

func (db *Database) CreateTodo(todo *model.Todo) error {
	sqlStatement := `
		INSERT INTO todos (name, done, user_id, created_at, updated_at)
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
		RETURNING id`
	err := db.QueryRow(sqlStatement, todo.Name, todo.Done, todo.UserID).Scan(&todo.ID)

	return err
}

func (db *Database) UpdateTodo(todo *model.Todo) error {
	sqlStatement := `
		UPDATE todos
		SET name = $1, done = $2, updated_at = CURRENT_TIMESTAMP
		WHERE id = $3
		AND deleted_at IS NULL`

	res, err := db.Exec(sqlStatement, todo.Name, todo.Done, todo.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteTodoById(id uint) error {
	sqlStatement := `
		UPDATE todos
		SET deleted_at = CURRENT_TIMESTAMP
		WHERE id = $1
		AND deleted_at IS NULL`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// Unused right now
func (db *Database) HardDeleteTodoById(id uint) error {
	sqlStatement := `DELETE from todos where id = $1`

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
