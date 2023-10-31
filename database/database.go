package database

import (
	"context"
	"log"
	"time"

	"github.com/akashabbasi/gogql/graph/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connectionString string = "mongodb+srv://localhost:27017/gogql"

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(connectionString),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		30*time.Second,
	)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) GetJob(id string) *model.JobListing {
	var jobListing model.JobListing
	return &jobListing
}

func (db *DB) GetJobs() []*model.JobListing {
	var jobListings []*model.JobListing
	return jobListings
}

func (db *DB) CreateJobListing(
	jobInfo model.CreateJobListingInput,
) *model.JobListing {
	var returnJobListing *model.JobListing
	return returnJobListing
}

func (db *DB) UpdateJobListing(
	jobId string,
	jobInfo model.UpdateJobListingInput,
) *model.JobListing {
	var jobListing *model.JobListing
	return jobListing
}

func (db *DB) DeleteJobListing(jobId string) *model.DeleteJobResponse {
	return &model.DeleteJobResponse{}
}
