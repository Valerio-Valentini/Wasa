package api

import (
	"database/sql"
)



func (rt *_router) VerifyUser (username string) (bool, error) {
	res, err:= rt.db.VerifyUser(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		
		return false, err
	}

	return res, nil
}

func (rt *_router) VerifyUserIsMamberOfChat (user_id string, chat_id int) (bool, error) {
	res, err:= rt.db.VerifyUserIsMamberOfChat (user_id, chat_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		
		return false, err
	}

	return res, nil
}

func (rt *_router) GetIdProfilePicture (user_id string) (int, error) {
	id, err:= rt.db.GetIdProfilePicture (user_id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) CreateNewId (user_id string) (int, error) {
	id, err:= rt.db.CreateNewId (user_id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) CreateNewPhotoId (chat_id string) (int, error) {
	id, err:= rt.db.CreateNewPhotoId (chat_id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) GetIdGroupPicture (chat_id string) (int, error) {
	id, err:= rt.db.GetIdGroupPicture (chat_id)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (rt *_router) GetIdPhoto (user_id string) (int, error) {
	id, err:= rt.db.GetIdPhoto (user_id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (rt *_router) CreateNewMediaId (user_id string) (int, error) {
	id, err:= rt.db.CreateNewMediaId (user_id)
	if err != nil {
		return -1, err
	}
	return id, nil
}