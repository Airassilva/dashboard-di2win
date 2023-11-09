package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"api/graph/model"
	"api/service"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Count is the resolver for the count field.
func (r *queryResolver) Count(ctx context.Context, groupBy string, userID *int, tipoDocumento *string, dataComeco *string, dataFinal *string) ([]*model.Count, error) {
	conn_str := os.Getenv("POSTGRE_URL")

	if conn_str == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}

		conn_str = os.Getenv("POSTGRE_URL")
	}

	rep := service.RepositorioPostgre{
		ConnStr: conn_str,
	}

	filter := service.FiltroExtract{}

	if dataComeco != nil {
		parsed_data_start, err := time.Parse(time.DateOnly, *dataComeco)
		if err != nil {
			log.Fatal(err)
		}

		filter.DataStart = parsed_data_start
	}

	if dataFinal != nil {
		parsed_data_final, err := time.Parse(time.DateOnly, *dataFinal)
		if err != nil {
			log.Fatal(err)
		}

		filter.DataEnd = parsed_data_final
	}

	if tipoDocumento != nil {
		filter.DocType = *tipoDocumento
	}

	if userID != nil {
		filter.UserId = *userID
	}

	count := rep.CountExtracts(groupBy, filter)

	return count, nil
}

// Extract is the resolver for the extract field.
func (r *queryResolver) Extract(ctx context.Context, id *int, userID *int, tipoDocumento *string, dataComeco *string, dataFinal *string) ([]*model.Extract, error) {
	conn_str := os.Getenv("POSTGRE_URL")

	if conn_str == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}

		conn_str = os.Getenv("POSTGRE_URL")
	}

	rep := service.RepositorioPostgre{
		ConnStr: conn_str,
	}

	if id != nil {
		extract, err := rep.FindExtract(*id)
		if err != nil {
			log.Fatal(err)
		}

		extract_parsed := model.Extract{
			ID:             fmt.Sprintf("%d", extract.Id),
			CreatedAt:      extract.CreatedAt.String(),
			PagesProcessed: extract.PagesProcessed,
			DocType:        extract.DocType,
			UserID:         extract.UserId,
		}

		return []*model.Extract{&extract_parsed}, nil
	}

	filter := service.FiltroExtract{}

	if dataComeco != nil {
		parsed_data_start, err := time.Parse(time.DateOnly, *dataComeco)
		if err != nil {
			log.Fatal(err)
		}

		filter.DataStart = parsed_data_start
	}

	if dataFinal != nil {
		parsed_data_final, err := time.Parse(time.DateOnly, *dataFinal)
		if err != nil {
			log.Fatal(err)
		}

		filter.DataEnd = parsed_data_final
	}

	if tipoDocumento != nil {
		filter.DocType = *tipoDocumento
	}

	if userID != nil {
		filter.UserId = *userID
	}

	extracts_return := []*model.Extract{}
	extracts_found := rep.FindExtracts(filter)

	for _, value := range extracts_found {
		extract := model.Extract{
			ID:             fmt.Sprintf("%d", value.Id),
			CreatedAt:      value.CreatedAt.String(),
			PagesProcessed: value.PagesProcessed,
			DocType:        value.DocType,
			UserID:         value.UserId,
		}

		extracts_return = append(extracts_return, &extract)
	}

	return extracts_return, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id *int, segment *string) ([]*model.User, error) {
	conn_str := os.Getenv("POSTGRE_URL")

	if conn_str == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal(err)
		}

		conn_str = os.Getenv("POSTGRE_URL")
	}

	rep := service.RepositorioPostgre{
		ConnStr: conn_str,
	}

	if id != nil {
		user_found, err := rep.FindUser(*id)
		if err != nil {
			log.Fatal(err)
		}

		user_return := model.User{
			ID:      fmt.Sprintf("%d", user_found.Id),
			Name:    user_found.Name,
			Segment: user_found.Segment,
		}

		return []*model.User{&user_return}, nil
	}

	if segment != nil {
		user_found := rep.FindUsers(*segment)
		users_return := []*model.User{}

		for _, value := range user_found {
			user_return := model.User{
				ID:      fmt.Sprintf("%d", value.Id),
				Name:    value.Name,
				Segment: value.Segment,
			}

			users_return = append(users_return, &user_return)
		}

		return users_return, nil
	}

	users_found := rep.FindUsers("")
	users_return := []*model.User{}

	for _, value := range users_found {
		user_return := model.User{
			ID:      fmt.Sprintf("%d", value.Id),
			Name:    value.Name,
			Segment: value.Segment,
		}

		users_return = append(users_return, &user_return)
	}

	return users_return, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
