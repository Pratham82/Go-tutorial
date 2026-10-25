package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var pointerInitialVotes *int

	pointerInitialVotes = &initialVotes

	return pointerInitialVotes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var result ElectionResult

	result.Name = candidateName
	result.Votes = votes

	return &result
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {

	candidateName := result.Name
	votes := result.Votes

	return fmt.Sprintf("%s (%d)", candidateName, votes)

}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for k, v := range results {
		if k == candidate {
			results[k] = results[k] - 1
		}
		fmt.Println(k, v)
	}
}
