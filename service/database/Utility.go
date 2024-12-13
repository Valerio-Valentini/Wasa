package database

func (db *appdbimpl) VerifyUser (username string) (bool, error)
{
	var presence int
	err:= db.c.QueryRow("SELECT 1 FROM users WHERE user_id = ? LIMIT 1 ", username).Scan(&presence)
	if err != nil
	{
		if err == sql.ErrNoRows
		{
			return false, nil
		}
		
		return false, err
	}

	return true, nil
}

func (db *appdbimpl) VerifyUserIsMamberOfChat (user_id string, chat_id int) (bool, error)
{
	var presence int
	err:= db.c.QueryRow("SELECT 1 FROM users WHERE user_id = ? AND chat_id = ? LIMIT 1 ", user_id, chat_id).Scan(&presence)
	if err != nil
	{
		if err == sql.ErrNoRows
		{
			return false, nil
		}
		
		return false, err
	}

	return true, nil
}