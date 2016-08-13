package models


type Designer struct {
	Id int
	Name string
	Email string
	Created_at string
	Updated_at string
}


func (d Designer) GetName() string {
	return d.Name
}