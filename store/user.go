package store

import "new-aspect/MindStream/common"

type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	WxOpenId   string `json:"wxOpenId"`
	GithubName string `json:"githubName"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

func GetUserById(id string) (User, error) {
	query := `SELECT id, username, password, wx_open_id, github_name, created_at, updated_at FROM users WHERE id=?`
	var user User
	err := DB.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Password, &user.WxOpenId, &user.GithubName, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func CreateNewUser(username, password, githubName, wxOpenId string) (User, error) {
	nowDateTimeStr := common.GetNowDateTimeStr()
	newUser := User{
		Id:         common.GetUUID(),
		Username:   username,
		Password:   password,
		WxOpenId:   wxOpenId,
		GithubName: githubName,
		CreatedAt:  nowDateTimeStr,
		UpdatedAt:  nowDateTimeStr,
	}

	query := `INSERT INTO users (id, username, password, wx_open_id, github_name, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`
	_, err := DB.Exec(query, newUser.Id, newUser.Username, newUser.Password, newUser.WxOpenId, newUser.GithubName, newUser.CreatedAt, newUser.UpdatedAt)

	return newUser, err
}

func GetUserByUsernameAndPassword(username, password string) (User, error) {
	query := `SELECT id, username, password, wx_open_id, github_name,created_at, updated_at FROM users WHERE username=? AND password=?`
	var user User
	err := DB.QueryRow(query, username, password).Scan(&user.Id, &user.Username, &user.Password, &user.WxOpenId, &user.GithubName, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}
