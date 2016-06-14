package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ApplicationContext struct {
	db                   *gorm.DB
	FreelancerRepository *FreelancerRepository
	ProjectRepository    *ProjectRepository
	ClientRepository     *ClientRepository
	ReferenceRepository  *ReferenceRepository
	JobRepository        *JobRepository
	JwtSecret            string
}

type ContextOptions struct {
	DbName string
	DbUser string
	DbPass string
	Secret string
}

func NewContext(options ContextOptions) (*ApplicationContext, error) {
	db, err := gorm.Open("postgres", "user="+options.DbUser+" password="+options.DbPass+" dbname="+options.DbName+" sslmode=disable")
	if err != nil {
		return nil, err
	}

	freelancerRepository, _ := NewFreelancerRepository(db)
	projectRepository, _ := NewProjectRepository(db)
	clientRepository, _ := NewClientRepository(db)
	referenceRepository, _ := NewReferenceRepository(db)
	jobRepository, _ := NewJobRepository(db)

	context := &ApplicationContext{
		db:                   db,
		FreelancerRepository: freelancerRepository,
		ProjectRepository:    projectRepository,
		ClientRepository:     clientRepository,
		ReferenceRepository:  referenceRepository,
		JobRepository:        jobRepository,
		JwtSecret:            options.Secret, //base64.StdEncoding.EncodeToString([]byte(options.Secret)),
	}

	return context, nil
}

func (ac *ApplicationContext) DropCreateFillTables() {
	ac.DropTables()
	ac.CreateTables()
	ac.FillTables()
}

func (ac *ApplicationContext) DropTables() {
	ac.db.DropTableIfExists(&Freelancer{}, &Project{}, &Client{}, &Job{}, &Review{}, &Reference{}, &Media{})
}

func (ac *ApplicationContext) CreateTables() {
	ac.db.CreateTable(&Freelancer{}, &Project{}, &Client{}, &Job{}, &Review{}, &Reference{}, &Media{})
}

func (ac *ApplicationContext) FillTables() {
	ac.FreelancerRepository.AddFreelancer(NewFreelancer("First", "Last", "Dev", "Pass", "first@mail.com", 3, 55, "UTC"))

	ac.FreelancerRepository.AddReview(&Review{
		Title:        "text2",
		Content:      "content",
		Rating:       4.1,
		ClientId:     1,
		FreelancerId: 1,
	})

	ac.ReferenceRepository.AddReference(&Reference{
		Title:        "title",
		Content:      "content",
		Media:        Media{"image", "video"},
		FreelancerId: 1,
	})

	ac.FreelancerRepository.AddFreelancer(NewFreelancer(
		"Pera",
		"Peric",
		"Dev",
		"123456",
		"second@mail.com",
		12,
		22,
		"CET",
	))

	ac.FreelancerRepository.AddReview(&Review{
		Title:        "text2",
		Content:      "content",
		Rating:       4.1,
		JobId:        1,
		ClientId:     1,
		FreelancerId: 2,
	})

	ac.FreelancerRepository.AddReview(&Review{
		Title:        "text2",
		Content:      "content",
		Rating:       2.4,
		JobId:        2,
		ClientId:     2,
		FreelancerId: 2,
	})
	ac.ReferenceRepository.AddReference(&Reference{
		Title:        "title",
		Content:      "content",
		Media:        Media{"image", "video"},
		FreelancerId: 2,
	})

	ac.db.Create(&Project{
		Name:        "Project",
		Description: "Description",
		ClientId:    1,
		IsActive:    true,
	})

	ac.db.Create(&Client{
		Name:        "Client",
		Description: "Desc Client",
	})

	ac.db.Create(&Job{
		Name:        "Job",
		Description: "Desc Job",
		ClientId:    1,
	})

	ac.FreelancerRepository.AddFreelancer(NewFreelancer("Third", "Last", "Dev", "Pass", "third@mail.com", 3, 55, "UTC"))
	ac.FreelancerRepository.DeleteFreelancer(3)
}
