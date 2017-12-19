package service

import (
    "agenda_api/cli/entity"
    "agenda_api/cli/dao"
)

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}


// FindByUsername .
func (*UserInfoAtomicService) FindByUsername(name string) (*entity.User, error) {
    dao := dao.UserInfoDao{entity.Mydb}
    return dao.FindByUsername(name)
}

//InsertAUser
func (*UserInfoAtomicService) InsertUser(user entity.User) error {
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.UserInfoDao{tx}
    err = dao.InsertUser(user)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//Insert Login Info 
func (*UserInfoAtomicService) LoginInfoInsert(session entity.Session) error {
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.UserInfoDao{tx}
    err = dao.LoginInfoInsert(session)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}

//delete User's log in infomation by username
func (*UserInfoAtomicService) LoginInfoDelete(session entity.Session) error {
    tx, err := entity.Mydb.Begin()
    entity.CheckErr(err)

    dao := dao.UserInfoDao{tx}
    err = dao.LoginInfoDelete(session)

    if err == nil {
        tx.Commit()
    } else {
		tx.Rollback()
		return err;
    }
    return nil
}
//determine whether a user has been logged in by username or not
func (*UserInfoAtomicService) UserHasLogin(loginid string) bool {
    dao := dao.UserInfoDao{entity.Mydb}
    session, _ := dao.UserHasLogin(loginid)
    return session != nil
}

//get all users' infomation
func (*UserInfoAtomicService) GetAllUsersInfo(session entity.Session) ([]entity.User, error) {
    dao := dao.UserInfoDao{entity.Mydb}
    return dao.GetAllUsersInfo(session)
}
