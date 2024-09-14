package db

import "ministryOfEducation/models"

func Migrate() error {
	err := dbConn.AutoMigrate(models.Teacher{},
		models.TeacherType{},
		models.University{},
		models.AcademicDegree{},
		models.AcademicPosition{},
		models.City{},
		models.DirectionSpec{},
		models.Spec{},
		models.Admin{})
	if err != nil {
		return err
	}
	return nil
}
