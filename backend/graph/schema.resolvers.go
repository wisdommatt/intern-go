package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/wisdommatt/intern-go/backend/graph/generated"
	"github.com/wisdommatt/intern-go/backend/graph/model"
)

func (r *mutationResolver) CreateAccount(ctx context.Context, input model.NewAccount) (bool, error) {
	_, err := r.AccountService.CreateAccount(ctx, model.Account{
		Name:        input.Name,
		Email:       input.Email,
		Password:    input.Password,
		ResumeURL:   input.ResumeURL,
		Skills:      input.Skills,
		Bio:         input.Bio,
		Description: input.Description,
		AccountType: input.AccountType,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) EmailLogin(ctx context.Context, email string, password string) (*model.AuthResponse, error) {
	account, err := r.AccountService.EmailLogin(ctx, email, password)
	if err != nil {
		return nil, err
	}
	return &model.AuthResponse{
		Account: account,
	}, nil
}

func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {
	return r.AccountService.CreateJob(ctx, model.Job{
		Role:        input.Role,
		Location:    input.Location,
		CompanyName: input.CompanyName,
		JobType:     input.JobType,
		Skills:      input.Skills,
		Description: input.Description,
	})
}

func (r *queryResolver) GetAvailableJobs(ctx context.Context) ([]model.Job, error) {
	return r.AccountService.GetJobs(ctx)
}

func (r *queryResolver) GetAccount(ctx context.Context, accountID string) (*model.Account, error) {
	return r.AccountService.GetAccount(ctx, accountID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) ApplyForJob(ctx context.Context, job string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
