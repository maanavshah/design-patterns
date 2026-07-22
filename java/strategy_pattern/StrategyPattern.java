interface PaymentInterface {
    void pay(double amount);
}

class CreditCard implements PaymentInterface {
    private final String cardNumber;

    public CreditCard(String cardNumber) {
        this.cardNumber = cardNumber;
    }

    @Override
    public void pay(double amount) {
        if (amount <= 0) {
            throw new IllegalArgumentException("invalid amount");
        }
        String last4 = cardNumber.substring(cardNumber.length() - 4);
        System.out.printf("Paid $%.2f using Credit Card ****%s%n", amount, last4);
    }
}

class UPI implements PaymentInterface {
    private final String upiId;

    public UPI(String upiId) {
        this.upiId = upiId;
    }

    @Override
    public void pay(double amount) {
        if (amount <= 0) {
            throw new IllegalArgumentException("invalid amount");
        }
        System.out.printf("Paid $%.2f using UPI ID: %s%n", amount, upiId);
    }
}

class PaymentContext {
    private PaymentInterface strategy;

    public void setStrategy(PaymentInterface strategy) {
        if (strategy == null) {
            throw new IllegalArgumentException("strategy cannot be nil");
        }
        this.strategy = strategy;
    }

    public void pay(double amount) {
        if (strategy == null) {
            throw new IllegalStateException("no payment strategy set");
        }
        strategy.pay(amount);
    }
}

public class StrategyPattern {
    public static void main(String[] args) {
        PaymentContext payment = new PaymentContext();

        try {
            payment.pay(100.0);
        } catch (IllegalStateException e) {
            System.out.println("Error: " + e.getMessage());
        }

        CreditCard cc = new CreditCard("1234567890123456");
        payment.setStrategy(cc);
        payment.pay(150.50);

        UPI upi = new UPI("user@paytm");
        payment.setStrategy(upi);
        payment.pay(75.25);
    }
}
