package endpoints

import (
	get "portfolio-backend/endpoints/get"
	post "portfolio-backend/endpoints/post"
	utils "portfolio-backend/utils"
)

/*  TODO: To include/merge the files/endpoint
 * to call it into the frontend
 */

var Routes = []utils.Route{
	// TODO: GET Requests
	get.Index,
	get.Baybayin,
	get.Experiences,
	get.Feedback,
	get.Poetry,
	get.Projects,

	// TODO: Created AI Endpoint
	post.AIAgent,

	// TODO: POST Requests
	post.Feedback,
	post.Poetry,

	// TODO: Cookie Handler
	get.Cookie,
}
