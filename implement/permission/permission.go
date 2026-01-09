// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build permission

// Package permission
package permission

import "half-nothing.cn/service-core/utils"

type Permission uint64

// 权限节点上限是64
const (
	// AdminEntry 显示管理员入口
	AdminEntry Permission = 1 << iota
	// UserShowList 显示用户管理入口
	UserShowList
	// UserEditInfo 编辑用户信息
	UserEditInfo
	// UserShowPermission 显示用户权限管理入口
	UserShowPermission
	// UserEditPermission 编辑用户权限
	UserEditPermission
	// UserShowRole 显示用户角色
	UserShowRole
	// UserEditRole 编辑用户角色
	UserEditRole
	// UserBan 封禁用户
	UserBan
	// ControllerShowList 显示管制员管理入口
	ControllerShowList
	// ControllerEditRating 编辑管制员权限
	ControllerEditRating
	// ControllerShowRecord 显示管制员履历入口
	ControllerShowRecord
	// ControllerCreateRecord 创建管制员履历
	ControllerCreateRecord
	// ControllerDeleteRecord 删除管制员履历
	ControllerDeleteRecord
	// ControllerTier2Rating 编辑管制员程序塔权限
	ControllerTier2Rating
	// ControllerChangeUnderMonitor 编辑管制员实习权限
	ControllerChangeUnderMonitor
	// ControllerChangeSolo 编辑管制员SOLO权限
	ControllerChangeSolo
	// ControllerChangeGuest 编辑管制员客座权限
	ControllerChangeGuest
	// ControllerApplicationShowList 显示管制员申请管理入口
	ControllerApplicationShowList
	// ControllerApplicationConfirm 确认管制员申请
	ControllerApplicationConfirm
	// ControllerApplicationPass 通过管制员申请
	ControllerApplicationPass
	// ControllerApplicationReject 拒绝管制员申请
	ControllerApplicationReject
	// ActivityShowList 显示活动管理入口
	ActivityShowList
	// ActivityPublish 发布活动
	ActivityPublish
	// ActivityEdit 编辑活动
	ActivityEdit
	// ActivityEditState 编辑活动状态
	ActivityEditState
	// ActivityEditPilotState 编辑活动飞行员状态
	ActivityEditPilotState
	// ActivityDelete 删除活动
	ActivityDelete
	// AuditLogShow 显示审计日志入口
	AuditLogShow
	// TicketShowList 显示工单管理入口
	TicketShowList
	// TicketReply 回复工单
	TicketReply
	// TicketRemove 删除工单
	TicketRemove
	// FlightPlanShowList 显示飞行计划管理入口
	FlightPlanShowList
	// FlightPlanChangeLock 更改飞行计划锁定状态
	FlightPlanChangeLock
	// FlightPlanDelete 删除飞行计划
	FlightPlanDelete
	// ClientManagerEntry 显示在线客户端管理入口
	ClientManagerEntry
	// ClientSendMessage 向客户端发送消息
	ClientSendMessage
	// ClientSendBroadcastMessage 发送广播消息
	ClientSendBroadcastMessage
	// ClientKill 从服务器踢出客户端
	ClientKill
	// AnnouncementShowList 显示公告管理入口
	AnnouncementShowList
	// AnnouncementPublish 发布公告
	AnnouncementPublish
	// AnnouncementEdit 编辑公告
	AnnouncementEdit
	// AnnouncementDelete 删除公告
	AnnouncementDelete
	// RoleShowList 显示角色列表
	RoleShowList
	// RoleCreate 创建角色
	RoleCreate
	// RoleDelete 删除角色
	RoleDelete
	// RoleEdit 编辑角色
	RoleEdit
	// RoleEditPermission 编辑角色权限
	RoleEditPermission
	// InstructorShowList 显示教员管理入口
	InstructorShowList
	// InstructorCreate 创建教员
	InstructorCreate
	// InstructorDelete 删除教员
	InstructorDelete
	// InstructorEdit 编辑教员
	InstructorEdit
	// InstructorEditRating 编辑教员权限
	InstructorEditRating
	// InstructorShowAllocation 显示教员分配入口
	InstructorShowAllocation
	// InstructorAllocation 分配教员
	InstructorAllocation
)

