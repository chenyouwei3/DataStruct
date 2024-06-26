package main

import (
	"database/sql"
	"log"
	"strings"
	"time"
)

type User struct {
	UserId string `json:"userid"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type VIP struct {
	UserId      string
	VIPStart    time.Time
	VIPEnd      time.Time
	PaidVIPDays int
	FreeVIPDays int
}

func main() {
	db, err := sql.Open("sqlite3", "./users_vip.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 查询users表中的所有记录
	rows, err := db.Query("SELECT userid, email, phone FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var Users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Email, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		if !isValidEmail(user.Email) || !isValidPhone(user.Phone) {
			Users = append(Users, user)
		}
	}
	for _, user := range Users {
		rows, err := db.Query("SELECT vip_start_date, vip_end_date, paid FROM vip_records WHERE userid = ?", user.UserId)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

	}
}
func isValidEmail(email string) bool {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return false
	}
	local, domain := parts[0], parts[1]
	if strings.Contains(local, "..") || strings.Contains(domain, "..") || !strings.HasSuffix(domain, ".com") || strings.Count(domain, ".") != 1 {
		return false
	}
	return true
}

func isValidPhone(phone string) bool {
	expectedFormat := "+86-###-####-####"
	if len(phone) != len(expectedFormat) {
		return false
	}
	for i := 0; i < len(expectedFormat); i++ {
		if expectedFormat[i] != '#' && phone[i] != expectedFormat[i] {
			return false
		}
	}
	return true
}
