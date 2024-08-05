
在store包里面，先搭建一个SQLite，先写一个初始化的InitDBConn(),然后


下面是sqlite.go文件
```go
var DB *sql.DB

func InitDBConn() {
	db, err := sql.Open("sqlite3", "./resources/memos.db")
	if err != nil {
		fmt.Println("connect failed")
	} else {
		DB = db
		fmt.Println("connect to sqlite success")
	}
}

```

下面是user.go文件
```go
type User struct {
	Id         string `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	WxOpenId   string `json:"wxOpenId"`
	GithubName string `json:"githubName"`
	CreatedAt  string `json:"createdAt"`
	UpdateAt   string `json:"updateAt"`
}

func GetUserById(id string) (User, error) {
	query := `SELECT id, username, password, wx_open_id, github_name, create_at, update_at FROM users WHERE id=?`
	var user User
	err := DB.QueryRow(query, id).Scan(&user.Id, &user.Username, &user.Password, &user.WxOpenId, &user.GithubName, &user.CreatedAt, &user.UpdateAt)
	return user, err
}
```

在执行上面的语句报错，意识到需要导入SQL的表语句

```
CREATE TABLE `users` (
    `id` TEXT NOT NULL PRIMARY KEY,
    `username` TEXT NOT NULL,
    `password` TEXT NOT NULL,
    `github_name` TEXT NULL DEFAULT '',
    `wx_open_id` TEXT NULL DEFAULT '',
    `created_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `memos` (
    `id` TEXT NOT NULL PRIMARY KEY,
    `content` TEXT NOT NULL,
    `user_id` TEXT NOT NULL,
    `created_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TEXT,
    FOREIGN KEY(`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `queries` (
    `id` TEXT NOT NULL PRIMARY KEY,
    `user_id` TEXT NOT NULL,
    `title` TEXT NOT NULL,
    `querystring` TEXT NOT NULL,
    `created_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `pinned_at`  TEXT NULL,
    FOREIGN KEY(`user_id`) REFERENCES `users`(`id`)
)
```


### 导入sqlite3
```go
sqlite3 mindStream.db < sqlite3.sql
```