package data_model

import {} 

// STRUCT of Projects with projectId, adminId, list of literature, Total rating, total number of ratings, and the list of ratings.
type Projects struct {
	ProjectId string `json:"projectId"`
	AdminId string `json:"adminId"`
	Literature []Literature `json:"literature"`
	TotalRating int `json:"totalRating"`
	TotalRatings int `json:"totalRatings"`
}

// STRUCT of Literature with literatureId, title, description, and the list of ratings.
type Literature struct {
	LiteratureId string `json:"literatureId"`
	Title string `json:"title"`
	Description string `json:"description"`
}

// STRUCT of Ratings with ratingId, userId, rating, and the list of comments.
type Ratings struct {
	RatingId string `json:"ratingId"`
	ProjectId string `json:"projectId"`
	UserId string `json:"userId"`
	Rating int `json:"rating"`
	Comments []Comments `json:"comments"`
}

// STRUCT of Comments with commentId, userId, and comment.
type Comments struct {
	CommentId string `json:"commentId"`
	UserId string `json:"userId"`
	Comment string `json:"comment"`
}

// STRUCT of Users with userId, username, password, and the list of projects.
type Users struct {
	UserId string `json:"userId"`
	Projects []Projects `json:"projects"`
}

// STRUCT of Admins with adminId, username, password, and the list of projects.
type Admins struct {
	AdminId string `json:"adminId"`
	Username string `json:"username"`
	Password string `json:"password"`
	Projects []Projects `json:"projects"`
}


