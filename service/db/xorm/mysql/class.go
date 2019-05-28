package db

import (
	"github.com/go-xorm/core"

	"github.com/qinhan-shu/gp-server/logger"
	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// GetClassNum : get the number of all class
func (m *MysqlDriver) GetClassNum() (int64, error) {
	return m.conn.Count(&model.Class{})
}

// GetClasses : get classes by page num and page index
func (m *MysqlDriver) GetClasses(pageNum, pageIndex int64, keyword string) ([]*transform.IntactClass, error) {
	classes := make([]*model.Class, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Where("name like ?", "%"+keyword+"%").
		Find(&classes); err != nil {
		return nil, err
	}

	var intactClasses []*transform.IntactClass
	for _, class := range classes {
		announcements, err := m.GetAnnouncementsByClassID(class.Id)
		if err != nil {
			return nil, err
		}
		intactClasses = append(intactClasses, &transform.IntactClass{
			Class:         class,
			Announcements: announcements,
		})
	}

	return intactClasses, nil
}

// AddClass : add a new class
func (m *MysqlDriver) AddClass(intactClass *transform.IntactClass) error {
	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Insert(intactClass.Class)
	if err != nil {
		session.Rollback()
		return err
	}
	// set class id
	var announcements []*model.Announcement
	for _, announcement := range intactClass.Announcements {
		announcement.ClassId = intactClass.Class.Id
		announcements = append(announcements, &announcement.Announcement)
	}
	_, err = session.Insert(announcements)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// GetClassByID : get detail of class
func (m *MysqlDriver) GetClassByID(id int64) (*transform.IntactClass, error) {
	class := new(model.Class)
	ok, err := m.conn.Id(id).Get(class)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, ErrNoRowsFound
	}

	announcements, err := m.GetAnnouncementsByClassID(class.Id)
	if err != nil {
		return nil, err
	}
	return &transform.IntactClass{
		Class:         class,
		Announcements: announcements,
	}, nil
}

// UpdateClass : update class
func (m *MysqlDriver) UpdateClass(intactClass *transform.IntactClass) error {
	// set class id
	var announcements []model.Announcement
	for _, announcement := range intactClass.Announcements {
		announcement.Announcement.ClassId = intactClass.Class.Id
		announcements = append(announcements, announcement.Announcement)
	}

	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Id(intactClass.Class.Id).Update(intactClass.Class)
	if err != nil {
		session.Rollback()
		return err
	}

	_, err = session.Insert(announcements)
	if err != nil {
		session.Rollback()
		return err
	}

	return session.Commit()
}

// MemberManage : manage the menbers
func (m *MysqlDriver) MemberManage(manageType, classID, memberID int64) error {
	var affectd int64
	var err error
	switch manageType {
	case int64(protocol.MemberManageReq_DELETE):
		affectd, err = m.conn.Delete(&model.UserClass{
			ClassId: classID,
			UserId:  memberID,
		})
	case int64(protocol.MemberManageReq_SET_ADMINISTRATOR):
		affectd, err = m.conn.Id(core.PK{classID, memberID}).Update(&model.UserClass{
			IsAdministrator: 1,
			IsChecked:       1,
		})
	case int64(protocol.MemberManageReq_CANCEL_ADMINISTRATOR):
		affectd, err = m.conn.Id(core.PK{classID, memberID}).Update(&model.UserClass{
			IsAdministrator: 0,
			IsChecked:       1,
		})
	}
	if err != nil {
		return err
	}
	if affectd == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// GetMembers : get all class members
func (m *MysqlDriver) GetMembers(classID, pageNum, pageIndex int64) ([]*transform.UserClass, int64, error) {
	members := make([]*transform.UserClass, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
		Join("INNER", "user", "user.id = user_class.user_id").
		Where("class_id = ?", classID).
		Find(&members); err != nil {
		return nil, 0, err
	}

	num, err := m.conn.Where("class_id = ?", classID).Count(&model.UserClass{})
	if err != nil {
		logger.Sugar.Errorf("failed to get class members : %v", err)
		return nil, 0, err
	}

	return members, num, nil
}

// EnterClass : add member to class
func (m *MysqlDriver) EnterClass(userID, classID int64) error {
	class, err := m.GetClassByID(classID)
	if err != nil {
		logger.Sugar.Errorf("failed to enter class (get class fail) : %v", err)
		return err
	}

	isChecked := 1
	if class.Class.IsCheck == 1 {
		isChecked = 0
	}

	member := &model.UserClass{
		UserId:    userID,
		ClassId:   classID,
		IsChecked: isChecked,
	}

	i, err := m.conn.Insert(member)
	if err != nil {
		logger.Sugar.Errorf("failed to insert : %v", err)
		return err
	}
	if i == 0 {
		logger.Sugar.Error(ErrNoRowsAffected.Error())
		return ErrNoRowsAffected
	}
	return nil
}

// QuitClass : remove the player of class
func (m *MysqlDriver) QuitClass(userID, classID int64) {
	m.conn.Delete(&model.UserClass{
		UserId:  userID,
		ClassId: classID,
	})
}

// ApplyEnterRequest : agree of disagree the player request of enter class
func (m *MysqlDriver) ApplyEnterRequest(userID, classID int64, isApply bool) error {
	isChecked := 0
	if isApply {
		isChecked = 1
	}
	member := &model.UserClass{IsChecked: isChecked}

	_, err := m.conn.Where("user_id = ? and class_id = ?", userID, classID).Update(member)
	if err != nil {
		return err
	}
	return nil
}
