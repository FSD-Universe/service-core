// Copyright (c) 2025-2026 Half_nothing
// SPDX-License-Identifier: MIT

//go:build database

// Package entity
package entity

import (
	"time"

	"half-nothing.cn/service-core/utils"
)

type AuditLog struct {
	ID            uint          `gorm:"primarykey"`
	Subject       string        `gorm:"type:text;not null"`
	Object        string        `gorm:"type:text;not null"`
	Event         string        `gorm:"size:64;index:idx_audit_logs_event;not null"`
	Ip            string        `gorm:"size:128;not null"`
	UserAgent     string        `gorm:"type:text;not null"`
	ChangeDetails *ChangeDetail `gorm:"type:text;serializer:json;default:null"`
	CreatedAt     time.Time
}

type ChangeDetail struct {
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}

func (a *AuditLog) GetId() uint {
	return a.ID
}

func (a *AuditLog) SetId(id uint) {
	a.ID = id
}

type AuditEvent *utils.Enum[string, string]

//goland:noinspection GoCommentStart
var (
	// 用户相关
	AuditEventUserRegistered       AuditEvent = utils.NewEnum("UserRegistered", "用户注册")
	AuditEventUserResetPassword    AuditEvent = utils.NewEnum("UserResetPassword", "用户重置密码")
	AuditEventUserInformationEdit  AuditEvent = utils.NewEnum("UserInformationEdit", "管理员修改用户信息")
	AuditEventUserPermissionGrant  AuditEvent = utils.NewEnum("UserPermissionGrant", "管理员授予用户权限")
	AuditEventUserPermissionRevoke AuditEvent = utils.NewEnum("UserPermissionRevoke", "管理员撤销用户权限")
	AuditEventUserBan              AuditEvent = utils.NewEnum("UserBan", "管理员封禁用户")
	AuditEventUserUnban            AuditEvent = utils.NewEnum("UserUnban", "管理员解封用户")

	// 角色相关
	AuditEventRoleCreated          AuditEvent = utils.NewEnum("RoleCreated", "管理员创建角色")
	AuditEventRoleUpdated          AuditEvent = utils.NewEnum("RoleUpdated", "管理员修改角色信息")
	AuditEventRoleDeleted          AuditEvent = utils.NewEnum("RoleDeleted", "管理员删除角色")
	AuditEventRolePermissionGrant  AuditEvent = utils.NewEnum("RolePermissionGrant", "管理员授予角色权限")
	AuditEventRolePermissionRevoke AuditEvent = utils.NewEnum("RolePermissionRevoke", "管理员撤销角色权限")
	AuditEventRoleGrant            AuditEvent = utils.NewEnum("RoleGrant", "管理员授予用户角色")
	AuditEventRoleRevoke           AuditEvent = utils.NewEnum("RoleRevoke", "管理员撤销用户角色")

	// 活动相关
	AuditEventActivityCreated             AuditEvent = utils.NewEnum("ActivityCreated", "管理员创建活动")
	AuditEventActivityDeleted             AuditEvent = utils.NewEnum("ActivityDeleted", "管理员删除活动")
	AuditEventActivityUpdated             AuditEvent = utils.NewEnum("ActivityUpdated", "管理员修改活动信息")
	AuditEventActivityPilotSign           AuditEvent = utils.NewEnum("ActivityPilotSign", "飞行员报名活动")
	AuditEventActivityPilotLeave          AuditEvent = utils.NewEnum("ActivityPilotLeave", "飞行员退出活动")
	AuditEventActivityPilotStatusChange   AuditEvent = utils.NewEnum("ActivityPilotStatusChange", "管理员修改飞行员活动状态")
	AuditEventActivityControllerJoin      AuditEvent = utils.NewEnum("AuditEventActivityControllerJoin", "管制员加入活动")
	AuditEventActivityControllerLeave     AuditEvent = utils.NewEnum("AuditEventActivityControllerLeave", "管制员退出活动")
	AuditEventActivityStatusChange        AuditEvent = utils.NewEnum("ActivityStatusChange", "管理员修改活动状态")
	AuditEventActivityCoordinationCreated AuditEvent = utils.NewEnum("ActivityCoordinationCreated", "管理员创建活动协调表")
	AuditEventActivityCoordinationUpdated AuditEvent = utils.NewEnum("ActivityCoordinationUpdated", "活动协调信息表被编辑")

	// 在线管理相关
	AuditEventClientKickedFsd        AuditEvent = utils.NewEnum("ClientKickedFromFsd", "管理员在FSD中踢出用户")
	AuditEventClientKicked           AuditEvent = utils.NewEnum("ClientKickedFromWeb", "管理员在WEB中踢出用户")
	AuditEventClientMessage          AuditEvent = utils.NewEnum("ClientMessage", "管理员发送消息给用户")
	AuditEventClientBroadcastMessage AuditEvent = utils.NewEnum("ClientBroadcastMessage", "管理员广播消息给用户")

	// 异常访问
	AuditEventUnlawfulOverreach AuditEvent = utils.NewEnum("UnlawfulOverreach", "用户发生非法越权访问")

	// 工单相关
	AuditEventTicketOpen    AuditEvent = utils.NewEnum("TicketOpen", "用户创建工单")
	AuditEventTicketClose   AuditEvent = utils.NewEnum("TicketClose", "用户或管理员关闭工单")
	AuditEventTicketDeleted AuditEvent = utils.NewEnum("TicketDeleted", "管理员删除工单")

	// 管制员相关
	AuditEventControllerRecordCreated         AuditEvent = utils.NewEnum("ControllerRecordCreated", "管理员创建管制员履历")
	AuditEventControllerRecordDeleted         AuditEvent = utils.NewEnum("ControllerRecordDeleted", "管理员删除管制员履历")
	AuditEventControllerRatingChange          AuditEvent = utils.NewEnum("ControllerRatingChange", "管理员修改管制员权限")
	AuditEventControllerApplicationSubmit     AuditEvent = utils.NewEnum("ControllerApplicationSubmit", "用户提交管制员申请")
	AuditEventControllerApplicationCancel     AuditEvent = utils.NewEnum("ControllerApplicationCancel", "用户取消管制员申请")
	AuditEventControllerApplicationPassed     AuditEvent = utils.NewEnum("ControllerApplicationPassed", "管理员通过管制员申请")
	AuditEventControllerApplicationProcessing AuditEvent = utils.NewEnum("ControllerApplicationProcessing", "管理员正在处理管制员申请")
	AuditEventControllerApplicationRejected   AuditEvent = utils.NewEnum("ControllerApplicationRejected", "管理员拒绝管制员申请")
	AuditEventControllerInstructorChange      AuditEvent = utils.NewEnum("ControllerInstructorChange", "管理员修改管制员教员")

	// 管制教员相关
	AuditEventInstructorSendEmail    AuditEvent = utils.NewEnum("InstructorSendEmail", "管制教员发送邮件")
	AuditEventInstructorRatingChange AuditEvent = utils.NewEnum("InstructorRatingChange", "管制教员权限变更")

	// 飞行计划相关
	AuditEventFlightPlanDeleted     AuditEvent = utils.NewEnum("FlightPlanDeleted", "管理员删除用户飞行计划")
	AuditEventFlightPlanSelfDeleted AuditEvent = utils.NewEnum("FlightPlanSelfDeleted", "用户删除自己的飞行计划")
	AuditEventFlightPlanLock        AuditEvent = utils.NewEnum("FlightPlanLock", "管制员锁定飞行计划")
	AuditEventFlightPlanUnlock      AuditEvent = utils.NewEnum("FlightPlanUnlock", "管制员解锁飞行计划")

	// 文件相关
	AuditEventFileUpload AuditEvent = utils.NewEnum("FileUpload", "用户上传文件")

	// 公告相关
	AuditEventAnnouncementPublished AuditEvent = utils.NewEnum("AnnouncementPublished", "管理员发布公告")
	AuditEventAnnouncementUpdated   AuditEvent = utils.NewEnum("AnnouncementUpdated", "管理员修改公告")
	AuditEventAnnouncementDeleted   AuditEvent = utils.NewEnum("AnnouncementDeleted", "管理员删除公告")
)

