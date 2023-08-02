# Artifact Search Specification

The Artifact Search Specification provides the definition of Artifact Search. Artifact Search is the capability to discover Artifacts and their relationships regardless of where those Artifacts reside.

## Characteristics
A system that implements the Artifact Search Specification has the following characteristics:

1. The Artifact Search specification leverages Open Containers Initiative (OCI) for the handling of artifact information. 

2. Artifacts are described with [semantic triples](https://en.wikipedia.org/wiki/Semantic_triple) and addressed by one or many of the component parts of their triple. This means that if an object, predicate, or subject is queried for; its resource locater is resolved and returned in the query response. More than one artifact can be referenced in many ways, one way notable here is by assigning a predicate to the subject of triple. 

3. While some semantics are reserved by the core system, the semantics used to interact with artifacts can be dynamically registered with the system. 

4. Artifact Interaction interfaces are linked to their semantic definitions. These interfaces act as drivers for clients in their retrieval of artifacts. They can also provide runtime instruction for further processing of artifacts.

5. Artifacts are addressed by resource address, which can be any value to include a content address, namespaced locations, or a predicate. If a client does not know how to retrieve an artifact from a given resource address, it can optionally dynamically load that resource type's driver. 

6. Name resolution is provided by a GraphQL API.


