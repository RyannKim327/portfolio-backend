package endpoints

import utils "portfolio-backend/utils"

/*  TODO: To include/merge the files/endpoint
 * to call it into the frontend
 */

var Routes = []utils.Route{
	// TODO: GET Requests
	Index,
	GetBaybayin,
	GetExperiences,
	GetFeedback,
	GetPoetry,
	GetProjects,

	// TODO: Created AI Endpoint
	PostAIAgent,

	// TODO: POST Requests
	PostFeedback,
	PostPoetry,
}
