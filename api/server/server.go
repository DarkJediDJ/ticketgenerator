package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	cloud "github.com/darkjedidj/ticketgenerator/api/aws"
	pb "github.com/darkjedidj/ticketgenerator/api/proto"
	ticket "github.com/darkjedidj/ticketgenerator/internal/filegen"
	"github.com/joho/godotenv"
)

type Server struct {
	pb.UnimplementedTicketGeneratorServer
}

func (s *Server) GetTicket(ctx context.Context, in *pb.TicketRequset) (*pb.LinkReply, error) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = ticket.Generate(in)
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

	svc := s3.New(session)
	params := &s3.GetObjectInput{
		Bucket: aws.String(S3BucketName),
		Key:    aws.String(fmt.Sprintf("%v", in.Id)),
	}

	req, _ := svc.GetObjectRequest(params)

	url, err := req.Presign(15 * time.Minute) // Set link expiration time
	if err != nil {
		log.Fatalf("[AWS GET LINK]: %q,%v", params, err)
		return nil, err
	}

	return &pb.LinkReply{Link: url}, nil
}
