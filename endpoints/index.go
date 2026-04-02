package endpoints

import (
	get "portfolio-backend/endpoints/get"
	post "portfolio-backend/endpoints/post"
	put "portfolio-backend/endpoints/put"
	utils "portfolio-backend/utils"
)

/*  TODO: To include/merge the files/endpoint
 * to call it into the frontend
 */

var Routes = []utils.Route{
	// TODO: GET Requests
	get.Index,
	get.Baybayin,
	get.Blog,
	get.Certificates,
	get.Contact,
	get.Dev,
	get.Experiences,
	get.Feedback,
	get.Manga,
	get.Poetry,
	get.Projects,
	get.Retrieve,
	get.YoutubeDL,

	// TODO: Created AI Endpoint
	post.AIAgent,

	// TODO: POST Requests
	post.Blog,
	post.Certificates,
	post.Contact,
	post.Feedback,
	post.Poetry,
	post.Upload,

	// TODO: PUT Requests
	put.Blog,
	put.Experience,

	// TODO: Cookie Handler
	get.Cookie,
}
