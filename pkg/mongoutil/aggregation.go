package mongoutil

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AggregateFromText(ctx context.Context, collection *mongo.Collection, query string, subqueries ...string) (*mongo.Cursor, error) {
	var stages bson.A

	err := json.Unmarshal([]byte(query), &stages)
	if err != nil {
		return nil, fmt.Errorf("AggregateFromText: Error unmarshaling query: %v", err)
	}
	if len(subqueries) > 0 {
		for i, sq := range subqueries {
			var substages bson.A
			err := json.Unmarshal([]byte(sq), &substages)
			if err != nil {
				return nil, fmt.Errorf("AggregateFromText: Error unmarshaling subquery %d: %v", i, err)
			}
			stages = append(stages, substages...)
		}
	}

	return collection.Aggregate(ctx, stages)
}
