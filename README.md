# dmcs-l
Distributed Minimal Composable Superposition λ-System 

OVERVIEW
In this two repositories DMCS-λ and DMCS-π we develop distributed systems for the purpose of exploring aspects like: security weaknesses or strengths, elegance, performance, etc.
The λ version (λ-calculus) is a process/services implementation while the π version (π-calculus) is a distributed (containerized) implementation using the usual suspects.
We're Following a general guideline: Superposition of Minimal Orthogonal Compositions which to us means the following 4 things:
we are going to write minimal code that is sufficient to make a point and no more (Occam's Razor)
code modules will generally do one thing, be composable and independent of each other in the same layer

RANT
I'm tired of demos that are entire systems with bell and whistles and are already mounted on an undecomposable abstraction.


COMPONENTS
We want to explore the following in order of appearance
• Go:     because it's one of my favorite languages with Haskell, Rust, and Ruby.
• Docker: to Model π-systems
• gRPC:   for speed
• Swarm:  for simple scalability
• Consul: for simple security (could move to Kube-Istio later)
• Vault:  to explore how it improves secure key management