var AuditEventManager = utils.NewEnums(
	AuditEventUserRegistered,
	AuditEventUserResetPassword,
	AuditEventUserInformationEdit,
	AuditEventUserPermissionGrant,
	AuditEventUserPermissionRevoke,
	AuditEventUserBan,
	AuditEventUserUnban,
	AuditEventRoleCreated,
	AuditEventRoleUpdated,
	AuditEventRoleDeleted,
	AuditEventRolePermissionGrant,
	AuditEventRolePermissionRevoke,
	AuditEventRoleGrant,
	AuditEventRoleRevoke,
	AuditEventActivityCreated,
	AuditEventActivityDeleted,
	AuditEventActivityUpdated,
	AuditEventActivityPilotSign,
	AuditEventActivityPilotLeave,
	AuditEventActivityPilotStatusChange,
	AuditEventActivityControllerJoin,
	AuditEventActivityControllerLeave,
	AuditEventActivityStatusChange,
	AuditEventActivityCoordinationCreated,
	AuditEventActivityCoordinationUpdated,
	AuditEventClientKickedFsd,
	AuditEventClientKicked,
	AuditEventClientMessage,
	AuditEventClientBroadcastMessage,
	AuditEventUnlawfulOverreach,
	AuditEventTicketOpen,
	AuditEventTicketClose,
	AuditEventTicketDeleted,
	AuditEventControllerRecordCreated,
	AuditEventControllerRecordDeleted,
	AuditEventControllerRatingChange,
	AuditEventControllerApplicationSubmit,
	AuditEventControllerApplicationCancel,
	AuditEventControllerApplicationPassed,
	AuditEventControllerApplicationProcessing,
	AuditEventControllerApplicationRejected,
	AuditEventControllerInstructorChange,
	AuditEventInstructorSendEmail,
	AuditEventInstructorRatingChange,
	AuditEventFlightPlanDeleted,
	AuditEventFlightPlanSelfDeleted,
	AuditEventFlightPlanLock,
	AuditEventFlightPlanUnlock,
	AuditEventFileUpload,
	AuditEventAnnouncementPublished,
	AuditEventAnnouncementUpdated,
	AuditEventAnnouncementDeleted,
)
