package utils

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
    // Generate a hashed password with bcrypt and a cost of bcrypt.DefaultCost
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash compares a hashed password with a plain text password.
func CheckPasswordHash(password, hash string) bool {
    // Compare the hashed password with the plain text password
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
