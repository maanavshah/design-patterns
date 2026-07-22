import java.util.logging.Logger;

public class SingletonPattern {

    private static Logger newLogger() {
        try {
            Thread.sleep(1000);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
        }
        return Logger.getLogger("app-" + System.nanoTime());
    }

    public static Logger getNewLogger() {
        return newLogger();
    }

    private static class Holder {
        private static final Logger INSTANCE = newLogger();
    }

    public static Logger getSingletonLogger() {
        return Holder.INSTANCE;
    }

    public static void main(String[] args) {
        System.out.println("Starting application...");

        System.out.println("Using a new logger instance each time:");
        for (int i = 0; i < 5; i++) {
            Logger log = getNewLogger();
            log.info("Log message " + (i + 1) + " from new logger");
        }

        System.out.println("Using a singleton logger instance:");
        for (int i = 0; i < 5; i++) {
            Logger log = getSingletonLogger();
            log.info("Log message " + (i + 1) + " from singleton logger");
        }
    }
}
