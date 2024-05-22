package data

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// randomString генерирует случайную строку заданной длиной и содержащей только буквы латинского алфавита
func randomString(length int) string {
	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = alphabet[rand.Intn(len(alphabet))]
	}

	return string(bytes)
}

// randomInt генерирует случайное целое число в диапазоне [min, max]
func randomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// randomPhone генерирует случайный телефонный номер из 10 цифр.
func randomPhone() string {
	phoneStr := "+7"
	phone := randomInt(9000000000, 9999999999)
	phoneStr += strconv.Itoa(phone)
	return phoneStr
}
// randomEmail генерирует случайный адрес электронной почты
func randomEmail() string {
	domains := []string{"example.com", "test.com", "mail.com", "inbox.com"}
	return fmt.Sprintf("%s@%s", randomString(10), domains[rand.Intn(len(domains))])
}

// randomTimeгенерирует случайное время в формате "2006-01-02T15:04:05Z"
func randomTime() time.Time {
	return time.Now()
}

// randomUIID генерирует случайный UUID
func randomUIID() uuid.UUID{
	return uuid.New()
}