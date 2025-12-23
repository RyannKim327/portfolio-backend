package endpoints

import utils "portfolio-backend/utils"

/*  TODO: To include/merge the files/endpoint
 * to call it into the frontend
 */

var Routes = []utils.Route{
	// TODO: GET Requests
	Index,
	GetExperiences,
	GetFeedback,
	GetPoetry,
	GetProjects,

	// TODO: POST Requests
	PostAIAgent,
	PostFeedback,
	PostPoetry,
}
