Pub/Sub in Golang
=============

"Publish-subscribe," or "pub/sub," is a messaging pattern that separates systems and the communications that occur between them. Publishers and Subscribers are the two categories of entities in a pub/sub system, as the name suggests. While Subscribers are entities that receive messages by subscribing to a single topic, Publishers are entities that create messages and publish them to a specific topic.

Publishers do not communicate directly with subscribers in a pub/sub system. Rather, the messages are sent via a broker, an intermediary tasked with obtaining messages from publishers and sending them to the relevant subscribers. As a result, messages can be received by subscribers without requiring them to know who the publishers are, and vice versa. When it comes to Golang, goroutines and channels can be used to create this pub/sub paradigm.

[Schema of pub/sub pattern](image.png)


