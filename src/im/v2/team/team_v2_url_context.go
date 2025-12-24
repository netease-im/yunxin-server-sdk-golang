package team

// URL context for Team V2 APIs
const (
	CreateTeam                       = "/im/v2.1/teams"                                     // Create Team endpoint
	UpdateTeam                       = "/im/v2.1/teams/{team_id}"                           // Update Team endpoint
	DisbandTeam                      = "/im/v2.1/teams/{team_id}"                           // Disband Team endpoint
	BatchQueryTeams                  = "/im/v2.1/teams"                                     // Batch Query Teams endpoint
	TransferOwner                    = "/im/v2/teams/{team_id}/actions/transfer_owner"      // Transfer Team Owner endpoint
	AddManager                       = "/im/v2/teams/{team_id}/actions/add_manager"         // Add Team Manager endpoint
	RemoveManager                    = "/im/v2/teams/{team_id}/actions/remove_manager"      // Remove Team Manager endpoint
	GetTeamInfo                      = "/im/v2.1/teams/{team_id}"                           // Get Team Info endpoint
	ListTeamMembers                  = "/im/v2.1/teams/{team_id}/actions/list_members"      // List Team Members endpoint
	ListOnlineTeamMembers            = "/im/v2/teams/{team_id}/actions/list_online_members" // List Online Team Members endpoint
	BatchQueryTeamOnlineMembersCount = "/im/v2/teams/actions/online_members_count"          // Batch Query Team Online Members Count endpoint
)
