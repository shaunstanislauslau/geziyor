package exporter

import (
	"encoding/csv"
	"fmt"
	"github.com/geziyor/geziyor"
	"log"
	"os"
	"reflect"
	"sync"
)

// CSVExporter exports response data as CSV streaming file
type CSVExporter struct {
	FileName string

	once   sync.Once
	mut    sync.Mutex
	writer *csv.Writer
}

func (e *CSVExporter) Export(response *geziyor.Response) {

	// Default filename
	if e.FileName == "" {
		e.FileName = "out.csv"
	}

	// Create file
	e.once.Do(func() {
		newFile, err := os.OpenFile(e.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "output file creation error: %v", err)
			return
		}
		e.writer = csv.NewWriter(newFile)
	})

	// Export data as responses came
	for res := range response.Exports {
		var values []string

		// Detect type and extract CSV values
		val := reflect.ValueOf(res)
		switch val.Kind() {

		case reflect.Slice:
			for i := 0; i < val.Len(); i++ {
				values = append(values, fmt.Sprint(val.Index(i)))
			}

			//case reflect.Map:
			//	iter := val.MapRange()
			//	for iter.Next() {
			//		values = append(values, fmt.Sprint(iter.Value()))
			//	}
		}

		// Write to file
		e.mut.Lock()
		if err := e.writer.Write(values); err != nil {
			log.Printf("CSV writing error on exporter: %v\n", err)
		}
		e.mut.Unlock()
	}

	e.mut.Lock()
	e.writer.Flush()
	e.mut.Unlock()
}
