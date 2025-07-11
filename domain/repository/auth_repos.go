package repository

// IReposAUTH AUTH仓储接口
type IReposAUTH interface {
    GenerateToken(userID string, role string) (string, error)
    ValidateToken(tokenStr string) (string, string, error) // returns userID, role
}
