package bug

import (
	"context"
	"entgo.io/bug/ent/leaguecertificatetype"
	"fmt"
	"net"
	"strconv"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
)

func TestBugSQLite(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, client *ent.Client) {
	ctx := context.Background()
	client.User.Delete().ExecX(ctx)
	client.User.Create().SetName("Ariel").SetAge(30).ExecX(ctx)
	if n := client.User.Query().CountX(ctx); n != 1 {
		t.Errorf("unexpected number of users: %d", n)
	}

	client.LeagueCertificateType.Delete().ExecX(ctx)
	client.League.Delete().ExecX(ctx)
	client.CertificateType.Delete().ExecX(ctx)

	l1, err := client.League.Create().SetName("league_1").Save(ctx)
	assert.Equal(t, nil, err)

	c1, err := client.CertificateType.Create().SetName("cert_1").Save(ctx)
	assert.Equal(t, nil, err)

	c2, err := client.CertificateType.Create().SetName("cert_2").Save(ctx)
	assert.Equal(t, nil, err)

	_, err = client.LeagueCertificateType.Create().SetLeagueID(l1.ID).SetCertificateTypeID(c1.ID).Save(ctx)
	assert.Equal(t, nil, err)

	_, err = client.LeagueCertificateType.Create().SetLeagueID(l1.ID).SetCertificateTypeID(c2.ID).Save(ctx)
	assert.Equal(t, nil, err)

	// using previous m2m implementation
	result, err := client.Debug().League.Query().
		WithLeagueCertificateType(func(q *ent.LeagueCertificateTypeQuery) {
			// still wrongly selects non-existent league_certificate_type.id column
			q.Select(leaguecertificatetype.FieldLeagueID)
			q.Select(leaguecertificatetype.FieldCertificateTypeID)
			q.WithCertificates()
		}).
		All(ctx)
	assert.Equal(t, nil, err)
	t.Log(result)
	/*
		outputs non-existent id column: SELECT DISTINCT `league_certificate_type`.`id`

		2022/09/29 10:06:31 driver.Query: query=SELECT DISTINCT `leagues`.`id`, `leagues`.`name` FROM `leagues` args=[]
		2022/09/29 10:06:31 driver.Query: query=SELECT DISTINCT `league_certificate_type`.`id`, `league_certificate_type`.`league_id`, `league_certificate_type`.`certificate_type_id` FROM `league_certificate_type` WHERE `league_id` IN (?) args=[1]
		2022/09/29 10:06:31 driver.Query: query=SELECT DISTINCT `certificate_types`.`id`, `certificate_types`.`name` FROM `certificate_types` WHERE `certificate_types`.`id` IN (?, ?) args=[1 2]

	*/

	// using through
	//result, err := client.Debug().League.Query().WithCertTypes().All(ctx)
	//assert.Equal(t, nil, err)
	//t.Log(result)
}
