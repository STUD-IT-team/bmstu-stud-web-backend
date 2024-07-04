package responses

type GetClubMembers struct {
	ID       int          `json:"id"`
	MainOrgs []MainOrg    `json:"main_orgs"`
	SubOrgs  []SubClubOrg `json:"sub_orgs"`
}
