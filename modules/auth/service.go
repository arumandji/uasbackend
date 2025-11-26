package auth

import (
    "context"
    "errors"
    "fmt"

    "prestasi-api/database"
    "prestasi-api/middleware"
    "prestasi-api/utils"

    "github.com/jmoiron/sqlx"
)

type AuthService struct {
    db *sqlx.DB
}

type User struct {
    ID           string `db:"id" json:"id"`
    Username     string `db:"username" json:"username"`
    PasswordHash string `db:"password_hash" json:"-"`
    RoleName     string `db:"role_name" json:"role"`
}

func NewAuthService() *AuthService {
    return &AuthService{db: database.DB}
}

func (s *AuthService) Authenticate(ctx context.Context, username, password string) (string, error) {
    var u User
    query := `SELECT u.id, u.username, u.password_hash, r.name as role_name
              FROM users u LEFT JOIN roles r ON u.role_id=r.id
              WHERE u.username=$1`
    if err := s.db.GetContext(ctx, &u, query, username); err != nil {
        return "", err
    }
    if !utils.CheckPasswordHash(password, u.PasswordHash) {
        return "", errors.New("invalid credentials")
    }
    token, err := middleware.GenerateToken(u.ID, u.RoleName)
    if err != nil {
        return "", fmt.Errorf("token gen: %w", err)
    }
    return token, nil
}
