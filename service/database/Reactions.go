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
	_, err = db.c.Exec("INSERT INTO message_reactions (owner, reaction, message_id) VALUES (?,?,?)", owner, reaction, message)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) DeleteReaction(owner string, message string, chat_id string) error {
	res, err := db.VerifyUserIsMamberOfChat(owner, chat_id)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("User Is Not A Member")
	}

	_, err = db.c.Exec("DELETE FROM message_reactions WHERE (owner = ? AND message = ?)", owner, message)
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
