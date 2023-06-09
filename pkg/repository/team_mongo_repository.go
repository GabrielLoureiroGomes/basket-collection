package repository

import (
	"context"

	"github.com/GabrielLoureiroGomes/basket-collection/core/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type MongoRepository struct {
	mongo *mongoClient
}

type TeamMongoDocument struct {
	Name       string `bson:"name"`
	Conference string `bson:"conference"`
	State      string `bson:"state"`
}

func newTeamMongoDocument(team *domain.Team) TeamMongoDocument {
	return TeamMongoDocument{
		Name:       team.GetName(),
		Conference: team.GetConference(),
		State:      team.GetState(),
	}
}

func NewMongoRepository(ctx context.Context) MongoRepository {
	return MongoRepository{
		mongo: newMongoClient(ctx),
	}
}

// InsertTeam used to save on database team data
func (m MongoRepository) InsertTeam(ctx context.Context, team *domain.Team) error {

	teamDocumentToInsert := newTeamMongoDocument(team)
	_, err := m.getCollection().InsertOne(ctx, teamDocumentToInsert)

	if err != nil {
		log.Error("error to insert team data on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return err
	}

	return nil
}

// GetAll used to get all database team data
func (m MongoRepository) GetAll(ctx context.Context) ([]*domain.Team, error) {
	filter := bson.D{}

	cursor, err := m.getCollection().Find(ctx, filter)
	if err != nil {
		log.Error("error to get team data on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return []*domain.Team{}, err
	}

	teamsMongoDocument := []TeamMongoDocument{}
	if err = cursor.All(ctx, &teamsMongoDocument); err != nil {
		return []*domain.Team{}, err
	}

	teamsToReturn := []*domain.Team{}
	for _, team := range teamsMongoDocument {
		teamsToReturn = append(teamsToReturn, &domain.Team{
			Name:       team.Name,
			Conference: team.Conference,
			State:      team.State,
		})
	}

	return teamsToReturn, nil
}

// GetOne used to get one database team data
func (m MongoRepository) GetOne(ctx context.Context, teamName string) (*domain.Team, error) {
	filter := bson.D{{Key: "name", Value: teamName}}

	teamMongoDocument := TeamMongoDocument{}
	if err := m.getCollection().FindOne(ctx, filter).Decode(&teamMongoDocument); err != nil {
		log.Error("error to get one team on mongo", zap.Field{Type: zapcore.StringType, String: err.Error()})
		return &domain.Team{}, err
	}

	teamToReturn := &domain.Team{
		Name:       teamMongoDocument.Name,
		Conference: teamMongoDocument.Conference,
		State:      teamMongoDocument.State,
	}

	return teamToReturn, nil
}

func (m MongoRepository) getCollection() *mongo.Collection {
	databaseName := viper.GetString("MONGO_DATABASE_NAME")
	teamCollection := viper.GetString("MONGO_TEAM_COLLECTION")

	return m.mongo.client.Database(databaseName).Collection(teamCollection)
}
