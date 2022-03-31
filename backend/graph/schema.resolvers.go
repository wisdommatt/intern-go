package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wisdommatt/intern-go/backend/graph/generated"
	"github.com/wisdommatt/intern-go/backend/graph/model"
)

func (r *mutationResolver) CreateStudentAccount(ctx context.Context, input model.NewStudent) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateOrganizationAccount(ctx context.Context, input model.NewOrganization) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) EmailLogin(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ApplyForJob(ctx context.Context, job string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAvailableJobs(ctx context.Context, search *string, pagination model.Pagination) ([]*model.Job, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAccount(ctx context.Context) (model.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
