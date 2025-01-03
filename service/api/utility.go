package api

import (
	"database/sql"
)



func (rt *_router) VerifyUser (username string) (bool, error) {
	var presence int
	err:= rt.db.c.QueryRow("SELECT 1 FROM users WHERE user_id = ? LIMIT 1 ", username).Scan(&presence)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		
		return false, err
	}

	return true, nil
}

func (rt *_router) VerifyUserIsMamberOfChat (user_id string, chat_id int) (bool, error) {
	var presence int
	err:= rt.db.c.QueryRow("SELECT 1 FROM users WHERE user_id = ? AND chat_id = ? LIMIT 1 ", user_id, chat_id).Scan(&presence)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		
		return false, err
	}

	return true, nil
}

func (rt *_router) GetIdProfilePicture (user_id string) (int, error) {
	var id int
	err:= rt.db.c.QueryRow("SELECT photo_id FROM profile_photo WHERE user_id = ?", user_id).Scan(&id)
	if err != nil {
		return -1, err
	}
	return -1, nil
}

func (rt *_router) CreateNewId (user_id string) (int, error) {
	err:= rt.db.c.Exec("DELETE FROM profile_photo WHERE user_id = ?", user_id)
	if err != nil {
		return -1, err
	}
	result, err:= rt.db.c.Exec("INSERT INTO profile_photo (user_id) VALUES(?)", user_id)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) CreateNewPhotoId (chat_id string) (int, error) {
	err:= rt.db.c.Exec("DELETE FROM group_photo WHERE chat_id = ?", chat_id)
	if err != nil {
		return nil, err
	}
	result, err:= rt.db.c.Exec("INSERT INTO group_photo (user_id) VALUES(?)", chat_id)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) GetIdGroupPicture (chat_id string) (int, error) {
	var id int
	err:= rt.db.c.QueryRow("SELECT photo_id FROM group_photo WHERE chat_id = ?", chat_id).Scan(&id)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (rt *_router) GetIdPhoto (user_id string) (int, error) {
	var id int
	err:= rt.db.c.QueryRow("SELECT photo_id FROM media_chat WHERE user_id = ?", user_id).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) CreateNewMediaId (user_id string) (int, error) {
	err:= rt.db.c.Exec("DELETE FROM media_chat WHERE user_id = ?", user_id)
	if err != nil {
		return -1, err
	}
	result, err:= rt.db.c.Exec("INSERT INTO media_chat (user_id) VALUES(?)", user_id)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return id, nil
}