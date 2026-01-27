# statusphere

This repository contains an aggressively-simplified Go implementation of the ATProtocol [statusphere](https://github.com/bluesky-social/statusphere-example-app) sample app.

It differs from [indigo](https://github.com/bluesky-social/indigo)-based ATProtocol applications in several ways:

- Authentication is handled by the [IO](https://agent.io/posts/io) proxy. This moves all OAuth signin-handling, request signing, retry, and request routing out of the app, greatly simplifying the application.
- Instead of consuming the firehose, the app uses [Jetstream](https://docs.bsky.app/blog/jetstream) to watch for new updates. This dramatically reduces bandwidth requirements and eliminates the need for CBOR and [ugly circular dependencies](https://github.com/bluesky-social/indigo/issues/931) in Lexicon code generation.
- The generated lexicon code is produced by [slink](https://github.com/agentio/slink), a fresh XRPC code generator that omits CBOR overhead, flattens the package space, produces type-safe unions, and, when run with a [list of Lexicon IDs](xrpc.json), generates only the code needed to use those methods and records.

For demonstration purposes, generated code is checked into this repo, but that's neither necessary nor recommended. See the `xrpc` make target for an easy way to regenerate `xrpc` handlers at build time.
