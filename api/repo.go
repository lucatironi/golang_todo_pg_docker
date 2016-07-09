package main

func RepoFindAllTodos() []Todo {
  var todos []Todo
  err := db.Model(&todos).Select()
  if err != nil {
    panic(err)
  }

  return todos
}

func RepoFindTodo(todoId int) Todo {
  todo := Todo{Id: todoId}
  err := db.Select(&todo)
  if err != nil {
    panic(err)
  }

  return todo
}

func RepoCreateTodo(newTodo Todo) Todo {
  err := db.Create(&newTodo)
  if err != nil {
    panic(err)
  }

  return newTodo
}
