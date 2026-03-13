package types

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age" validate:"required,gte=0,lte=150"`
	Grade string `json:"grade"`
}
