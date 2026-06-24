package dto

type CreateUserRequest struct {
    Name        string `json:"name"`
    Lastname    string `json:"lastname"`
    Email       string `json:"email"`
    NumberPhone string `json:"numberphone"`
    Birthdate   string `json:"birthdate"`
    Password    string `json:"password"`
    UserType    int    `json:"user_type"`
}

type UserResponse struct {
    IdUser      int32  `json:"iduser"`
    Name        string `json:"name"`
    Lastname    string `json:"lastname"`
    Email       string `json:"email"`
    NumberPhone string `json:"numberphone"`
    Birthdate   string `json:"birthdate"`
    UserType    int    `json:"usertype"`
    ImageURL    string `json:"image_url"`
}

type UpdateUserRequest struct {
    Name        string `json:"name,omitempty"`
    Lastname    string `json:"lastname,omitempty"`
    Email       string `json:"email,omitempty"`
    NumberPhone string `json:"numberphone,omitempty"`
    Birthdate   string `json:"birthdate,omitempty"`
}