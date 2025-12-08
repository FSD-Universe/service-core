//go:build permission

// Package utils
package utils

type Permission uint64

// 权限节点上限是64
const (
	PermissionAdminEntry Permission = 1 << iota
	PermissionUserShowList
	PermissionUserGetProfile
	PermissionUserSetPassword
	PermissionUserEditBaseInfo
	PermissionUserShowPermission
	PermissionUserEditPermission
	PermissionControllerShowList
	PermissionControllerTier2Rating
	PermissionControllerEditRating
	PermissionControllerShowRecord
	PermissionControllerCreateRecord
	PermissionControllerDeleteRecord
	PermissionControllerChangeUnderMonitor
	PermissionControllerChangeSolo
	PermissionControllerChangeGuest
	PermissionControllerApplicationShowList
	PermissionControllerApplicationConfirm
	PermissionControllerApplicationPass
	PermissionControllerApplicationReject
	PermissionActivityPublish
	PermissionActivityShowList
	PermissionActivityEdit
	PermissionActivityEditState
	PermissionActivityEditPilotState
	PermissionActivityDelete
	PermissionAuditLogShow
	PermissionTicketShowList
	PermissionTicketReply
	PermissionTicketRemove
	PermissionFlightPlanShowList
	PermissionFlightPlanChangeLock
	PermissionFlightPlanDelete
	PermissionClientManagerEntry
	PermissionClientSendMessage
	PermissionClientSendBroadcastMessage
	PermissionClientKill
	PermissionAnnouncementShowList
	PermissionAnnouncementPublish
	PermissionAnnouncementEdit
	PermissionAnnouncementDelete
)

var Permissions = NewEnums[string, Permission](
	NewEnum("AdminEntry", PermissionAdminEntry),
	NewEnum("UserShowList", PermissionUserShowList),
	NewEnum("UserGetProfile", PermissionUserGetProfile),
	NewEnum("UserSetPassword", PermissionUserSetPassword),
	NewEnum("UserEditBaseInfo", PermissionUserEditBaseInfo),
	NewEnum("UserShowPermission", PermissionUserShowPermission),
	NewEnum("UserEditPermission", PermissionUserEditPermission),
	NewEnum("ControllerShowList", PermissionControllerShowList),
	NewEnum("ControllerTier2Rating", PermissionControllerTier2Rating),
	NewEnum("ControllerEditRating", PermissionControllerEditRating),
	NewEnum("ControllerShowRecord", PermissionControllerShowRecord),
	NewEnum("ControllerCreateRecord", PermissionControllerCreateRecord),
	NewEnum("ControllerDeleteRecord", PermissionControllerDeleteRecord),
	NewEnum("ControllerChangeUnderMonitor", PermissionControllerChangeUnderMonitor),
	NewEnum("ControllerChangeSolo", PermissionControllerChangeSolo),
	NewEnum("ControllerChangeGuest", PermissionControllerChangeGuest),
	NewEnum("ControllerApplicationShowList", PermissionControllerApplicationShowList),
	NewEnum("ControllerApplicationConfirm", PermissionControllerApplicationConfirm),
	NewEnum("ControllerApplicationPass", PermissionControllerApplicationPass),
	NewEnum("ControllerApplicationReject", PermissionControllerApplicationReject),
	NewEnum("ActivityPublish", PermissionActivityPublish),
	NewEnum("ActivityShowList", PermissionActivityShowList),
	NewEnum("ActivityEdit", PermissionActivityEdit),
	NewEnum("ActivityEditState", PermissionActivityEditState),
	NewEnum("ActivityEditPilotState", PermissionActivityEditPilotState),
	NewEnum("ActivityDelete", PermissionActivityDelete),
	NewEnum("AuditLogShow", PermissionAuditLogShow),
	NewEnum("TicketShowList", PermissionTicketShowList),
	NewEnum("TicketReply", PermissionTicketReply),
	NewEnum("TicketRemove", PermissionTicketRemove),
	NewEnum("FlightPlanShowList", PermissionFlightPlanShowList),
	NewEnum("FlightPlanChangeLock", PermissionFlightPlanChangeLock),
	NewEnum("FlightPlanDelete", PermissionFlightPlanDelete),
	NewEnum("ClientManagerEntry", PermissionClientManagerEntry),
	NewEnum("ClientSendMessage", PermissionClientSendMessage),
	NewEnum("ClientKill", PermissionClientKill),
	NewEnum("ClientSendBroadcastMessage", PermissionClientSendBroadcastMessage),
	NewEnum("AnnouncementShowList", PermissionAnnouncementShowList),
	NewEnum("AnnouncementPublish", PermissionAnnouncementPublish),
	NewEnum("AnnouncementEdit", PermissionAnnouncementEdit),
	NewEnum("AnnouncementDelete", PermissionAnnouncementDelete),
)

func (p *Permission) HasPermission(perm Permission) bool {
	return *p&perm == perm
}

func (p *Permission) Merge(perm Permission) {
	*p |= perm
}

func (p *Permission) Grant(perm Permission) {
	*p |= perm
}

func (p *Permission) Revoke(perm Permission) {
	*p &^= perm
}
