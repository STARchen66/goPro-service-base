package session_dao

import (
	"goImPro-service/pkg/date"
	"goImPro-service/pkg/mysql"
	baseSession "goImPro-service/src/models/sessions"
	"goImPro-service/src/models/user"
)

type SessionDao struct {
}

func (s *SessionDao) CreateSession(formId int64, toId int64, channelType int) (sessions *baseSession.BaseSessions) {

	var users user.BaseUsers
	mysql.DB.Table("base_users").Where("id=?", toId).First(&users)
	session := baseSession.BaseSessions{
		ToId:        toId,
		FormId:      formId,
		CreatedAt:   date.NewDate(),
		TopStatus:   sessions.TopStatus,
		TopTime:     date.NewDate(),
		Note:        users.Name,
		ChannelType: channelType,
		Name:        users.Name,
		Avatar:      users.Avatar,
		Status:      baseSession.SessionStatusOk,
	}

	mysql.DB.Save(&session)

	return &session

}
