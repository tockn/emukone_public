package repository

type HashID interface {
	Decode(string) (int64, error)
	Encode(int64) string
}
