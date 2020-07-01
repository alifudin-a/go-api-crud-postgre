package models

type Employee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}
