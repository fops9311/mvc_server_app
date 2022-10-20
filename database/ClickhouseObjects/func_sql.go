package clickhouseobjects

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

func readDbSamples(object_id string) (result dbSamples) {
	log.Printf("[debug][clickhouse][readDbSamples]...%s", object_id)
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return
	}
	defer conn.Close()
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}), clickhouse.WithProfileInfo(func(p *clickhouse.ProfileInfo) {
		fmt.Println("profile info: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			log.Printf("[debug][clickhouse][readDbSamples]Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return
	}
	const limit = 2000
	result = make(dbSamples, 0)
	query := fmt.Sprintf("SELECT object_id,timestamp,value FROM example WHERE object_id LIKE ? ORDER BY object_id,timestamp DESC limit %d by object_id", limit)
	log.Printf("[debug][clickhouse][readDbSamples]query %s\n", query)
	rows, err := conn.Query(ctx,
		query,
		object_id+"/%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var (
		read_object_id string
		read_timestamp int64
		read_value     float32
	)
	for rows.Next() {
		if err := rows.Scan(&read_object_id, &read_timestamp, &read_value); err != nil {
			log.Fatal(err)
		}
		var newSample = dbSample{object_id: read_object_id, timestamp: read_timestamp, value: read_value}
		result = append(result, newSample)
	}
	log.Printf("[debug][clickhouse][readDbSamples]result len %d\n", len(result))
	return result

}

func dbWtireSamples(samples dbSamples) (err error) {
	log.Printf("[debug][clickhouse][dbWtireSamples] start writing %d samples", len(samples))
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:19000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
	})
	if err != nil {
		return err
	}
	defer conn.Close()
	ctx := clickhouse.Context(context.Background(), clickhouse.WithSettings(clickhouse.Settings{
		"max_block_size": 10,
	}), clickhouse.WithProgress(func(p *clickhouse.Progress) {
		fmt.Println("progress: ", p)
	}), clickhouse.WithProfileInfo(func(p *clickhouse.ProfileInfo) {
		fmt.Println("profile info: ", p)
	}))
	if err := conn.Ping(ctx); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("Catch exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		}
		return err
	}

	batch, err := conn.PrepareBatch(ctx, "INSERT INTO example (object_id,timestamp,value)")
	if err != nil {
		return err
	}
	for _, v := range samples {
		if err := batch.Append(v.object_id+"/", v.timestamp, v.value); err != nil {
			return err
		}
	}
	if err := batch.Send(); err != nil {
		return err
	}

	log.Printf("[debug][clickhouse][dbWtireSamples]added %d samples", len(samples))

	return nil
}
