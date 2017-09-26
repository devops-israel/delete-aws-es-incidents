// Copyright 2012-present Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package elastic

// AvgAggregation is a single-value metrics aggregation that computes
// the average of numeric values that are extracted from the
// aggregated documents. These values can be extracted either from
// specific numeric fields in the documents, or be generated by
// a provided script.
// See: https://www.elastic.co/guide/en/elasticsearch/reference/5.2/search-aggregations-metrics-avg-aggregation.html
type AvgAggregation struct {
	field           string
	script          *Script
	format          string
	subAggregations map[string]Aggregation
	meta            map[string]interface{}
}

func NewAvgAggregation() *AvgAggregation {
	return &AvgAggregation{
		subAggregations: make(map[string]Aggregation),
	}
}

func (a *AvgAggregation) Field(field string) *AvgAggregation {
	a.field = field
	return a
}

func (a *AvgAggregation) Script(script *Script) *AvgAggregation {
	a.script = script
	return a
}

func (a *AvgAggregation) Format(format string) *AvgAggregation {
	a.format = format
	return a
}

func (a *AvgAggregation) SubAggregation(name string, subAggregation Aggregation) *AvgAggregation {
	a.subAggregations[name] = subAggregation
	return a
}

// Meta sets the meta data to be included in the aggregation response.
func (a *AvgAggregation) Meta(metaData map[string]interface{}) *AvgAggregation {
	a.meta = metaData
	return a
}

func (a *AvgAggregation) Source() (interface{}, error) {
	// Example:
	//	{
	//    "aggs" : {
	//      "avg_grade" : { "avg" : { "field" : "grade" } }
	//    }
	//	}
	// This method returns only the { "avg" : { "field" : "grade" } } part.

	source := make(map[string]interface{})
	opts := make(map[string]interface{})
	source["avg"] = opts

	// ValuesSourceAggregationBuilder
	if a.field != "" {
		opts["field"] = a.field
	}
	if a.script != nil {
		src, err := a.script.Source()
		if err != nil {
			return nil, err
		}
		opts["script"] = src
	}

	if a.format != "" {
		opts["format"] = a.format
	}

	// AggregationBuilder (SubAggregations)
	if len(a.subAggregations) > 0 {
		aggsMap := make(map[string]interface{})
		source["aggregations"] = aggsMap
		for name, aggregate := range a.subAggregations {
			src, err := aggregate.Source()
			if err != nil {
				return nil, err
			}
			aggsMap[name] = src
		}
	}

	// Add Meta data if available
	if len(a.meta) > 0 {
		source["meta"] = a.meta
	}

	return source, nil
}
