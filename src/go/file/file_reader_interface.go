package file

type FileReaderInterface interface {
    Read(filename string) (string, error)
}
