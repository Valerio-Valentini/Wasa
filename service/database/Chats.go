package database

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func (db *appdbimpl) StartChat(group bool, members []string) (int64, error) {
	res, err := db.c.Exec("INSERT INTO chat (is_group) VALUES (?)", group)
	if err != nil {
		return -1, err
	}

	id_chat, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	id_chat_str := strconv.FormatInt(id_chat, 10)

	for _, member := range members {
		err = db.AddMember(id_chat_str, member)
		if err != nil {
			return -1, err
		}
	}

	if !group && len(members) == 2 {
		user1 := fmt.Sprintf("%v", members[0])
		user2 := fmt.Sprintf("%v", members[1])

		chatName := fmt.Sprintf("%s-%s", user1, user2)
		_, err = db.c.Exec("UPDATE chat SET chat_name = ? WHERE chat_id = ?", chatName, id_chat)
		if err != nil {
			return -1, err
		}
	}

	return id_chat, nil
}

func (db *appdbimpl) AddMember(chat_id string, user_id string) error {
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS (SELECT 1 FROM chat_members WHERE chat_id = ? AND user_id = ?)", chat_id, user_id).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return nil // Member already exists, so do nothing
	}

	_, err = db.c.Exec("INSERT INTO chat_members (chat_id, user_id) VALUES (?, ?)", chat_id, user_id)
	return err
}

func (db *appdbimpl) LeaveChat(chat_id string, user_id string) error {
	var memberCount int
	err := db.c.QueryRow("SELECT COUNT(*) FROM chat_members WHERE chat_id = ?", chat_id).Scan(&memberCount)
	if err != nil {
		return err
	}

	if memberCount == 2 {
		_, err = db.c.Exec("DELETE FROM chat_members WHERE chat_id = ?", chat_id)
		if err != nil {
			return err
		}
		_, err = db.c.Exec("DELETE FROM chat WHERE chat_id = ?", chat_id)
		if err != nil {
			return err
		}
	} else {
		_, err = db.c.Exec("DELETE FROM chat_members WHERE user_id = ? AND chat_id = ?", user_id, chat_id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *appdbimpl) GetChats(user_id string) ([]ChatPreview, error) {
	rows, err := db.c.Query("SELECT chat_id FROM chat_members WHERE user_id = ? ", user_id)
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var chats []ChatPreview
	for rows.Next() {
		var id int
		var chat ChatPreview
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		err = db.c.QueryRow("SELECT * FROM chat WHERE chat_id = ? ", id).Scan(&chat.Chat_id, &chat.Chat_group, &chat.Chat_photo, &chat.Chat_name)
		if err != nil {
			return nil, err
		}
		/*
			var users []string
			rows, err = db.c.Query("SELECT user_id FROM chat_members WHERE chat_id = ? ", id)
			if err != nil {
				return nil, err
			}
			for rows.Next() {
				var user string
				err = rows.Scan(&user)
				if err != nil {
					return nil, err
				}
				users = append(users, user)
			}
			// allUsers := strings.Join(users, "-")
			chat.Members = users
		*/
		members, err := db.GetMembers(chat.Chat_id)
		if err != nil {
			return nil, err
		}
		chat.Members = members

		err = db.c.QueryRow("SELECT owner, content, date FROM messages WHERE chat_id = ? ORDER BY date DESC LIMIT 1", id).Scan(&chat.Owner, &chat.Preview, &chat.Date)
		if err == sql.ErrNoRows {
			chat.Preview = ""
			chat.Owner = ""
			chat.Date = ""
		}
		if err != nil && err != sql.ErrNoRows {

			return nil, err
		}
		chats = append(chats, chat)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return chats, nil
}

func (db *appdbimpl) UpdateGroupPhoto(chat_id int, photo_id int) (int64, error) {

	_, err := db.c.Exec("DELETE FROM group_photo WHERE photo_id = ? AND chat_id = ?", photo_id, chat_id)
	if err != nil {
		return -1, err
	}

	res, err := db.c.Exec("INSERT INTO profile_photo (chat_id, photo_id) VALUES (?)", chat_id, photo_id)
	if err != nil {
		return -1, err
	}

	photo_id_db, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return photo_id_db, nil

}

func (db *appdbimpl) SetGroupName(user_id string, chat_id int, name string) error {

	res, err := db.VerifyUser(user_id)
	if err != nil || !res {
		return err
	}

	_, err = db.c.Exec("UPDATE chat SET chat_name = ? WHERE chat_id = ?", name, chat_id)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetMembers(chat_id int) (string, error) {
	var users []string
	rows, err := db.c.Query("SELECT user_id FROM chat_members WHERE chat_id = ? ", chat_id)
	if err != nil {
		return "", err
	}
	for rows.Next() {
		var user string
		err = rows.Scan(&user)
		if err != nil {
			return "", err
		}
		users = append(users, user)
	}
	allUsers := strings.Join(users, "-")

	return allUsers, nil
}
