package db

import (
	"database/sql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

// DB 커넥션 관련 설정입니다.
const (
	dbDriver = "postgres"
)

// sqlc로 생성된 트랜잭션을 가져와서 관련 테스트를 초기화하고 관리합니다.
var testQueries *Queries

func TestMain(m *testing.M) {

	// env 설정
	err := godotenv.Load("../../.env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	dbSource := os.Getenv("DB_SOURCE")

	// database/sql 라이브러리를 통해 DB 커넥션을 생성
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// 커넥션이 생성되면 해당 커넥션을 SQLC 드라이버와 연동함.
	testQueries = New(conn)

	// Test 시작
	os.Exit(m.Run())

	/*
		이외에도 runtime.GoEixt()과 같은 함수를 사용헤서 defer르 무시하지 않고 수행한 뒤 테스트를 종료하는 방식을 사용할 수도 있음을 알아두자.
	*/
}
