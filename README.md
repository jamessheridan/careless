# careless.go

This is a daft dabble in Golang.

It's a Twitter bot which looks for a particular grammar peeve, specifically when people mean to say that they "couldn't care less" but instead they state, inexplicably, that they actually "could care less". This is patently absurd, of course - so careless takes care of letting the offending Twitter user know of their faux-pas and replies with a correction.

Uses boltdb to store state (so we don't reply to the same Tweet more than once).

The config file should be self-documenting. Throw your API keys in here, and it should work.
