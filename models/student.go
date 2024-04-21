package models

import (
	"encoding/json"
	"time"
)

// 创建本地数据库，并创建一张Student表，有字段：id（主键），student_no，name，gender，birth，往数据库中随机生产10万条学生信息
type Student struct {
	Id        int       `gorm:"column:id;primaryKey" json:"id" form:"id"`
	StudentNo int       `gorm:"column:student_no;index" json:"student_no" form:"student_no" binding:"required" unique:"true"`
	Name      string    `gorm:"column:name" json:"name" form:"name" binding:"required"`
	Gender    string    `gorm:"column:gender" json:"gender" form:"gender" binding:"required" enum:"男,女"`
	Birth     time.Time `gorm:"column:birth" json:"birth" form:"birth" binding:"required"`
}

func (s *Student) MarshalJSON() ([]byte, error) {
	type Alias Student
	return json.Marshal(&struct {
		Birth string `json:"birth"`
		*Alias
	}{
		Birth: s.Birth.Format("2006-01-02"),
		Alias: (*Alias)(s),
	})
}

func GetStudent(page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetStudentByName(name string, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("name like ?", name+"%").Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetStudentByBirth(start time.Time, end time.Time, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("birth between ? and ?", start, end).Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func GetStudentByBirthRangeAndName(name string, start time.Time, end time.Time, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("name like ? and birth between ? and ?", name+"%", start, end).Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// 查询姓名和出生日期大于指定日期的学生
func GetStudentByNameAndBirthBiggerThan(name string, birth time.Time, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("name = ? and birth >= ?", name, birth).Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// 查询姓名和出生日期小于指定日期的学生
func GetStudentByNameAndBirthLessThan(name string, birth time.Time, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("name = ? and birth <= ?", name, birth).Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// 查询出生日期大于指定日期的学生
func GetStudentByBirthBiggerThan(birth time.Time, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("birth >= ?", birth).Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

// 查询出生日期小于指定日期的学生
func GetStudentByBirthLessThan(birth time.Time, page int, page_size int) ([]Student, error) {
	var students []Student
	if err := DB.Where("birth <= ?", birth).Offset((page - 1) * page_size).Limit(page_size).Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
