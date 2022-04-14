package internal

// MultipartTodoData is returned by utils.ExtractTodoMultipartFileAndData
// The utils function is called by handlers when creating and updating to-dos.
// FileData is byte content of the file uploaded as multipart/form-data
type MultipartTodoData struct {
	FileData []byte
	JSONData []byte
}
