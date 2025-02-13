/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// ----------------------------------------------------- Chats_Functions
	StartChat(group bool, members []string) (int64, error)       // ok
	AddMember(chat_id string, user_id string) error              // ok
	LeaveChat(chat_id string, user_id string) error              // ok
	GetChats(user_id string) ([]Chat, error)                     // ok
	UpdateGroupPhoto(chat_id int, photo_id int) (int64, error)   // ok
	SetGroupName(user_id string, chat_id int, name string) error // ok
	// ----------------------------------------------------- Messages_Functions
	GetMessagesFromChat(chat_id string) ([]Message, error)                                              // ok
	SendMedia(chat_id int, owner string, content string) (int64, error)                                 // ok
	SendMessage(chat_id int, owner string, message Message) (int64, error)                              // ok
	DeleteMessage(owner string, chat_id string, message_id string) error                                // ok
	ForwardMessage(owner string, chat1_id string, original_msg_id string, chat2_id string) (int, error) // ok
	ReplyMessage(owner string, reply int, content string) error                                         // ok
	DeleteMedia(user_id string, photo_id int, chat_id int) error                                        // ok
	// ----------------------------------------------------- Reactions_Functions
	GetMessageReactions(message_id string) ([]Reaction, error)
	ChangeReaction(owner string, reaction string, message string, chat_id string) error // ok
	DeleteReaction(owner string, message string, chat_id string) error                  // ok
	InsertReaction(owner string, reaction string, message string, chat_id string) error // ok
	// ----------------------------------------------------- Users_Functions
	ChangePhoto(user_id string, photo_id int) (int64, error)  // ok
	SearchUser(user_id string, sender string) ([]User, error) // ok
	UpdateUser(user_id string, new_user_id string) error      // ok
	InsertUser(username string) error                         // ok
	VerifyUser(username string) (bool, error)                 // ok
	VerifyUserIsMamberOfChat(user_id string, chat_id string) (bool, error)
	GetIdProfilePicture(user_id string) (int64, error)
	CreateNewId(user_id string) (int64, error)
	CreateNewPhotoId(chat_id string) (int64, error)
	GetIdPhoto(user_id string) (int, error)
	GetIdGroupPicture(chat_id string) (int64, error)
	CreateNewMediaId(user_id string) (int64, error)
	NewChat(group bool, photo int, name string) (int64, error) // <--
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	var tables = [9]string{
		`CREATE TABLE IF NOT EXISTS media_chat (
			photo_id INTEGER PRIMARY KEY AUTOINCREMENT,
			owner VARCHAR(16) NOT NULL,
			chat_id INTEGER NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS profile_photo (
			photo_id INTEGER PRIMARY KEY AUTOINCREMENT,
			owner VARCHAR(16) NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS group_photo (
			photo_id INTEGER PRIMARY KEY AUTOINCREMENT,
			chat_id INTEGER(16) NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS users (
			user_id VARCHAR(16) NOT NULL PRIMARY KEY,
			photo_id INTEGER NOT NULL DEFAULT 0
		);`,

		`CREATE TABLE IF NOT EXISTS chat (
			chat_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			is_group BOOLEAN NOT NULL,
			chat_photo INTEGER NOT NULL DEFAULT 0,
			chat_name VARCHAR(20) DEFAULT 'A chat'
		);`,

		`CREATE TABLE IF NOT EXISTS chat_members (
			user_id VARCHAR(16) NOT NULL,
			chat_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (chat_id) REFERENCES chat(chat_id) ON DELETE CASCADE,
			PRIMARY KEY (user_id, chat_id)
		);`,

		`CREATE TABLE IF NOT EXISTS messages (
			message_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			chat_id INTEGER NOT NULL,
			status INTEGER NOT NULL DEFAULT 0,
			date DATETIME DEFAULT CURRENT_TIMESTAMP,
			owner VARCHAR(16) NOT NULL,
			forwarded INTEGER NOT NULL DEFAULT 0,
			reply INTEGER NOT NULL DEFAULT -1,
			media INTEGER NOT NULL DEFAULT -1,
			content VARCHAR(16),
			FOREIGN KEY (chat_id) REFERENCES chat(chat_id) ON DELETE CASCADE,
			FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (reply) REFERENCES messages(message_id) ON DELETE SET NULL
		);`,

		`CREATE TABLE IF NOT EXISTS message_reactions (
			owner VARCHAR(16) NOT NULL,
			reaction VARCHAR(16) DEFAULT '-1',
			message_id INTEGER NOT NULL,
			FOREIGN KEY (owner) REFERENCES users(user_id) ON DELETE CASCADE,
			FOREIGN KEY (message_id) REFERENCES messages(message_id) ON DELETE CASCADE,
			PRIMARY KEY (owner, message_id)
		);`,
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	for i := 0; i < len(tables); i++ {
		_, err := db.Exec(tables[i])
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
