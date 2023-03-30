package users

import (
	"database/sql"
	"log"

	"github.com/brianlangdon/tada-auth/graph/model"
	database "github.com/brianlangdon/tada-auth/internal/pkg/db/mysql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (user *User) Create() bool {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Password, Email, Title, Location, Gender, Picture, Company, FaveLangauge, FaveFramework, FaveTool) VALUES(?,?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		return false
	}
	defer statement.Close()

	hashedPassword, _ := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword, user.Email, "", "", "OTHER", "", "", "", "", "")

	return err == nil
}

func (user *User) Authenticate() bool {
	statement, err := database.Db.Prepare("select Password from Users WHERE Email = ?")

	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	row := statement.QueryRow(user.Email)

	var hashedPassword string
	err = row.Scan(&hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}

	return CheckPasswordHash(user.Password, hashedPassword)
}

// GetUserIdByUsername check if a user exists in database by given username
func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("select ID from Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	row := statement.QueryRow(username)

	var Id int
	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}

	return Id, nil
}

// GetUserByID check if a user exists in database and return the user object.
func GetUsernameById(userId string) (User, error) {
	statement, err := database.Db.Prepare("select Username from Users WHERE ID = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()

	row := statement.QueryRow(userId)

	var username string
	err = row.Scan(&username)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return User{}, err
	}

	return User{ID: userId, Username: username}, nil
}

// GetUserByID check if a user exists in database and return the user object.
func GetUserByEmail(email string) (model.FullUser, error) {
	statement, err := database.Db.Prepare("select Username, Email, Title, Location, Gender, Picture, Company, FaveLangauge, FaveFramework, FaveTool from Users WHERE Email = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer statement.Close()
	row := statement.QueryRow(email)

	var user model.FullUser
	err = row.Scan(&user.Username, &user.Email, &user.Title, &user.Location, &user.Gender, &user.Picture, &user.Company, &user.FaveLangauge, &user.FaveFramework, &user.FaveTool)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return user, err
	}

	return user, nil
}

// HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
