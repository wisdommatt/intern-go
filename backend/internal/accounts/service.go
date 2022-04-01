package accounts

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wisdommatt/intern-go/backend/graph/model"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	collection     *firestore.CollectionRef
	jobsCollection *firestore.CollectionRef
}

var (
	errSomethingWentWrong = fmt.Errorf("an error occured, please try again later")
)

// NewService creates a new accounts service.
func NewService(firestoreClient *firestore.Client) *Service {
	return &Service{
		collection: firestoreClient.Collection("accounts"),
	}
}

func (s *Service) CreateAccount(ctx context.Context, account model.Account) (*model.Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("an error occured while hashing password")
	}
	log.Println("creating account: ", account)
	account.Password = string(hashedPassword)
	doc, _, err := s.collection.Add(ctx, account)
	if err != nil {
		log.Println("error: ", err)
		return nil, errSomethingWentWrong
	}
	account.ID = doc.ID
	return &account, nil
}

func (s *Service) EmailLogin(ctx context.Context, email, password string) (*model.Account, error) {
	iter := s.collection.Where("Email", "==", email).Documents(ctx)
	doc, err := iter.Next()
	if err != nil {
		log.Println("error: ", err)
		return nil, fmt.Errorf("invalid credentials")
	}
	var acc model.Account
	err = doc.DataTo(&acc)
	if err != nil {
		log.Println("error: ", err)
		return nil, fmt.Errorf("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(acc.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	acc.ID = doc.Ref.ID
	return &acc, nil
}

func (s *Service) CreateJob(ctx context.Context, job model.Job) (*model.Job, error) {
	job.TimeAdded = time.Now()
	doc, _, err := s.collection.Add(ctx, job)
	if err != nil {
		log.Println("error: ", err)
		return nil, errSomethingWentWrong
	}
	job.ID = doc.ID
	return &job, nil
}

func (s *Service) GetJobs(ctx context.Context) ([]model.Job, error) {
	iter := s.collection.Documents(ctx)
	docs, err := iter.GetAll()
	if err != nil {
		log.Println("error: ", err)
		return nil, errSomethingWentWrong
	}
	var jobs []model.Job
	for _, doc := range docs {
		var job model.Job
		err = doc.DataTo(&job)
		if err != nil {
			log.Println("error: ", err)
			return nil, errSomethingWentWrong
		}
		job.ID = doc.Ref.ID
		jobs = append(jobs, job)
	}
	return jobs, nil
}

func (s *Service) GetAccount(ctx context.Context, id string) (*model.Account, error) {
	doc, err := s.collection.Doc(id).Get(ctx)
	if err != nil {
		log.Println("error: ", err)
		return nil, errSomethingWentWrong
	}
	var acc model.Account
	err = doc.DataTo(&acc)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	}
	acc.ID = doc.Ref.ID
	return &acc, nil
}
