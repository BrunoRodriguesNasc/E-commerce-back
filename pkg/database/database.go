package database

import (
	"database/sql"
	"log"
	"time"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	connStr := "host=db port=5432 user=postgres password=postgres dbname=ecommerce sslmode=disable"
	
	// Tenta conectar com retries
	var db *sql.DB
	var err error
	
	for i := 0; i < 30; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Tentativa %d: Erro ao abrir conexão: %v", i+1, err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = db.Ping()
		if err == nil {
			log.Println("Conectado ao banco de dados com sucesso!")
			return db
		}

		log.Printf("Tentativa %d: Erro ao pingar banco: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Não foi possível conectar ao banco de dados após várias tentativas")
	return nil
}

func Close(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatal(err)
	}
}


