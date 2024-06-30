package repositories

import (
    "database/sql"
    "golang-crud/models"
)

func CreateUser(db *sql.DB, user models.User) (int64, error) {
    result, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
    if err != nil {
        return 0, err
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func GetUser(db *sql.DB, id int) (models.User, error) {
    var user models.User
    err := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email)
    if err != nil {
        return user, err
    }

    return user, nil
}

func GetUserByEmail(db *sql.DB, email string) (models.User, error) {
    var user models.User
    err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
    if err != nil {
        return user, err
    }

    return user, nil
}

func GetUsers(db *sql.DB) ([]models.User, error) {
    rows, err := db.Query("SELECT id, name, email FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    users := []models.User{}
    for rows.Next() {
        var user models.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func UpdateUser(db *sql.DB, id int, user models.User) error {
    _, err := db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, id)
    if err != nil {
        return err
    }

    return nil
}

func DeleteUser(db *sql.DB, id int) error {
    _, err := db.Exec("DELETE FROM users WHERE id = ?", id)
    if err != nil {
        return err
    }

    return nil
}
