package server

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	cloud "github.com/darkjedidj/ticketgenerator/api/aws"
	ticket "github.com/darkjedidj/ticketgenerator/internal/filegen"
	pb "github.com/darkjedidj/ticketgenerator/internal/proto"
)

type Server struct {
	pb.UnimplementedTicketGeneratorServer
}

// GetTicket creates PDF file, stores it in S3 and returns ID
func (s *Server) GetTicket(ctx context.Context, in *pb.TicketRequset) (*pb.IDReply, error) {

	err := ticket.Generate(in)
	if err != nil {
		return nil, err
	}
	defer ticket.Delete(in.Id)

	file, err := os.Open(fmt.Sprintf("internal/tickets/%v.pdf", in.Id))
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
		return nil, err
	}

	session := cloud.AwsService.GetSession()

	up := s3manager.NewUploader(session)

	S3BucketName := os.Getenv("BUCKET_NAME")

	_, err = up.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(S3BucketName),
		Key:         aws.String(fmt.Sprintf("%v", in.Id)),
		Body:        file,
		ContentType: aws.String("application/pdf"),
	})
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	return &pb.IDReply{ID: in.Id}, nil
}
