# statusphere-io

This repository contains an aggressively-simplified Go implementation of the ATProtocol [statusphere](https://github.com/bluesky-social/statusphere-example-app) sample app.

It differs from [indigo](https://github.com/bluesky-social/indigo)-based ATProtocol applications in several ways:

- Authentication is handled by the [IO](https://agent.io/io) proxy. This moves all OAuth signin-handling, request signing, retry, and request routing out of the app, greatly simplifying the xrpc client.
- Instead of consuming the firehose, the app uses [Jetstream](https://docs.bsky.app/blog/jetstream) to watch for new updates. This dramatically reduces bandwidth requirements and eliminates the need for CBOR and [ugly circular dependencies](https://github.com/bluesky-social/indigo/issues/931) in Lexicon code generation.
- The generated lexicon code is hand-edited output of the indigo [lexgen](https://github.com/bluesky-social/indigo/tree/main/cmd/lexgen) (with an eye toward future automatic generation) that removes all CBOR overhead and replaces the hard-coded dependency on a concrete xrpc client with an xrpc interface, allowing multiple clients to be easily used with the generated lexicon code.
- Two small xrpc clients are included: one that makes unauthenticated (aka "anonymous") requests to Bluesky servers and one that makes authenticated requests through IO.
