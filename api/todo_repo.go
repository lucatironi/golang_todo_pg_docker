package main

import "gopkg.in/pg.v4"

type (
	TodoRepo struct {
		dbConnection *pg.DB
	}
)

func NewTodoRepo(db *pg.DB) *TodoRepo {
	return &TodoRepo{db}
}

func (tr TodoRepo) FindAllTodos() Todos {
	var todos Todos
	if err := tr.dbConnection.Model(&todos).Select(); err != nil {
		panic(err)
	}
	return todos
}

func (tr TodoRepo) FindTodo(todoId int) Todo {
	todo := Todo{Id: todoId}
	if err := tr.dbConnection.Select(&todo); err != nil {
		return Todo{}
	}
	return todo
}

func (tr TodoRepo) CreateTodo(todo Todo) Todo {
	if err := tr.dbConnection.Create(&todo); err != nil {
		panic(err)
	}
	return todo
}

func (tr TodoRepo) UpdateTodo(todo Todo) Todo {
	if err := tr.dbConnection.Update(&todo); err != nil {
		panic(err)
	}
	return todo
}

func (tr TodoRepo) DeleteTodo(todo Todo) Todo {
	if err := tr.dbConnection.Delete(&todo); err != nil {
		panic(err)
	}
	return todo
}
