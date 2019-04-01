package db

import (
	"github.com/go-xorm/core"

	"github.com/qinhan-shu/gp-server/model/transform"
	"github.com/qinhan-shu/gp-server/model/xorm"
	"github.com/qinhan-shu/gp-server/protocol"
)

// GetClassNum : get the number of all class
func (m *MysqlDriver) GetClassNum() (int64, error) {
	return m.conn.Count(&model.Class{})
}

// GetClasses : get classes by page num and page index
func (m *MysqlDriver) GetClasses(pageNum, pageIndex int64) ([]*transform.IntactClass, error) {
	classes := make([]*model.Class, 0)
	if err := m.conn.
		Limit(int(pageNum), int((pageIndex-1)*pageNum)).
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
	for _, announcement := range intactClass.Announcements {
		announcement.ClassId = intactClass.Class.Id
	}
	_, err = session.Insert(intactClass.Announcements)
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
	for _, announcement := range intactClass.Announcements {
		announcement.ClassId = intactClass.Class.Id
	}

	session := m.conn.NewSession()
	defer session.Close()

	err := session.Begin()
	_, err = session.Id(intactClass.Class.Id).Update(intactClass.Class)
	if err != nil {
		session.Rollback()
		return err
	}

	_, err = session.Insert(intactClass.Announcements)
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
