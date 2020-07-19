package userdefinetype

//The Name should be capitalize as it is need to export
//the fields which needs to be asseible from outside the package must also be capotalize
type Worker struct {
	Id          int
	Password    string
	Name        string
	Deptartment string
	Salary      float64
	//basic way of defining the Address field
	//Adres Address

	//making the Address filed as anonymous field so that it can be accesss direclty
	// /anonymous fields offer much more than just the ability to skip providing a name for a field
	//in a struct definition. This Address struct is called as enonymus
	Address
}
type Address struct {
	Street   string
	Landmark string
	City     string
	State    string
	Country  string
	Pin      int
}
