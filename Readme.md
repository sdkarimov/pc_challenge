# PaulCamper offline challenge

## Application process

To provide your solution in the most suitable way please use following process:
* fork a git repository
* iterate on the solution and commit changes step by step (one commit per logical change set)
with descriptive one line comments
* open a pull request (merge request) to this repo in a 72 hours
(all commits withing 72 hours after challenge has been shared with you are considered as an accepted solution)

You will find 3 parts of the task below in this Readme file. Parts differs in complexity.
They can be treated as separate subtasks and implemented one by one starting from top to bottom.
Partial solutions also accepted.


## Domain

We are working on the service which uses external 3rd-party translation provider API.
We have translation provider API interface (TranslatorAPI) and implementation (translatorStub) for testing purposes.
We know that translation API calls:
* can take long time
* sometimes fails
* costs us money
* result may change sometimes but business decided we can cache them for at least N hours


## Task

* Create a service that will properly handle external translation API with:
1. retry requests N times with exponential back off before failing with an error
2. cache request results in the storage to avoid charges for the same queries (simplest in-memory storage is enough)
3. deduplicate simultaneous queries for the same parameters to avoid charges for same query burst
* Cover new functionality with tests. This is to the same extent important as a solution code itself.


## Source code

translator.go and main.go should not be modified. Please use service.go and any new files for the solution.
