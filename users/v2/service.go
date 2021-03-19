package v2

import (
	"context"
	"fmt"
	"io"
	"user-service/proto"
)

type service struct {
	database Database
}

func (s service) GetAllUserSummariesPaginated(stream proto.UserSummary_GetAllUserSummariesPaginatedServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if in.LastId == 0{

		}
		users := s.database.GetAllUsersPaginated(in.LastId)
		for _, user := range users {
			if err := stream.Send(user); err != nil {
				return err
			}
		}
	}

}

func (s service) GetAllUserSummaries(request *proto.UserSummaryRequest, stream proto.UserSummary_GetAllUserSummariesServer) error {
	users := s.database.GetAllUsers()
	for _, user := range users {
		fmt.Println(user.CreatedAt)

		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s service) GetUserSummary(ctx context.Context, req *proto.UserSummaryRequest) (*proto.UserSummaryResponse, error) {
	users := s.database.GetUserSummaryV2(req.Ids)
	results := &proto.UserSummaryResponse{Users: users}
	return results, nil

}

func New(db Database) service {
	return service {
		database: db,
	}
}