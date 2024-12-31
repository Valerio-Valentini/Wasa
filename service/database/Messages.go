package database

import(
	"errors"
)

func (db *appdbimpl) SendMessage(chat_id int, owner string, content string) (int64, error) {
	res, err:= db.c.Exec("INSERT INTO messages (chat_id, owner, content) VALUES (?,?,?)", chat_id, owner, content)
	if err != nil {	
		return -1, err
	}

	message_id, err := res.LastInsertId()
	if err != nil {	
		return -1, err
	}

	return message_id, nil
}

func (db *appdbimpl) DeleteMessage(owner string, chat_id int, message_id int) error {
	_, err:= db.c.Exec("DELETE FROM messages WHERE (owner = ?, chat_id = ?, message_id = ?)", owner, chat_id, message_id)
	if err != nil {	
		return err
	}

	return nil
}

func (db *appdbimpl) ForwardMessage(owner string, chat1_id int, content string, chat2_id int) (int, error) {
	res1, err := db.VerifyUserIsMamberOfChat(owner, chat1_id)
	if err != nil {	
		return -1,err
	}
	if !res1{
		return -1, nil
	}
	res2, err := db.VerifyUserIsMamberOfChat(owner, chat2_id)
	if err != nil {	
		return -1, err
	}
	if !res2{
		return -1, nil
	}

	_, err = db.c.Exec("INSERT INTO messages (owner, content, chat_id, forwarded) VALUES (?, ?,?,?)", owner, content, chat2_id, true)
	if err != nil {	
		return -1, err
	}

	return 1, nil
}

func (db *appdbimpl) ReplyMessage(owner string, reply int, content string) error {
	chat_id, err:= db.c.QueryRow("SELECT chat_id FROM messages WHERE message_id = ? ", reply)
	if err != nil {	
		return err
	}
	res, err := db.VerifyUserIsMamberOfChat(owner, chat_id)
	if err != nil {	
		return err
	}
	if !res{
		return errors.New("User Is Not A Member")
	}

	_, err:= db.c.Exec("INSERT INTO messages (owner, reply, content) VALUES (?,?,?)", owner, reply, content)
	if err != nil {	
		return err
	}

	return nil
}

func (db *appdbimpl) GetMessagesFromChat(chat_id int) ([]Message, error) {
	rows, err:= db.c.QueryRow("SELECT * FROM messages WHERE chat_id = ? ", chat_id)
	
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var messages []Message
	for rows.Next() {
		var message Message
		err = rows.Scan(&message)
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

func (db *appdbimpl) SendMedia(chat_id int, owner string, content string) (int, error) {
	_, err:= db.c.Exec("INSERT INTO messages (chat_id, owner, content) VALUES (?,?,?)", chat_id, owner, content)
	if err != nil {	
		return -1, err
	}

	message_id, err := res.LastInsertId()
	if err != nil {	
		return -1, err
	}

	return media_id, nil
}

func (db *appdbimpl) DeleteMedia(owner string, photo_id int, chat_id int) error {
	_, err:= db.c.Exec("DELETE FROM media_chat WHERE (owner = ?, chat_id = ?, photo_id = ?)", owner, chat_id, photo_id)
	if err != nil {	
		return err
	}

	return nil
}