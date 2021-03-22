package dvras

import "fmt"

// SequenceConflictError - When write to event store fails due to conflict
// of event sequence and stream sequence: Usually leads to a retry of the
// command with updated version of the aggregate.
type SequenceConflictError struct {
    have     uint64
    expected uint64
}

// Error TODO
func (err *SequenceConflictError) Error() string {
    return fmt.Sprintf(
        "event sequence conflict: have version %d, but expected version %d",
        err.have, err.expected,
    )
}
