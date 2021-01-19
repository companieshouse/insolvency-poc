package models

type AppointLiquidatorResponse struct {
	Id    string                         `json:"id"`
	Links AppointLiquidatorResponseLinks `json:"links"`
}

type AppointLiquidatorResponseLinks struct {
	Self string `json:"self"`
}

type NoticeStatementAffairsResponse struct {
	Id    string                              `json:"statement_id"`
	Links NoticeStatementAffairsResponseLinks `json:"links"`
}

type NoticeStatementAffairsResponseLinks struct {
	Self string `json:"self"`
}
