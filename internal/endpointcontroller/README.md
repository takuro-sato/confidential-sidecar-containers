# Endpoint Controllers

This package provides [controller layer](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html#interface-adapters) for API endpoints of any protocol (http, gRPC, etc.).

We want to provide same endpoints for multiple protocols, but still we want to make them have consistent interface as much as possible.
By having a rule that we need to use this controller for the same use case, we can reduce the chance to have inconsistent interface for the same use case.