package domain

/*
* Loggger Interface for question model
*
*	CampaignError(err error) --> to inform an error
*	CampaignInfo(info string) --> to general information
*
 */
type QuestionLogger interface {
	QuestionError(err error)
	QuestionInfo(info string)
}
