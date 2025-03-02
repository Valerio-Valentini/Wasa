package database

import (
	"errors"
)

func (db *appdbimpl) InsertReaction(owner string, reaction string, message string, chat_id string) error {
	res, err := db.VerifyUserIsMamberOfChat(owner, chat_id)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("User Is Not A Member")
	}

	_, err = db.c.Exec("INSERT OR IGNORE INTO message_reactions (owner, reaction, message_id) VALUES (?, ?, ?)", owner, reaction, message)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetMessageReactions(message_id string) ([]Reaction, error) {
	rows, err := db.c.Query("SELECT owner, reaction, message_id FROM message_reactions WHERE message_id = ?", message_id)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	var reactions []Reaction
	for rows.Next() {
		var reaction Reaction
		err = rows.Scan(&reaction.Owner, &reaction.Reaction, &reaction.Messageid)
		if err != nil {
			return nil, err
		}
		reactions = append(reactions, reaction)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return reactions, nil
}

func (db *appdbimpl) DeleteReaction(owner string, message string, chat_id string) error {
	res, err := db.VerifyUserIsMamberOfChat(owner, chat_id)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("User Is Not A Member")
	}

	_, err = db.c.Exec("DELETE FROM message_reactions WHERE owner = ? AND message_id = ?", owner, message)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) ChangeReaction(owner string, reaction string, message string, chat_id string) error {
	res, err := db.VerifyUserIsMamberOfChat(owner, chat_id)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("User Is Not A Member")
	}

	_, err = db.c.Exec("UPDATE message_reactions SET reaction = ? WHERE (owner = ? AND message = ?)", reaction, owner, message)
	if err != nil {
		return err
	}

	return nil
}
