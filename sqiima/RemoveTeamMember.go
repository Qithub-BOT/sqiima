package sqiima

// RemoveTeamMember は指定のユーザーをチームから離脱させる際のオブジェクトを表します.
// Note: 自身とチームのオーナーはこの API では離脱させられません.
type RemoveTeamMember map[string]interface{}
