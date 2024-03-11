# Notes

## Difference between an entity and an aggregate
(might be overly simplistic)
- An entity is the thing you store in the database
- An aggregate is the thing you work with inside the application and (sort of) end up showing to users

__Example: `Subscription`__
A `Feed` is an _entity_, it has an identity and is stored in the database. The application's user can _subscribe_ to a feed. A `Subscription` is __not__ an _entity_. A `subscription` is an _aggregate_ that contains a reference to the `User` and the `Feed` entities.
