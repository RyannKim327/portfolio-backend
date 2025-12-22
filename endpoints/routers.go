package endpoints

import utils "portfolio-backend/utils"

/*  TODO: To include/merge the files/endpoint
 * to call it into the frontend
 */

var Routes = []utils.Route{
	// TODO: GET Requests
	Index,
	Experiences,
	Feedback,
	Poetry,
	Projects,

	// TODO: POST Requests
	PostFeedback,
}
