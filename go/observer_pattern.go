package main

import "fmt"

type ObservableInterface interface {
	Add(obs ObserverInterface)
	Notify(data string)
}

type RelevanceObservable struct {
	observers []ObserverInterface
}

func (ro *RelevanceObservable) Add(obs ObserverInterface) {
	ro.observers = append(ro.observers, obs)
}

func (ro *RelevanceObservable) Notify(data string) {
	for _, observer := range ro.observers {
		observer.Update(data)
	}
}

type ObserverInterface interface {
	Update(data string)
}

type RedisObserver struct {
	data string
}

func (o *RedisObserver) Update(data string) {
	o.data = data
	fmt.Printf("Redis data updated: %s\n", data)
}

type MongoObserver struct {
	data string
}

func (o *MongoObserver) Update(data string) {
	o.data = data
	fmt.Printf("Mongo data updated: %s\n", data)
}

func main() {
	relevanceObserable := RelevanceObservable{}
	
	redisClient := RedisObserver{}
	mongoClient := MongoObserver{}

	relevanceObserable.Add(&redisClient)
	relevanceObserable.Add(&mongoClient)

	relevanceObserable.Notify("data 1")
	relevanceObserable.Notify("data 2")
}
