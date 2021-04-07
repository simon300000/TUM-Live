package model

type Student struct {
	ID      string `gorm:"primaryKey"` // currently matrikelnr. as soon as we get a reply from the it service "obfuscatedID"
	Name    string
	Courses []Course `gorm:"many2many:course_students;"` // sql back reference
}

func (s *Student) CoursesForSemester(year int, term string) []Course {
	var cRes []Course
	for _, c := range s.Courses {
		if c.Year == year && c.TeachingTerm == term {
			cRes = append(cRes, c)
		}
	}
	return cRes
}
