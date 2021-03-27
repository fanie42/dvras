package dvras

// SequenceConflictError - When write to event store fails due to conflict
// of event sequence and stream sequence: Usually leads to a retry of the
// command with updated version of the aggregate.
// type SequenceConflictError struct {
//     Have uint64
//     Want uint64
// }

// // Error TODO
// func (err SequenceConflictError) Error() string {
//     return fmt.Sprintf(
//         "event sequence conflict: have version %d, but wanted version %d\n",
//         err.Have, err.Want,
//     )
// }
