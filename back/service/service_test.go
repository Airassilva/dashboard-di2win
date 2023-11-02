package service

import "testing"

func TestUser(t *testing.T) {
	rep := RepositorioPostgre{
		connStr: "user=postgres dbname=database sslmode=disable",
	}

	t.Run("find one user", func(t *testing.T) {
		expected := User{id: 3, name: "Caio", segment: "imobiliaria"}
		got := rep.findUser(3)

		if expected != got {
			t.Errorf("Expected %v, got %v", expected, got)
		}
	})

	t.Run("could not find one user", func(t *testing.T) {
		_, err := rep.findUser(10)

		if err != nil {
			t.Error("Should've failed but didn't")
		}
	})
}
