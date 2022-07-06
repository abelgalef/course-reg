package storage

import "testing"

func TestStorage(t *testing.T) {
	// CHECK DB HAS BEEN PROPERLY INSTANTIATED BY COUNTING THE NUMBER OF TABLES IN THE DATABASE
	db := NewDatabaseConnection(false)

	want := 9   // THE NUMBER OF TABLES NEEDED IN THE DATABASE
	var got int // THE NUMBER OF TABLES WE GOT
	if err := db.Connection.Raw("SELECT count(*) AS TOTAL_DB_NUM FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = ? ;", "coursereg").Scan(&got).Error; err != nil {
		t.Fatalf("Could not execute query: %v", err)
	}

	if want != got {
		t.Fatalf("SQLite database table number mismatch; wanted %d tables got %d", want, got)
	}

	// Create MySQL Database
	db = NewDatabaseConnection(true)
	got = 0

	if err := db.Connection.Raw("SELECT count(*) FROM sqlite_master AS tables WHERE type='table'").Scan(&got).Error; err != nil {
		t.Fatalf("Could not execute query: %v", err)
	}

	if want != got {
		t.Fatalf("MySQL database table number mismatch; wanted %d tables got %d", want, got)
	}
}
