package entities

type User struct {
	IdUser      int32  `json:"id"`        
    Name        string `json:"name"`      
    Lastname    string `json:"lastname"`  
    Email       string `json:"email"`     
    NumberPhone string `json:"numberphone"`
    Password    string `json:"password"`
    UserType    int    `json:"user_type"`
    Birthdate   string `json:"birthday"`
    ImageURL    string `json:"image_url"`
}