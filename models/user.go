package models

var (
	UserList map[string]*User
)

type User struct {
	Id int
	Name string
	Age int
}

func init()  {
	//UserList = make(map[string]*User)
	//defaultUser := User{
	//	"user1",
	//	"Axell",
	//	20,
	//	}
	//UserList["user1"] = &defaultUser
}