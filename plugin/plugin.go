package plugin

import (
	"gopkg.in/sensorbee/sensorbee.v0/bql"
	"gopkg.in/sensorbee/sensorbee.v0/bql/udf"
//	"github.com/yukimori/udfs"
)

func init() {
	udf.MustRegisterGlobalUDF("my_inc", udf.MustConvertGeneric(udfs.Inc))
}
