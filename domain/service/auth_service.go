package service

type JWTService interface {
    GenerateToken(userID string, role string) (string, error)
    ValidateToken(tokenStr string) (string, string, error) // returns userID, role
}
