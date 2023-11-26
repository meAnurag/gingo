package request

type CreateTaskRequest struct {
	Id          int    `validate:"required,min=1,max=200" json:"id" binding:"required"`
	Name        string `validate:"required" json:"name" binding:"required"`
	Description string `validate:"required" json:"description" binding:"required"`
}

type FindAllTaskRequest struct{}

type UpdateTaskRequest struct {
	Id          int    `validate:"required,min=1,max=200" json:"id" binding:"required"`
	Name        string `validate:"required" json:"name" binding:"required"`
	Description string `validate:"required" json:"description" binding:"required"`
}
