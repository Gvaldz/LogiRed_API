package domain

type TokenService interface {
    GenerateToken(userID int32, email string, usertype int, citywork string) (Token, error)
    ValidateToken(tokenString string) (int32, error) 
}
