package utils

import (
    "encoding/json"
)

// serializeRules converts []uint into a string for storage
func SerializeRules(rules []uint) string {
    b, _ := json.Marshal(rules)
    return string(b)
}

// parseRules converts stored string back into []uint
func ParseRules(s string) []uint {
    var rules []uint
    _ = json.Unmarshal([]byte(s), &rules)
    return rules
}
