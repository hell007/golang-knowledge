package FilterPattern

//为标准（Criteria）创建一个接口。
type Criteria interface {
	MeetCriteria(persons []Person) []Person
}

type CriteriaMale struct {
}

//按照性别男过滤
func (s *CriteriaMale) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, person := range persons {
		if person.Gender == "Male" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type CriteriaFemale struct {
}

//按照性别女过滤
func (s *CriteriaFemale) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, person := range persons {
		if person.Gender == "Female" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type CriteriaSingle struct {
}

//按照未婚过滤
func (s *CriteriaSingle) MeetCriteria(persons []Person) []Person {
	var femalePersons []Person
	for _, person := range persons {
		if person.MaritalStatus == "Single" {
			femalePersons = append(femalePersons, person)
		}
	}
	return femalePersons
}

type AndCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}

//使用需要的过滤组合
func (s *AndCriteria) AndCriteria(criteria Criteria, otherCriteria Criteria) {
	s.criteria = criteria
	s.otherCriteria = otherCriteria
}

//多重组合过滤
func (s *AndCriteria) MeetCriteria(persons []Person) []Person {
	firstCriteriaPersons := s.criteria.MeetCriteria(persons)
	return s.otherCriteria.MeetCriteria(firstCriteriaPersons)
}

type OrCriteria struct {
	criteria      Criteria
	otherCriteria Criteria
}