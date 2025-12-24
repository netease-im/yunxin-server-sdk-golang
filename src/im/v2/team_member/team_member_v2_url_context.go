package team_member

// URL context for Team Member V2 APIs
const (
	InviteTeamMembers    = "/im/v2/team_members"                                     // Invite Team Members endpoint
	KickTeamMembers      = "/im/v2/team_members/actions/kick_member"                 // Kick Team Members endpoint
	LeaveTeam            = "/im/v2/team_members/actions/leave"                       // Leave Team endpoint
	UpdateTeamMember     = "/im/v2.1/team_members/{account_id}"                      // Update Team Member endpoint
	BatchMuteTeamMembers = "/im/v2/team_members/actions/batch_mute"                  // Batch Mute Team Members endpoint
	JoinedTeams          = "/im/v2.1/team_members/{account_id}/actions/joined_teams" // Query Joined Teams endpoint
)
