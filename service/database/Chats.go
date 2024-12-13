package database

func (db *appdbimpl) StartChat(group bool, members []string) (int, error)
{
	res, err:= db.c.Exec("INSERT INTO chat (group) VALUES (?)", group)
	if err != nil
	{	
		return -1, err
	}

	id_chat, err := res.LastInsertId()
	if err != nil
	{	
		return -1, err
	}
	
	for _,member := range members
	{
		_, err:= AddMember(id_chat,member)
		if err != nil
		{	
			return -1, err
		}
	}

	return id_chat, nil
}

func (db *appdbimpl) AddMember(chat_id int, user_id string) error
{
	_, err:= db.c.Exec("INSERT INTO chat_members (chat_id, user_id) VALUES (?,?)", chat_id, user_id)
	if err != nil
	{	
		return err
	}

	return nil
}

func (db *appdbimpl) LeaveChat(chat_id int, user_id string) error
{
	_, err:= db.c.Exec("DELETE FROM chat_members WHERE (user_id = ? AND chat_id = ?)", user_id, chat_id)
	if err != nil
	{	
		return err
	}

	return nil
}

