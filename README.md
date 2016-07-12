# careless.go

This is a daft dabble in Golang.

It's a Twitter bot which looks for a particular grammar peeve, specifically when people misuse "I couldn't care less" as "I could care less", and replies with a correction.

Uses boltdb to store state (so we don't reply to the same Tweet more than once).

The config file should be self-documenting. Throw your API keys in here, and it should work.
