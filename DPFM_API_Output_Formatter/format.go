package dpfm_api_output_formatter

import (
	"database/sql"
	"fmt"
)

func ConvertToGlobalRegion(rows *sql.Rows) (*GlobalRegion, error) {
	defer rows.Close()
	globalRegion := GlobalRegion{}
	i := 0

	for rows.Next() {
		i++
		err := rows.Scan(
			&globalRegion.GlobalRegion,
			&globalRegion.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &globalRegion, err
		}

	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &globalRegion, nil
	}

	return &globalRegion, nil
}
