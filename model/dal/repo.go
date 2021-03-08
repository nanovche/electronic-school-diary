package dal

type IRepository interface {
	GetTeacherRepository() IRepositoryTeacher
	GetAdminRepository() RepositoryAdminImpl
	GetAuthRepository() IRepositoryAuth
	GetSubjectRepository() IRepositorySubject
	GetStudentRepository() IStudentRepository
	GetMarkRepository() IMarkRepository
	GetTermRepository() ITermRepository
}

type RepositoryImpl struct {
	subjectRepo IRepositorySubject
	teacherRepo IRepositoryTeacher
	adminRepo RepositoryAdminImpl
	authRepo IRepositoryAuth
	studentRepo IStudentRepository
	markRepo IMarkRepository
	termRepo ITermRepository
}

func(repo *RepositoryImpl) GetTermRepository() ITermRepository{
	return repo.termRepo
}

func(repo *RepositoryImpl) SetTermRepository(termRepo ITermRepository) {
	repo.termRepo = termRepo
}

func(repo *RepositoryImpl) GetMarkRepository() IMarkRepository{
	return repo.markRepo
}

func(repo *RepositoryImpl) SetMarkRepository(markRepo IMarkRepository) {
	repo.markRepo = markRepo
}

func(repo *RepositoryImpl) GetSubjectRepository() IRepositorySubject{
	return repo.subjectRepo
}

func(repo *RepositoryImpl) SetSubjectRepository(subjectRepo IRepositorySubject) {
	repo.subjectRepo = subjectRepo
}

func(repo *RepositoryImpl) GetStudentRepository() IStudentRepository{
	return repo.studentRepo
}

func(repo *RepositoryImpl) SetStudentRepository(studentRepo IStudentRepository) {
	repo.studentRepo = studentRepo
}

func(repo *RepositoryImpl) GetAuthRepository() IRepositoryAuth{
	return repo.authRepo
}

func(repo *RepositoryImpl) SetAuthRepository(authRepo IRepositoryAuth) {
	repo.authRepo = authRepo
}

func(repo *RepositoryImpl) GetTeacherRepository() IRepositoryTeacher{
	return repo.teacherRepo
}

func(repo *RepositoryImpl) SetRepositoryTeacher(teacherRepo IRepositoryTeacher) {
	repo.teacherRepo = teacherRepo
}

func(repo *RepositoryImpl) GetAdminRepository() RepositoryAdminImpl{
	return repo.adminRepo
}

func(repo *RepositoryImpl) SetRepositoryAdmin(adminRepo RepositoryAdminImpl) {
	repo.adminRepo = adminRepo
}


