package load_balancer

type ClassInfo struct {
	Enrollment    int
	MaxEnrollment int
	ClassNumber   int
}

type Cache struct {
	Classes map[int]*ClassInfo // the maps from classNumbers to enrollment stats
}

func (cache *Cache) AddStudent(classNumber int, studentID int) {

}
