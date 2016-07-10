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
    return Todo{}
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

func RepoUpdateTodo(updatedTodo Todo) Todo {
  err := db.Update(&updatedTodo)
  if err != nil {
    panic(err)
  }

  return updatedTodo
}

func RepoDeleteTodo(deletedTodo Todo) Todo {
  err := db.Delete(&deletedTodo)
  if err != nil {
    panic(err)
  }

  return deletedTodo
}
