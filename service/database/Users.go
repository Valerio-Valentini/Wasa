package database

func (db *appdbimpl) InsertUser(username string) error {
	_, err := db.c.Exec("INSERT INTO users (user_id, photo_id) VALUES (?, 0)", username)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) ChangePhoto(user_id string, photo_id int) (int64, error) {
	_, err := db.c.Exec("DELETE FROM profile_photo WHERE photo_id = ? AND user_id = ?", photo_id, user_id)
	if err != nil {
		return -1, err
	}

	res, err := db.c.Exec("INSERT INTO profile_photo (user_id, photo_id) VALUES (?)", user_id, photo_id)
	if err != nil {
		return -1, err
	}

	photo_id_db, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return photo_id_db, nil
}

func (db *appdbimpl) UpdateUser(user_id string, new_user_id string) error {

	res, err := db.VerifyUser(new_user_id)
	if err != nil || res {
		return err
	}

	_, err = db.c.Exec("UPDATE users SET user_id = ? WHERE user_id = ?", new_user_id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) SearchUser(user_id string) ([]User, error) {
	rows, err := db.c.Query("SELECT * FROM users WHERE user_id LIKE ? ", user_id+"%") // QUI

	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()

	var users []User
	for rows.Next() {
		var user User
		err = rows.Scan(&user.User_id, &user.Photo_id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return users, nil
}
