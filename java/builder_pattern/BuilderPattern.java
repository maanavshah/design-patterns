final class Request {
    private final String requestID;
    private final int userID;
    private final String keyword;
    private final int sessionID;
    private final String mode;

    Request(String requestID, int userID, String keyword, int sessionID, String mode) {
        this.requestID = requestID;
        this.userID = userID;
        this.keyword = keyword;
        this.sessionID = sessionID;
        this.mode = mode;
    }

    @Override
    public String toString() {
        return "{requestID=" + requestID
                + ", userID=" + userID
                + ", keyword=" + keyword
                + ", sessionID=" + sessionID
                + ", mode=" + mode + "}";
    }
}

class RequestBuilder {
    private String requestID = "";
    private int userID = 0;
    private String keyword = "";
    private int sessionID = 0;
    private String mode = "";

    public RequestBuilder setRequestID(String requestID) {
        this.requestID = requestID;
        return this;
    }

    public RequestBuilder setUserID(int userID) {
        this.userID = userID;
        return this;
    }

    public RequestBuilder setKeyword(String keyword) {
        this.keyword = keyword;
        return this;
    }

    public RequestBuilder setSessionID(int sessionID) {
        this.sessionID = sessionID;
        return this;
    }

    public RequestBuilder setMode(String mode) {
        this.mode = mode;
        return this;
    }

    public Request build() {
        return new Request(requestID, userID, keyword, sessionID, mode);
    }

    public Request buildWithValidation() {
        if (requestID == null || requestID.isEmpty()) {
            throw new IllegalStateException("requestID is required");
        }
        return build();
    }

    public RequestBuilder reset() {
        this.requestID = "";
        this.userID = 0;
        this.keyword = "";
        this.sessionID = 0;
        this.mode = "";
        return this;
    }
}

public class BuilderPattern {
    public static void main(String[] args) {
        Request request = new RequestBuilder().setRequestID("request1").build();
        System.out.println("request: " + request);

        request = new RequestBuilder().setMode("autosuggest").setKeyword("query").build();
        System.out.println("request: " + request);

        RequestBuilder builder = new RequestBuilder();
        Request req1 = builder.setRequestID("request1").setUserID(100).build();
        Request req2 = builder.setRequestID("request2").setUserID(200).build();

        System.out.println("req1: " + req1);
        System.out.println("req2: " + req2);

        try {
            Request req3 = new RequestBuilder().setMode("search").buildWithValidation();
            System.out.println("req3: " + req3);
        } catch (IllegalStateException e) {
            System.out.println("Validation error: " + e.getMessage());
        }

        builder.reset().setRequestID("request3").setMode("search");
        Request req4 = builder.build();
        System.out.println("req4: " + req4);
    }
}
