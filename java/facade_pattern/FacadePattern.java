class Parser {
    public void parseRelevanceData() {
        System.out.println("Parsing Relevance Data");
    }

    public void parseRankerData() {
        System.out.println("Parsing Ranker Data");
    }
}

class RedisClient {
    public void saveRankerToRedis() {
        System.out.println("Saved Ranker Data in Redis");
    }

    public void saveRelevanceToRedis() {
        System.out.println("Saved Relevance Data in Redis");
    }
}

class MongoClient {
    public void saveRankerToMongo() {
        System.out.println("Saved Ranker Data in Mongo");
    }

    public void saveRelevanceToMongo() {
        System.out.println("Saved Relevance Data in Mongo");
    }
}

class RelevanceFacade {
    private final Parser parser;
    private final RedisClient redisClient;
    private final MongoClient mongoClient;

    public RelevanceFacade(Parser parser, RedisClient redisClient, MongoClient mongoClient) {
        this.parser = parser;
        this.redisClient = redisClient;
        this.mongoClient = mongoClient;
    }

    public void saveRelevance() {
        parser.parseRelevanceData();
        redisClient.saveRelevanceToRedis();
        mongoClient.saveRelevanceToMongo();
    }
}

class RankerFacade {
    private final Parser parser;
    private final RedisClient redisClient;

    public RankerFacade(Parser parser, RedisClient redisClient) {
        this.parser = parser;
        this.redisClient = redisClient;
    }

    public void saveRanker() {
        parser.parseRankerData();
        redisClient.saveRankerToRedis();
    }
}

public class FacadePattern {
    public static void main(String[] args) {
        Parser parser = new Parser();
        RedisClient redisClient = new RedisClient();
        MongoClient mongoClient = new MongoClient();

        RelevanceFacade relevanceFacade = new RelevanceFacade(parser, redisClient, mongoClient);
        relevanceFacade.saveRelevance();

        RankerFacade rankerFacade = new RankerFacade(parser, redisClient);
        rankerFacade.saveRanker();
    }
}
