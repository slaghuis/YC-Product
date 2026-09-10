package services

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func StrToUint(str string) (uint, error) {
  // Convert string to uint64 first
	// Base 10, 64-bit size
	u64, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}

	// Explicitly convert the uint64 to a uint
	// Note: The size of uint is implementation defined (either 32 or 64 bits)
	// You may lose data if the number is too large for the target uint type
	// on a 32-bit system
	u := uint(u64)
  return u, err
}

func StrToInt(str string, def int) (int) {
  num, err := strconv.Atoi(str)
  if err != nil {
    return def
  }
  return num
}

func UintToStr(u uint) (string) {
  // Cast 'u' to uint64 and specify base 10
  return strconv.FormatUint(uint64(u), 10)
}

func IntToStr(i int) (string) {
  return strconv.Itoa(i)
}

func LogJson(data any) error {
  // Marshal the struct with indentation (empty prefix, tab indent)
  jsonData, err := json.MarshalIndent(data, "", "  ")
  if err != nil {
    return err
  }

  // Print the pretty-printed JSON string
  fmt.Println(string(jsonData))

  return nil
}


func Paginate(data []int, page, limit int) []int {
    if limit <= 0 || page <= 0 {
        return nil
    }

    start := (page - 1) * limit
    if start >= len(data) {
        // No records for this page
        return nil
    }

    end := start + limit
    if end > len(data) {
        end = len(data) // clamp to slice length
    }

    return data[start:end]
}
