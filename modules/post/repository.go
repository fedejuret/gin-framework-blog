package post

import (
	"go.fedejuret.blog/modules/database"
)

var connection = database.GetConnection()

func GetPostById(id int) (Post, error) {

	var post Post

	err := connection.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&post.Id, &post.Title, &post.Content, &post.Image, &post.Status)

	if err != nil {
		panic(err.Error())
	}

	return post, nil
}

func GetAllPosts() []Post {

	var post []Post

	results, err := connection.Query("SELECT * FROM posts")

	if err != nil {
		panic(err.Error())
	}

	for results.Next() {

		var p Post

		err = results.Scan(&p.Id, &p.Title, &p.Content, &p.Image, &p.Status, &p.Created)

		if err != nil {
			panic(err.Error())
		}

		post = append(post, p)

	}

	return post
}
