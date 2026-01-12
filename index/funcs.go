package index

import (
	"fmt"

	"github.com/basebytes/elastic-go/client/entity"
	"github.com/mitchellh/mapstructure"
)

func Decode(input, output any) error {
	decoderConfig := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		DecodeHook:       decodeHook,
		Result:           output,
	}
	decoder, err := mapstructure.NewDecoder(decoderConfig)
	if err == nil {
		err = decoder.Decode(input)
	}
	return err
}

var decodeHook = mapstructure.ComposeDecodeHookFunc(
	entity.BucketsMapToSliceHookFunc(),
	entity.AnyToStringHookFunc(),
	entity.SettingItemHookFunc())

type SkipFieldFunc func(string) bool

type CollectFunc func(map[string]any) map[string]any

func ExtractAggResult(aggs *entity.Aggregations, lastField string, skip SkipFieldFunc, collect CollectFunc) (results []map[string]any, imprecise byte) {
	if lastField == "" {
		stats, _imprecise := extractStatistics(aggs, skip)
		if collect != nil {
			stats = collect(stats)
		}
		if len(stats) > 0 {
			results = append(results, stats)
			imprecise = _imprecise
		}
		return
	}
	for key, agg := range *aggs {
		if key == Item || (skip != nil && skip(key)) {
			for k, value := range agg.Other {
				if innerAggs := transToAggs(k, value); innerAggs != nil {
					return ExtractAggResult(innerAggs, lastField, skip, collect)
				}
			}
			continue
		}
		if agg.DocCountErrorUpperBound > 0 || agg.SumOtherDocCount > 0 {
			imprecise |= 1
		}
		if agg.Buckets != nil {
			for _, bucket := range *agg.Buckets {
				if key == lastField {
					stats, _imprecise := extractStatistics(&bucket.Aggs, skip)
					if collect != nil {
						stats = collect(stats)
					}
					if len(stats) > 0 {
						stats[key] = bucket.Key
						results = append(results, stats)
						imprecise |= _imprecise
					}
				} else {
					_results, _imprecise := ExtractAggResult(&bucket.Aggs, lastField, skip, collect)
					for _, result := range _results {
						result[key] = bucket.Key
					}
					results = append(results, _results...)
					imprecise |= _imprecise
				}
			}
		}
	}
	return
}

func extractStatistics(aggs *entity.Aggregations, skipField SkipFieldFunc) (map[string]any, byte) {
	var (
		result    = make(map[string]any)
		imprecise byte
	)
	for key, agg := range *aggs {
		if key == Item || (skipField != nil && skipField(key)) {
			for k, value := range agg.Other {
				if innerAggs := transToAggs(k, value); innerAggs != nil {
					res, _imprecise := extractStatistics(innerAggs, skipField)
					for _k, _v := range res {
						result[_k] = _v
					}
					imprecise |= _imprecise
				}
			}
			continue
		}
		if agg.SumOtherDocCount > 0 || agg.DocCountErrorUpperBound > 0 {
			imprecise |= 1
		}
		if agg.Buckets != nil {
			if res := extractResultFromBucket(agg.Buckets); len(res) > 0 {
				result[key] = res
			}
		} else if count := agg.Other["value"]; count != nil && fmt.Sprintf("%v", count) != zeroValue {
			result[key] = count
		} else if count = agg.Other["doc_count"]; count != nil && fmt.Sprintf("%v", count) != zeroValue {
			result[key] = count
		}
	}
	return result, imprecise
}

func extractResultFromBucket(buckets *[]*entity.BucketItem) (result []map[string]any) {
	for _, bucket := range *buckets {
		if bucket.DocCount > 0 {
			result = append(result, map[string]any{bucket.Key: bucket.DocCount})
		}
	}
	return
}

func transToAggs(key string, value any) *entity.Aggregations {
	if mapValue, OK := value.(map[string]any); OK {
		var result entity.AggregationsResult
		if err := Decode(mapValue, &result); err == nil {
			return &entity.Aggregations{key: &result}
		}
	}
	return nil
}

const zeroValue = "0"
