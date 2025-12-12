// Copyright (c) 2025 Half_nothing
// SPDX-License-Identifier: MIT

//go:build permission

// Package utils
package utils

type Permission uint64

// 权限节点上限是64
const (
	// PermissionAdminEntry 显示管理员入口
	PermissionAdminEntry Permission = 1 << iota
	// PermissionUserShowList 显示用户管理入口
	PermissionUserShowList
	// PermissionUserSetPassword 设置用户密码
	PermissionUserSetPassword
	// PermissionUserEditBaseInfo 编辑用户基本信息
	PermissionUserEditBaseInfo
	// PermissionUserShowPermission 显示用户权限管理入口
	PermissionUserShowPermission
	// PermissionUserEditPermission 编辑用户权限
	PermissionUserEditPermission
	// PermissionUserEditRole 编辑用户角色
	PermissionUserEditRole
	// PermissionUserBan 封禁用户
	PermissionUserBan
	// PermissionControllerShowList 显示管制员管理入口
	PermissionControllerShowList
	// PermissionControllerEditRating 编辑管制员权限
	PermissionControllerEditRating
	// PermissionControllerShowRecord 显示管制员履历入口
	PermissionControllerShowRecord
	// PermissionControllerCreateRecord 创建管制员履历
	PermissionControllerCreateRecord
	// PermissionControllerDeleteRecord 删除管制员履历
	PermissionControllerDeleteRecord
	// PermissionControllerTier2Rating 编辑管制员程序塔权限
	PermissionControllerTier2Rating
	// PermissionControllerChangeUnderMonitor 编辑管制员实习权限
	PermissionControllerChangeUnderMonitor
	// PermissionControllerChangeSolo 编辑管制员SOLO权限
	PermissionControllerChangeSolo
	// PermissionControllerChangeGuest 编辑管制员客座权限
	PermissionControllerChangeGuest
	// PermissionControllerApplicationShowList 显示管制员申请管理入口
	PermissionControllerApplicationShowList
	// PermissionControllerApplicationConfirm 确认管制员申请
	PermissionControllerApplicationConfirm
	// PermissionControllerApplicationPass 通过管制员申请
	PermissionControllerApplicationPass
	// PermissionControllerApplicationReject 拒绝管制员申请
	PermissionControllerApplicationReject
	// PermissionActivityShowList 显示活动管理入口
	PermissionActivityShowList
	// PermissionActivityPublish 发布活动
	PermissionActivityPublish
	// PermissionActivityEdit 编辑活动
	PermissionActivityEdit
	// PermissionActivityEditState 编辑活动状态
	PermissionActivityEditState
	// PermissionActivityEditPilotState 编辑活动飞行员状态
	PermissionActivityEditPilotState
	// PermissionActivityDelete 删除活动
	PermissionActivityDelete
	// PermissionAuditLogShow 显示审计日志入口
	PermissionAuditLogShow
	// PermissionTicketShowList 显示工单管理入口
	PermissionTicketShowList
	// PermissionTicketReply 回复工单
	PermissionTicketReply
	// PermissionTicketRemove 删除工单
	PermissionTicketRemove
	// PermissionFlightPlanShowList 显示飞行计划管理入口
	PermissionFlightPlanShowList
	// PermissionFlightPlanChangeLock 更改飞行计划锁定状态
	PermissionFlightPlanChangeLock
	// PermissionFlightPlanDelete 删除飞行计划
	PermissionFlightPlanDelete
	// PermissionClientManagerEntry 显示在线客户端管理入口
	PermissionClientManagerEntry
	// PermissionClientSendMessage 向客户端发送消息
	PermissionClientSendMessage
	// PermissionClientSendBroadcastMessage 发送广播消息
	PermissionClientSendBroadcastMessage
	// PermissionClientKill 从服务器踢出客户端
	PermissionClientKill
	// PermissionAnnouncementShowList 显示公告管理入口
	PermissionAnnouncementShowList
	// PermissionAnnouncementPublish 发布公告
	PermissionAnnouncementPublish
	// PermissionAnnouncementEdit 编辑公告
	PermissionAnnouncementEdit
	// PermissionAnnouncementDelete 删除公告
	PermissionAnnouncementDelete
	// PermissionRoleShowList 显示角色列表
	PermissionRoleShowList
	// PermissionRoleCreate 创建角色
	PermissionRoleCreate
	// PermissionRoleDelete 删除角色
	PermissionRoleDelete
	// PermissionRoleEdit 编辑角色
	PermissionRoleEdit
	// PermissionRoleEditPermission 编辑角色权限
	PermissionRoleEditPermission
	// PermissionInstructorShowList 显示教员管理入口
	PermissionInstructorShowList
	// PermissionInstructorCreate 创建教员
	PermissionInstructorCreate
	// PermissionInstructorDelete 删除教员
	PermissionInstructorDelete
	// PermissionInstructorEdit 编辑教员
	PermissionInstructorEdit
	// PermissionInstructorEditRating 编辑教员权限
	PermissionInstructorEditRating
	// PermissionInstructorShowAllocation 显示教员分配入口
	PermissionInstructorShowAllocation
	// PermissionInstructorAllocation 分配教员
	PermissionInstructorAllocation
)

var Permissions = NewEnums[string, Permission](
	NewEnum("AdminEntry", PermissionAdminEntry),
	NewEnum("UserShowList", PermissionUserShowList),
	NewEnum("UserSetPassword", PermissionUserSetPassword),
	NewEnum("UserEditBaseInfo", PermissionUserEditBaseInfo),
	NewEnum("UserShowPermission", PermissionUserShowPermission),
	NewEnum("UserEditPermission", PermissionUserEditPermission),
	NewEnum("UserEditRole", PermissionUserEditRole),
	NewEnum("UserBan", PermissionUserBan),
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
	NewEnum("RoleShowList", PermissionRoleShowList),
	NewEnum("RoleCreate", PermissionRoleCreate),
	NewEnum("RoleDelete", PermissionRoleDelete),
	NewEnum("RoleEdit", PermissionRoleEdit),
	NewEnum("RoleEditPermission", PermissionRoleEditPermission),
	NewEnum("InstructorShowList", PermissionInstructorShowList),
	NewEnum("InstructorCreate", PermissionInstructorCreate),
	NewEnum("InstructorDelete", PermissionInstructorDelete),
	NewEnum("InstructorEdit", PermissionInstructorEdit),
	NewEnum("InstructorEditRating", PermissionInstructorEditRating),
	NewEnum("InstructorShowAllocation", PermissionInstructorShowAllocation),
	NewEnum("InstructorAllocation", PermissionInstructorAllocation),
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
