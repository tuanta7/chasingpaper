# Chasing Paper 🤑

Chasing Paper provides a comprehensive solution for managing subscriptions, handling payments, and tracking user billing information. The service integrates with popular payment providers to enable smooth subscription creation, renewal, and cancellation workflows.

## Dependencies

Chasing Paper relies on some popular Go packages for its codebase.

| Package                     | Description                                       |
|-----------------------------|---------------------------------------------------|
| go-chi/chi                  | Lightweight HTTP router for Go                    |
| go-playground/validator/v10 | Value validator for structs and individual fields |
| sqlc-dev/sqlc               | SQL code generator for type-safe database queries |
| jackc/pgx/v5                | Go driver and toolkit for PostgreSQL              |
| vektra/mockery              | Mock code generator for Go interfaces             |
| stretchr/testify            | Go testing toolkit with assertions and mocks      |
| ~~stripe/stripe-go~~        | ~~Stripe Go library for payment processing~~      |

## PayPal Integration

Chasing Paper leverages PayPal's REST APIs as the primary payment gateway.

- [PayPal Developer Sandbox](https://developer.paypal.com/dashboard/applications/sandbox)
- [PayPal REST APIs](https://developer.paypal.com/api/rest/current-resources/)
- [PayPal Postman](https://www.postman.com/paypal/paypal-public-api-workspace/collection/ujhlb45/paypal-apis?sideView=agentMode)

## Stripe Integration

TBD: Stripe integration will be added after the initial release.

## Getting Started

```shell
make run-server
```

## Support

For issues and feature requests, please open an issue on GitHub.
