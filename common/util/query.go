package util

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/rohanchauhan02/clean/common/schemas"
)

type Value struct {
	Key string
	Val interface{}
}

type Statement struct {
	Name         string
	Table        string
	Queries      []string
	Conditions   []string
	Values       []string
	Using        []string
	MappedValues []Value
}

// Build query to be used to update the value of database
// @config: schemas of UpdateDataConfig
// @data: string of json encoded data of result set from the database
// @response: string of json encoded data of http response,
// @state: State of the result set status
// @statusProcess: Current status of the process
// @outputPreparedStatement: Output query using PreparedStatement or not
// It will return either query or error
func UpdateQueryBuilder(config schemas.UpdateDataConfig, data string, response string, state string, statusProcess string, outputPreparedStatement bool) (queries []string, err error) {
	var Statements []Statement
	var ProcessedMappings []*schemas.UpdateDataMapping

	if config.TriggerStatus != nil {
		for triggerKey, triggerVal := range config.TriggerStatus {
			// Search in the mappings based on the value and name
			for _, mapping := range config.Mappings {
				// Check mapping name and trigger is same
				// Also check the status
				if (strings.EqualFold(mapping.Name, triggerVal) && strings.EqualFold(triggerKey, state)) && strings.EqualFold(mapping.Status, statusProcess) {
					// append to process mappings
					ProcessedMappings = append(ProcessedMappings, mapping)
				}
			}
		}

		if len(ProcessedMappings) > 0 {
			for _, mapping := range ProcessedMappings {
				var value interface{}
				var Statement Statement

				Statement.Table = mapping.Table

				// Get the column mapping and get each of the value
				for k, v := range mapping.ColumnMappings {
					var exprVal string

					if strings.EqualFold(v.Operation, "GJSON") || strings.EqualFold(v.Operation, "JSON") {
						value = GetValueJson(response, v.Value, v.Type)
						if _, k := value.(string); k {
							value = MysqlRealEscapeString(fmt.Sprintf("%s", value))
						}
						if strings.EqualFold(v.Operation, "JSON") {
							if !strings.EqualFold(v.Type, "string") {
								keySplit := strings.Split(v.Key, ".")
								exprVal = fmt.Sprintf("%s = JSON_SET(%s, '$.%s', %v)", keySplit[0], keySplit[0], v.Value, value)
							} else {
								keySplit := strings.Split(v.Key, ".")
								exprVal = fmt.Sprintf("%s = JSON_SET(%s, '$.%s', '%s')", keySplit[0], keySplit[0], v.Value, value)
							}
						} else if strings.EqualFold(v.Operation, "GJSON") {
							if !strings.EqualFold(v.Type, "string") {
								exprVal = fmt.Sprintf("%s = %v", v.Key, value)
							} else {
								exprVal = fmt.Sprintf("%s = '%s'", v.Key, value)
							}
						}
					} else if strings.HasPrefix(strings.ToUpper(v.Operation), "REGEX") {
						value = GetValueJson(response, v.Value, "string")

						ops := strings.Split(v.Operation, "[")
						matchIndex := 0
						if len(ops) > 1 {
							opsIndex := strings.Split(ops[1], "]")[0]
							i, err := strconv.Atoi(opsIndex)
							if err == nil {
								matchIndex = i
							}
						}

						r, err := regexp.Compile(v.Type)
						if err == nil {
							subMatches := r.FindStringSubmatch(value.(string))
							if len(subMatches) > 0 && len(subMatches) > matchIndex {
								matchedValue := subMatches[matchIndex]
								exprVal = fmt.Sprintf("%s = '%s'", v.Key, MysqlRealEscapeString(matchedValue))
							}
						}
					} else {
						if !strings.EqualFold(v.Type, "string") {
							exprVal = fmt.Sprintf("%s = %v", v.Key, v.Value)
						} else {
							exprVal = fmt.Sprintf("%s = '%s'", v.Key, MysqlRealEscapeString(v.Value))
						}
					}

					if outputPreparedStatement {
						val := value
						Statement.Values = append(Statement.Values, fmt.Sprintf("%s = ?", v.Key))
						if strings.EqualFold(v.Operation, "JSON") {
							keySplit := strings.Split(v.Key, ".")
							val = fmt.Sprintf("JSON_SET(%s, '$.%s', %v)", keySplit[0], v.Value, value)
						} else {
							val = MysqlRealEscapeString(v.Value)
						}
						Statement.MappedValues = append(Statement.MappedValues, Value{
							Key: fmt.Sprintf("Values%v", k),
							Val: val,
						})
					} else {
						Statement.Values = append(Statement.Values, exprVal)
					}
				}

				// Get the conditions
				for k, v := range mapping.Conditions {
					var exprVal string
					value = GetValueJson(data, v.Value, "")
					typeVal := fmt.Sprintf("%T", value)

					if !strings.EqualFold(typeVal, "string") {
						exprVal = fmt.Sprintf("%v", value)
					} else {
						value = MysqlRealEscapeString(fmt.Sprintf("%s", value))
						exprVal = fmt.Sprintf("'%s'", value)
					}

					if outputPreparedStatement {
						Statement.Conditions = append(Statement.Conditions, v.Key)
						Statement.MappedValues = append(Statement.MappedValues, Value{
							Key: fmt.Sprintf("Cond%v", k),
							Val: value,
						})
					} else {
						condition := strings.Replace(v.Key, "?", exprVal, -1)
						Statement.Conditions = append(Statement.Conditions, condition)
					}
				}
				// Add statement to the statements
				// for query building
				Statements = append(Statements, Statement)
			}
		}
	}

	if len(Statements) < 1 {
		return nil, fmt.Errorf("there is no statement to be processed")
	}

	// Loop all the statement
	for _, Statament := range Statements {
		if outputPreparedStatement {
			rgx := regexp.MustCompile("[^a-zA-Z0-9]+")
			// Set the statement name
			Statament.Name = fmt.Sprintf("stmntupdate%s", rgx.ReplaceAllString(Statament.Table, ""))
			// Set the prepared statement
			updateStatement := fmt.Sprintf("UPDATE %s SET %s WHERE %s", Statament.Table, strings.Join(Statament.Values, ", "), strings.Join(Statament.Conditions, " AND "))
			// Prepare the Statament
			Statament.Queries = append(Statament.Queries, fmt.Sprintf("PREPARE %s FROM '%s';", Statament.Name, updateStatement))
			// Loop all the mapped values
			for _, v := range Statament.MappedValues {
				typeVal := fmt.Sprintf("%T", v.Val)
				if !strings.EqualFold(typeVal, "string") {
					Statament.Queries = append(Statament.Queries, fmt.Sprintf("SET @%s = %v;", v.Key, v.Val))
				} else {
					Statament.Queries = append(Statament.Queries, fmt.Sprintf("SET @%s = \"%s\";", v.Key, v.Val))
				}
				Statament.Using = append(Statament.Using, fmt.Sprintf("@%s", v.Key))
			}
			// Execute
			Statament.Queries = append(Statament.Queries, fmt.Sprintf("EXECUTE %s USING %s;", Statament.Name, strings.Join(Statament.Using, ", ")))
			// After Execute deallocate Statament
			Statament.Queries = append(Statament.Queries, fmt.Sprintf("DEALLOCATE PREPARE %s;", Statament.Name))

			queries = append(queries, strings.Join(Statament.Queries, ""))
		} else {
			query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", Statament.Table, strings.Join(Statament.Values, ", "), strings.Join(Statament.Conditions, " AND "))
			queries = append(queries, query)
		}
	}

	return queries, nil
}
