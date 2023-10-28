Circuit breaker in Go
=============

A service in a microservices architecture typically makes calls to other services in order to gather data, and there's a possibility that the upstream service might be unavailable. In the event that temporary network problems or temporary unavailability are the root of the problem, the client support team may attempt to resolve the issue multiple times.

But there could also be other major issues, including a database failure or a slow-moving service. In certain situations, an excessive number of pointlessly repeated queries could cause cascade failures across the entire system. The circuit breaker, then, is the hero. It's a safeguard that lets you keep your service from processing too many requests in a short amount of time.

An RPC, to put it simply, is like drawing two straight lines between two services. A request is sent from service A to service B via one line, and the response is received back from service B to service A via the other. However, in order to put that "limiting" policy into practise, we require a middleman of some sort to determine whether or not to route a request to the intended destination.

This intermediary, wrapper, or proxyproxy (not a network proxy) either allows the connection or circuit between two services to be “Closed” or prevents one from contacting the other, thereby “Opening” the circuit.

## The following is the basic concept of a circuit breaker:

The circuit is by default in closed mode, which permits you to make free calls to the destination. It will stop allowing you to make further requests after a predetermined number of unsuccessful responses from the destination (a threshold, let's say 5), and the circuit will be regarded as open for a while (a backoff period, say 30 seconds). It enters a state known as Half-Open after that time. If the purpose of the subsequent request is to ascertain whether we will remain in an open state or enter a closed state.

The circuit will be closed if the request is approved; otherwise, we will revert to an open condition and have to wait for another back-off period.

![Schema of pub/sub pattern](image.png)


