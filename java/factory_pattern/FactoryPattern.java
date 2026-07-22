interface Seat {
    String getDetails();

    int getPrice();
}

class EconomySeat implements Seat {
    @Override
    public String getDetails() {
        return "Economy: Basic seat, no extras";
    }

    @Override
    public int getPrice() {
        return 100;
    }
}

class PremiumSeat implements Seat {
    @Override
    public String getDetails() {
        return "Premium: Comfortable seat, extra legroom";
    }

    @Override
    public int getPrice() {
        return 200;
    }
}

class VIPSeat implements Seat {
    @Override
    public String getDetails() {
        return "VIP: Luxury recliner, complimentary snacks";
    }

    @Override
    public int getPrice() {
        return 500;
    }
}

public class FactoryPattern {

    public static Seat bookSeat(String seatType) {
        switch (seatType) {
            case "economy":
                return new EconomySeat();
            case "premium":
                return new PremiumSeat();
            case "vip":
                return new VIPSeat();
            default:
                return new EconomySeat();
        }
    }

    public static void main(String[] args) {
        String[] bookings = {"economy", "premium", "vip"};

        for (String booking : bookings) {
            Seat seat = bookSeat(booking);
            System.out.printf("%s - Price: ₹%d%n", seat.getDetails(), seat.getPrice());
        }
    }
}
