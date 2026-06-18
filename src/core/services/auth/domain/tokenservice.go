package domain

type TokenService interface {
    GenerateToken(userID int32, email string, usertype int, approved bool) (Token, error)
    ValidateToken(tokenString string) (int32, error) 
}
