package main

import (
	"net/http"
	"time"

	"example.com/m/database"
	"example.com/m/muxes"
	_ "github.com/lib/pq"
)

func main() {

	db, err := database.StartDB()
	checkErr(err)
	db.Exec(`CREATE TABLE IF NOT EXISTS USERS(
        ID INT PRIMARY KEY,
        NAME TEXT NOT NULL,
        SURNAME TEXT NOT NULL
      );
      CREATE SEQUENCE public.users_id_seq NO MINVALUE NO MAXVALUE NO CYCLE;
      ALTER TABLE public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq');
      ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;`)
	defer db.Close()

	s := &http.Server{
		Addr:           ":3000",
		Handler:        muxes.Serve(db),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
