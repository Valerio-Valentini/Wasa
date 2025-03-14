package database

import (
	"errors"
)

func (db *appdbimpl) SendMessage(chat_id int, owner string, message Message) (int64, error) {
	res, err := db.c.Exec("INSERT INTO messages (chat_id, owner, content, forwarded, reply) VALUES (?, ?, ?, ?, ?)",
		chat_id, owner, message.Content, message.Forwarded, message.Reply)
	if err != nil {
		return -1, err
	}

	message_id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return message_id, nil
}

func (db *appdbimpl) DeleteMessage(owner string, chat_id string, message_id string) error {
	_, err := db.c.Exec("DELETE FROM messages WHERE owner = ? AND chat_id = ? AND message_id = ?", owner, chat_id, message_id)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) ForwardMessage(owner string, chat1_id string, msg_id string, chat2_id string) (int, error) {
	// Verify if user is a member of the original chat
	res1, err := db.VerifyUserIsMamberOfChat(owner, chat1_id)
	if err != nil {
		return -1, err
	}
	if !res1 {
		return -1, errors.New("User is not a member of the original chat")
	}

	// Verify if user is a member of the target chat
	res2, err := db.VerifyUserIsMamberOfChat(owner, chat2_id)
	if err != nil {
		return -1, err
	}
	if !res2 {
		return -1, errors.New("User is not a member of the target chat")
	}

	var content string
	err = db.c.QueryRow("SELECT content FROM messages WHERE message_id = ? AND chat_id = ?", msg_id, chat1_id).
		Scan(&content)
	if err != nil {
		return -1, err
	}

	result, err := db.c.Exec(`
		INSERT INTO messages (chat_id, owner, forwarded, content) 
		VALUES (?, ?, ?, ?)`,
		chat2_id, owner, 1, content)

	if err != nil {
		return -1, err
	}

	// Return the newly inserted message ID
	insertedID, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(insertedID), nil
}

func (db *appdbimpl) ReplyMessage(owner string, reply int, content string) error {
	var chat_id string
	err := db.c.QueryRow("SELECT chat_id FROM messages WHERE message_id = ? ", reply).Scan(&chat_id)
	if err != nil {
		return err
	}
	res, err := db.VerifyUserIsMamberOfChat(owner, chat_id)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("User Is Not A Member")
	}

	_, err = db.c.Exec("INSERT INTO messages (owner, reply, content) VALUES (?,?,?)", owner, reply, content)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetMessagesFromChat(chat_id string) ([]Message, error) {
	rows, err := db.c.Query("SELECT * FROM messages WHERE chat_id = ? ORDER BY date ASC", chat_id)

	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var messages []Message
	for rows.Next() {
		var message Message
		err = rows.Scan(
			&message.Message_id, &message.Chat_id,
			&message.Status, &message.Date, &message.Owner,
			&message.Forwarded, &message.Reply, &message.Media,
			&message.Content)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return messages, nil
}

func (db *appdbimpl) SendMedia(chat_id int, owner string, content string) (int64, error) {
	res, err := db.c.Exec("INSERT INTO messages (chat_id, owner, content) VALUES (?,?,?)", chat_id, owner, content)
	if err != nil {
		return -1, err
	}

	message_id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return message_id, nil
}

func (db *appdbimpl) DeleteMedia(owner string, photo_id int, chat_id int) error {
	_, err := db.c.Exec("DELETE FROM media_chat WHERE (owner = ?, chat_id = ?, photo_id = ?)", owner, chat_id, photo_id)
	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) UpdateMessageStatus(user_id string, chat_id string, message_id string) error {
	var presence int
	err := db.c.QueryRow("SELECT COUNT(*) FROM messages_status WHERE user_id = ? AND chat_id = ?", user_id, chat_id).Scan(&presence)
	if err != nil {
		return err
	}
	if(presence > 0) {
		_, err := db.c.Exec("UPDATE messages_status SET message_id = ? WHERE user_id = ? AND chat_id = ?", message_id, user_id, chat_id)
		if err != nil {
			return err
		}
	} else {
		_, err = db.c.Exec("INSERT INTO messages_status (user_id, chat_id, message_id) VALUES (?,?,?)", user_id, chat_id, message_id)
		if err != nil {
			return err
		}
	}
	var message string
	err = db.c.QueryRow("SELECT MIN(message_id) FROM messages_status WHERE chat_id = ?", chat_id).Scan(&message)
	if err != nil {
		return err
	}
	var membri_letto int
	var membri_tot int
	err = db.c.QueryRow("SELECT COUNT(*) FROM messages_status WHERE chat_id = ?", chat_id).Scan(&membri_letto)
	if err != nil {
		return err
	}
	err = db.c.QueryRow("SELECT COUNT(*) FROM chat_members WHERE chat_id = ?", chat_id).Scan(&membri_tot)
	if err != nil {
		return err
	}
	if(membri_letto == membri_tot) {
		_, err = db.c.Exec("UPDATE messages SET status = 1 WHERE message_id <= ?", message)
		if err != nil {
			return err
		}
	}
	return nil
}