package dpfm_api_caller

import (
	dpfm_api_input_reader "data-platform-api-global-region-deletes-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-global-region-deletes-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"strings"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) GlobalRegion(
	input *dpfm_api_input_reader.SDC,
	log *logger.Logger,
) *dpfm_api_output_formatter.GlobalRegion {

	where := strings.Join([]string{
		fmt.Sprintf("WHERE globalRegion.GlobalRegion = \"%s\ ", input.GlobalRegion.GlobalRegion),
	}, "")

	rows, err := c.db.Query(
		`SELECT 
    	globalRegion.GlobalRegion
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_global_region_global_region_data as globalRegion 
		` + where + ` ;`)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToGlobalRegion(rows)
	if err != nil {
		log.Error("%+v", err)
		return nil
	}

	return data
}
