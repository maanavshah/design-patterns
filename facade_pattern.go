package main

import "fmt"

type Parser struct{}

func (p *Parser) ParseRelevanceData() {
	fmt.Println("Parsing Relevance Data")
}
func (p *Parser) ParseRankerData() {
	fmt.Println("Parsing Ranker Data")
}

type RedisClient struct{}

func (r *RedisClient) SaveRankerToRedis() {
	fmt.Println("Saved Ranker Data in Redis")
}
func (r *RedisClient) SaveRelevanceToRedis() {
	fmt.Println("Saved Relevance Data in Redis")
}

type MongoClient struct{}

func (r *MongoClient) SaveRankerToMongo() {
	fmt.Println("Saved Ranker Data in Mongo")
}
func (r *MongoClient) SaveRelevanceToMongo() {
	fmt.Println("Saved Relevance Data in Mongo")
}

type RelevanceFacade struct {
	parser      *Parser
	redisClient *RedisClient
	mongoClient *MongoClient
}

func NewRelevanceFacade(parser *Parser, redisClient *RedisClient, mongoClient *MongoClient) RelevanceFacade {
	return RelevanceFacade{
		parser:      parser,
		redisClient: redisClient,
		mongoClient: mongoClient,
	}
}
func (r RelevanceFacade) SaveRelevance() {
	r.parser.ParseRelevanceData()
	r.redisClient.SaveRelevanceToRedis()
	r.mongoClient.SaveRelevanceToMongo()
}

type RankerFacade struct {
	parser      *Parser
	redisClient *RedisClient
}

func NewRankerFacade(parser *Parser, redisClient *RedisClient) RankerFacade {
	return RankerFacade{
		parser:      parser,
		redisClient: redisClient,
	}
}
func (r RankerFacade) SaveRanker() {
	r.parser.ParseRankerData()
	r.redisClient.SaveRankerToRedis()
}

func main() {
	parser := &Parser{}
	redisClient := &RedisClient{}
	mongoClient := &MongoClient{}

	relevanceFacade := NewRelevanceFacade(parser, redisClient, mongoClient)
	relevanceFacade.SaveRelevance()

	rankerFacade := NewRankerFacade(parser, redisClient)
	rankerFacade.SaveRanker()
}
