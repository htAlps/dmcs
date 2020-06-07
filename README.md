# Distributed Minimal Composable Superimposed λ-System 

## Overview
In this two repositories DMCS-λ and DMCS-π we develop distributed systems for the purpose of exploring aspects like: security weaknesses or strengths, elegance, performance, etc.
The λ version (λ-calculus) is a process/services implementation while the π version (π-calculus) is a distributed (containerized) implementation using the usual suspects.
They are esentially the same abstract distributed system with two distinct implementations. 
We're Following a general guideline: Superposition of Minimal Orthogonal Compositions which to us means the following:

- write minimal code that is sufficient to make a point and no more (Occam's Razor)

- design modules to do one thing, be composable, and independent (orthogonal) of each other within the same layer

- don't compose when you can superimpose components


## Rant
We want to avoid superfuous bell and whistles or undecomposable abstractions


## Components
We want to explore the following in order of appearance

- Arch Linux: for better OS component control and package management

- Go (Golang): an obvious choice and because it's a favorite languages with Haskell, Rust, and C.

- Docker: to model scalability

- gRPC: to explore performance 

- etcd: to explore basic clustering 

- Swarm: to explore simple scalability

- Consul: to explore simple security 

- Vault: to explore improvements to secure key management

- Prometheus: to 


## Branching on Subjects 
We want to explore a number of subject areas while preserving key steps that exemplify an issue of interest. 
The Git platform seems ideal for a concrete representation of this. 

## Exposing Decomposition and Evolution 
Often we are presented a system in its current magestic (and sometimes indescipherable) state and wonder how it got there, 
i.e., what are its essential components, when was each component introduced, etc. 

We want to explore security and best practices as we evolve and branch the systems and try to establish a cannonical partial order.  


    # DMCS - Distributed Minimal Composable Superimposed System 

    ## Functional Evolution and Branching Tree


       DMCS - (λ: Services Based)
        │
        ├── ReST
        │
        ├── ProtoBuf
        │
        ├── GRPC
        │  
        ├───────────────> DMCSP - (π: Docker-Hashicorp)
        │                   │                                     
        ├── TLS             ├── Docker                            
        │                   │                                     
        ├── SystemD         ├───────────────> DMCSK - (π: Kube-Stack)
        │                   │                   │  
        ├── ETCD            ├── Swarm           ├── Kubernetes    
        │                   │                   │       
        ├──                 ├── Consul          ├── Istio    
        │                   │                   │       
        └─>                 ├── Vault           ├── Prometheus      
                            │                   │       
                            ├── Prometheus      ├──     
                            │                   │       
                            ├──                 └─> 
                            │                                     
                            └─>                           

    ### Legend:
       λ System  (λ-calculus) is a process/services implementation 
       π Systems (π-calculus) are distributed (containerized) implementations





## NON-Functional Components 

        Environments:
            Development: Local Arch-Linux Servers Inside  
            Integration & Testing: AZ / AWS / GCP Platforms 
            Production: Same Platforms (may not be needed due to the experimental nature of the project)

        DevOps (Pipelines): AZ DevOps



