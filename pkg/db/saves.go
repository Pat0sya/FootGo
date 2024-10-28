package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Match struct {
	Team1  string   `bson:"team1"`
	Team2  string   `bson:"team2"`
	Score1 int      `bson:"score1"`
	Score2 int      `bson:"score2"`
	Events []string `bson:"events"`
}

// SaveMatch сохраняет информацию о завершенном матче в MongoDB
func (m *Match) SaveMatch(collection *mongo.Collection) error {
	matchData := bson.M{
		"team1":  m.Team1,
		"team2":  m.Team2,
		"score1": m.Score1,
		"score2": m.Score2,
		"events": m.Events,
	}

	_, err := collection.InsertOne(context.TODO(), matchData)
	return err
}
