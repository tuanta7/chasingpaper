# Integration

## 1. PayPal

### 1.1. Price, Plan and Subscriptions

PayPal subscriptions rely on a structure where a Plan (template with frequency) defines the Price (fixed or quantity-based amount), which is then applied to individual user Subscriptions.

- **Price** (Cost): Defines how much is charged. This can be a fixed, recurring amount, or based on quantity.
- **Plan** (Template): Defines pricing and billing cycle details for subscriptions, including the name, description, frequency (day, week, month, year), and trial periods.
- **Subscriptions** (Contract): The specific instance of a user subscribing to a plan. It tracks the status of a specific user, their start date, and applied pricing.

### 1.2. Checkout

- **Order**: Represents a payment between two or more parties (an agreement between a buyer and a seller).
- **Payment**: Represents the actual settlement, often used in conjunction with the Orders API.

## 2. Stripe

TBD