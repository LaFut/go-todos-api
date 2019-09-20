package description

import (
	"todos-rest/model/entity"
	"todos-rest/model/repository"
)

// Todo data
// swagger:response TodoResponse
type TodoResponse struct {
	// in: body
	Body struct {
		Entity *entity.Todo
		*BaseResponseBody
	}
}

// Todos data
// swagger:response TodoResponse
type TodosResponse struct {
	// in: body
	Body struct {
		Entities []*entity.Todo
		*BaseResponseBody
	}
}

// swagger:parameters CreateTodo UpdateTodo
type TodoParameters struct {
	// in: body
	Body *entity.TodoFields
}

// swagger:parameters GetTodo
type TodosParameters struct {
	// in: query
	*repository.TodoParameters
}
