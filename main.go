package main

import (
        "fmt"
        "os"

        "gorm.io/driver/mysql"
        "gorm.io/gorm"
)

func ping() (bool, error) {
        dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=true&loc=Local",
                os.Getenv("DB_USER"),
                os.Getenv("DB_PASSWORD"),
                os.Getenv("DB_HOST"),
                os.Getenv("DB_PORT"))
        conn, err := gorm.Open(mysql.Open(dsn))
        if err != nil {
                return false, err
        }

        db, err := conn.DB()
        if err != nil {
                return false, err
        }
        if err := db.Ping(); err != nil {
                return false, err
        }

        return true, nil
}

func main() {
        fmt.Println("OK")
}