// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"Web/app/dao/internal"
)

// userLogDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type userLogDao struct {
	*internal.UserLogDao
}

var (
	// UserLog is globally public accessible object for table user_log operations.
	UserLog = &userLogDao{
		internal.UserLog,
	}
)

// Fill with you ideas below.