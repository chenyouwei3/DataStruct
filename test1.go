package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	UserId string `json:"userid"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

type VIPRecord struct {
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

	var invalidUsers []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.UserId, &user.Email, &user.Phone)
		if err != nil {
			log.Fatal(err)
		}
		if !isValidEmail(user.Email) || !isValidPhone(user.Phone) {
			invalidUsers = append(invalidUsers, user)
		}
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// 输出存储格式不正确的用户列表
	fmt.Println("Invalid Users:")
	for _, user := range invalidUsers {
		fmt.Printf("User ID: %s, Email: %s, Phone: %s\n", user.UserId, user.Email, user.Phone)
	}

	// 计算并输出VIP历史信息，并按要求格式写入文件
	var vipRecords []VIPRecord
	for _, user := range invalidUsers {
		// 查询该用户的VIP购买记录
		rows, err := db.Query("SELECT vip_start_date, vip_end_date, paid FROM vip_records WHERE userid = ?", user.UserId)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var vipStart, vipEnd string
			var paid int
			err := rows.Scan(&vipStart, &vipEnd, &paid)
			if err != nil {
				log.Fatal(err)
			}

			// 解析日期字符串为时间对象
			startTime, err := time.Parse("2006-01-02", vipStart)
			if err != nil {
				log.Fatal(err)
			}
			endTime, err := time.Parse("2006-01-02", vipEnd)
			if err != nil {
				log.Fatal(err)
			}

			// 计算VIP天数
			vipDays := int(endTime.Sub(startTime).Hours() / 24)

			// 根据付费状态添加到VIP记录列表中
			vipRecord := VIPRecord{
				UserId:      user.UserId,
				VIPStart:    startTime,
				VIPEnd:      endTime,
				PaidVIPDays: 0,
				FreeVIPDays: 0,
			}
			if paid == 1 {
				vipRecord.PaidVIPDays = vipDays
			} else {
				vipRecord.FreeVIPDays = vipDays
			}

			vipRecords = append(vipRecords, vipRecord)
		}
	}

	// 将结果按照指定格式输出并写入本地文件
	outputFile, err := os.Create("invalid_user_vip_summary.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	for _, record := range vipRecords {
		output := fmt.Sprintf("UserID: %s, Total VIP Count: %d, Free VIP Days: %d, Paid VIP Days: %d\n",
			record.UserId, record.FreeVIPDays+record.PaidVIPDays, record.FreeVIPDays, record.PaidVIPDays)
		_, err := outputFile.WriteString(output)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("VIP信息已成功写入文件 invalid_user_vip_summary.txt")
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