var Permissions = utils.NewEnums[string, Permission](
	utils.NewEnum("AdminEntry", AdminEntry),
	utils.NewEnum("UserShowList", UserShowList),
	utils.NewEnum("UserEditInfo", UserEditInfo),
	utils.NewEnum("UserShowPermission", UserShowPermission),
	utils.NewEnum("UserEditPermission", UserEditPermission),
	utils.NewEnum("UserShowRole", UserShowRole),
	utils.NewEnum("UserEditRole", UserEditRole),
	utils.NewEnum("UserBan", UserBan),
	utils.NewEnum("ControllerShowList", ControllerShowList),
	utils.NewEnum("ControllerTier2Rating", ControllerTier2Rating),
	utils.NewEnum("ControllerEditRating", ControllerEditRating),
	utils.NewEnum("ControllerShowRecord", ControllerShowRecord),
	utils.NewEnum("ControllerCreateRecord", ControllerCreateRecord),
	utils.NewEnum("ControllerDeleteRecord", ControllerDeleteRecord),
	utils.NewEnum("ControllerChangeUnderMonitor", ControllerChangeUnderMonitor),
	utils.NewEnum("ControllerChangeSolo", ControllerChangeSolo),
	utils.NewEnum("ControllerChangeGuest", ControllerChangeGuest),
	utils.NewEnum("ControllerApplicationShowList", ControllerApplicationShowList),
	utils.NewEnum("ControllerApplicationConfirm", ControllerApplicationConfirm),
	utils.NewEnum("ControllerApplicationPass", ControllerApplicationPass),
	utils.NewEnum("ControllerApplicationReject", ControllerApplicationReject),
	utils.NewEnum("ActivityPublish", ActivityPublish),
	utils.NewEnum("ActivityShowList", ActivityShowList),
	utils.NewEnum("ActivityEdit", ActivityEdit),
	utils.NewEnum("ActivityEditState", ActivityEditState),
	utils.NewEnum("ActivityEditPilotState", ActivityEditPilotState),
	utils.NewEnum("ActivityDelete", ActivityDelete),
	utils.NewEnum("AuditLogShow", AuditLogShow),
	utils.NewEnum("TicketShowList", TicketShowList),
	utils.NewEnum("TicketReply", TicketReply),
	utils.NewEnum("TicketRemove", TicketRemove),
	utils.NewEnum("FlightPlanShowList", FlightPlanShowList),
	utils.NewEnum("FlightPlanChangeLock", FlightPlanChangeLock),
	utils.NewEnum("FlightPlanDelete", FlightPlanDelete),
	utils.NewEnum("ClientManagerEntry", ClientManagerEntry),
	utils.NewEnum("ClientSendMessage", ClientSendMessage),
	utils.NewEnum("ClientKill", ClientKill),
	utils.NewEnum("ClientSendBroadcastMessage", ClientSendBroadcastMessage),
	utils.NewEnum("AnnouncementShowList", AnnouncementShowList),
	utils.NewEnum("AnnouncementPublish", AnnouncementPublish),
	utils.NewEnum("AnnouncementEdit", AnnouncementEdit),
	utils.NewEnum("AnnouncementDelete", AnnouncementDelete),
	utils.NewEnum("RoleShowList", RoleShowList),
	utils.NewEnum("RoleCreate", RoleCreate),
	utils.NewEnum("RoleDelete", RoleDelete),
	utils.NewEnum("RoleEdit", RoleEdit),
	utils.NewEnum("RoleEditPermission", RoleEditPermission),
	utils.NewEnum("InstructorShowList", InstructorShowList),
	utils.NewEnum("InstructorCreate", InstructorCreate),
	utils.NewEnum("InstructorDelete", InstructorDelete),
	utils.NewEnum("InstructorEdit", InstructorEdit),
	utils.NewEnum("InstructorEditRating", InstructorEditRating),
	utils.NewEnum("InstructorShowAllocation", InstructorShowAllocation),
	utils.NewEnum("InstructorAllocation", InstructorAllocation),
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
