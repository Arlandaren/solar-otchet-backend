package utils

import (
	"encoding/csv"
	"server/models"
	"strconv"
	"strings"
	"time"
)

func ConvertToCSV(stats *[]models.Stats) ([]byte, error) {
	var csvData [][]string

	// Write header
	header := []string{"ID", "Timestamp", "OrgName", "SrcIP", "SrcPort", "DstIP", "DstPort", "PacketsCount", "BytesCount"}
	csvData = append(csvData, header)

	// Iterate through stats and write each record
	for _, s := range *stats {
		record := []string{
			strconv.FormatUint(uint64(s.ID), 10),
			s.Timestamp.Format(time.RFC3339),
			s.OrgName,
			s.SrcIP,
			strconv.Itoa(s.SrcPort),
			s.DstIP,
			strconv.Itoa(s.DstPort),
			strconv.Itoa(s.PacketsCount),
			strconv.Itoa(s.BytesCount),
		}
		csvData = append(csvData, record)
	}

	// Create a buffer to write CSV data
	var csvBuffer strings.Builder
	writer := csv.NewWriter(&csvBuffer)

	// Write CSV data
	err := writer.WriteAll(csvData)
	if err != nil {
		return nil, err
	}

	// Return the CSV data as []byte
	return []byte(csvBuffer.String()), nil
}
