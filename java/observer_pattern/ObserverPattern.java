import java.util.ArrayList;
import java.util.List;

interface ObserverInterface {
    void update(String data);
}

interface ObservableInterface {
    void add(ObserverInterface obs);

    void notify(String data);
}

class RelevanceObservable implements ObservableInterface {
    private final List<ObserverInterface> observers = new ArrayList<>();

    @Override
    public void add(ObserverInterface obs) {
        observers.add(obs);
    }

    @Override
    public void notify(String data) {
        for (ObserverInterface observer : observers) {
            observer.update(data);
        }
    }
}

class RedisObserver implements ObserverInterface {
    private String data;

    @Override
    public void update(String data) {
        this.data = data;
        System.out.println("Redis data updated: " + data);
    }
}

class MongoObserver implements ObserverInterface {
    private String data;

    @Override
    public void update(String data) {
        this.data = data;
        System.out.println("Mongo data updated: " + data);
    }
}

public class ObserverPattern {
    public static void main(String[] args) {
        RelevanceObservable relevanceObservable = new RelevanceObservable();

        RedisObserver redisClient = new RedisObserver();
        MongoObserver mongoClient = new MongoObserver();

        relevanceObservable.add(redisClient);
        relevanceObservable.add(mongoClient);

        relevanceObservable.notify("data 1");
        relevanceObservable.notify("data 2");
    }
}
