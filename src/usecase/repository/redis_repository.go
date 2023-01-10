package repository

type RedisRepository interface {
	GetCookie(accountID string) error
}
