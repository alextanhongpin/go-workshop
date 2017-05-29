// package httputil

// type Error interface {
// 	error
// 	Status() int
// }

// type StatusError struct {
// 	Code int
// 	Err  error
// }

// // Allow StatusError to satisfy the error interface
// func (se StatusError) Error() string {
// 	return se.Err.Error()
// }

// func (se StatusError) Status() int {
// 	return se.Code
// }

// switch e := err(.type) {
// case Error:
// 	http.Error(w, e.Error(), e.Status())
// default:
// 	//
// }
