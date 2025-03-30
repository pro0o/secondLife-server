package data

// import (
// 	"context"
// 	"errors"
// 	"time"

// 	"github.com/pro0o/second-life/gen/bowie/public/model"
// 	"github.com/pro0o/second-life/gen/bowie/public/table"

// 	"github.com/go-jet/jet/v2/postgres"
// 	"github.com/golang-jwt/jwt/v5"
// 	"github.com/google/uuid"
// 	"golang.org/x/crypto/bcrypt"
// )

// // AuthRepository defines the interface for authentication operations
// type AuthRepository interface {
// 	SignUp(ctx context.Context, email string, password string, username string) (*model.Users, string, string, error)
// 	Login(ctx context.Context, email string, password string) (string, string, error)
// }

// // auth_repository_impl implements the AuthRepository interface
// type auth_repository_impl struct {
// 	DB           postgres.DB
// 	jwtSecret    []byte
// 	tokenExpiry  time.Duration
// 	refreshToken time.Duration
// }

// // NewAuthRepository creates a new instance of AuthRepository
// func NewAuthRepository(db postgres.DB, jwtSecret string, tokenExpiry, refreshToken time.Duration) AuthRepository {
// 	return &auth_repository_impl{
// 		DB:           db,
// 		jwtSecret:    []byte(jwtSecret),
// 		tokenExpiry:  tokenExpiry,
// 		refreshToken: refreshToken,
// 	}
// }

// // Additional errors for authentication
// var (
// 	ErrInvalidCredentials = errors.New("invalid credentials")
// 	ErrUserExists         = errors.New("user already exists")
// )

// // Password table definition (we'll assume this exists based on the auth proto)
// type PasswordRecord struct {
// 	UserID   uuid.UUID
// 	Password string
// }

// func (repo *auth_repository_impl) SignUp(ctx context.Context, email string, password string, username string) (*model.Users, string, string, error) {
// 	// Check if user already exists
// 	existingUser, err := repo.getUserByEmail(ctx, email)
// 	if err != nil && err != ErrNotFound {
// 		return nil, "", "", err
// 	}
// 	if existingUser != nil {
// 		return nil, "", "", ErrUserExists
// 	}

// 	// Create new user
// 	id := uuid.New()
// 	profilePicture := int32(0) // Default profile picture
// 	user := &model.Users{
// 		ID:             id,
// 		Email:          &email,
// 		ProfilePicture: &profilePicture,
// 		Username:       &username,
// 	}

// 	// Begin transaction
// 	tx, err := repo.DB.(postgres.DBTx).Begin()
// 	if err != nil {
// 		return nil, "", "", err
// 	}
// 	defer tx.Rollback()

// 	// Insert user
// 	err = table.Users.
// 		INSERT(table.Users.AllColumns).
// 		MODEL(user).
// 		RETURNING(table.Users.AllColumns).
// 		QueryContext(ctx, tx, user)
// 	if err != nil {
// 		return nil, "", "", err
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, "", "", err
// 	}

// 	// Insert password record (assuming table.Passwords exists)
// 	// Note: This would need adjustment based on your actual schema
// 	_, err = table.Passwords.
// 		INSERT(table.Passwords.UserID, table.Passwords.Password).
// 		VALUES(postgres.UUID(id), postgres.String(string(hashedPassword))).
// 		ExecContext(ctx, tx)
// 	if err != nil {
// 		return nil, "", "", err
// 	}

// 	// Generate tokens
// 	accessToken, refreshToken, err := repo.generateTokens(user.ID)
// 	if err != nil {
// 		return nil, "", "", err
// 	}

// 	// Commit transaction
// 	if err = tx.Commit(); err != nil {
// 		return nil, "", "", err
// 	}

// 	return user, accessToken, refreshToken, nil
// }

// func (repo *auth_repository_impl) Login(ctx context.Context, email string, password string) (string, string, error) {
// 	// Get user by email
// 	user, err := repo.getUserByEmail(ctx, email)
// 	if err != nil {
// 		if err == ErrNotFound {
// 			return "", "", ErrInvalidCredentials
// 		}
// 		return "", "", err
// 	}

// 	// Get password record
// 	var passwordRecord PasswordRecord
// 	err = table.Passwords.
// 		SELECT(table.Passwords.Password).
// 		WHERE(table.Passwords.UserID.EQ(postgres.UUID(user.ID))).
// 		QueryContext(ctx, repo.DB, &passwordRecord)
// 	if err != nil {
// 		if isNotFoundError(err) {
// 			return "", "", ErrInvalidCredentials
// 		}
// 		return "", "", err
// 	}

// 	// Verify password
// 	err = bcrypt.CompareHashAndPassword([]byte(passwordRecord.Password), []byte(password))
// 	if err != nil {
// 		return "", "", ErrInvalidCredentials
// 	}

// 	// Generate tokens
// 	accessToken, refreshToken, err := repo.generateTokens(user.ID)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return accessToken, refreshToken, nil
// }

// // Helper function to get user by email
// func (repo *auth_repository_impl) getUserByEmail(ctx context.Context, email string) (*model.Users, error) {
// 	user := &model.Users{}

// 	err := table.Users.
// 		SELECT(table.Users.AllColumns).
// 		WHERE(table.Users.Email.EQ(postgres.String(email))).
// 		QueryContext(ctx, repo.DB, user)

// 	if err != nil {
// 		if isNotFoundError(err) {
// 			return nil, ErrNotFound
// 		}
// 		return nil, err
// 	}

// 	return user, nil
// }

// // Generate JWT tokens
// func (repo *auth_repository_impl) generateTokens(userID uuid.UUID) (string, string, error) {
// 	// Create access token
// 	accessTokenClaims := jwt.MapClaims{
// 		"sub": userID.String(),
// 		"exp": time.Now().Add(repo.tokenExpiry).Unix(),
// 	}
// 	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
// 	accessTokenString, err := accessToken.SignedString(repo.jwtSecret)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	// Create refresh token
// 	refreshTokenClaims := jwt.MapClaims{
// 		"sub": userID.String(),
// 		"exp": time.Now().Add(repo.refreshToken).Unix(),
// 	}
// 	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
// 	refreshTokenString, err := refreshToken.SignedString(repo.jwtSecret)
// 	if err != nil {
// 		return "", "", err
// 	}

// 	return accessTokenString, refreshTokenString, nil
// }
