package post

import (
	"errors"

	"go.fedejuret.blog/modules/database"
)

var connection = database.GetConnection()

func GetPostById(id int) Post {

	var post Post

	err := connection.QueryRow("SELECT * FROM posts WHERE id = ?", id).Scan(&post.Id, &post.Title, &post.Content, &post.Image, &post.Status, &post.Created)

	if err != nil {
		panic(err.Error())
	}

	return post
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

func CreatePost(post Post) Post {

	result, err := connection.Exec("INSERT INTO posts (title, content, image, status) VALUES (?, ?, ?, ?)", post.Title, post.Content, post.Image, post.Status)

	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	newId := int(id)

	post.Id = &newId

	return post

}

func UpdatePost(postId int, post Post) (Post, error) {

	result, err := connection.Exec("UPDATE posts SET title = ?, content = ?, image = ?, status = ? WHERE id = ?", post.Title, post.Content, post.Image, post.Status, postId)

	if err != nil {
		return Post{}, err
	}

	if rows, _ := result.RowsAffected(); rows != 1 {
		return Post{}, errors.New("Post not found")
	}

	return GetPostById(postId), nil

}

func DeletePost(id int) (bool, error) {

	result, err := connection.Exec("DELETE FROM posts WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	if rows, _ := result.RowsAffected(); rows != 1 {
		return false, errors.New("Post not found")
	}

	return true, nil

}
