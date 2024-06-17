package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/common/model"
)

func (p *DeviceState) getMultipleMetricsAverage(ctx context.Context, metrics []string, deviceNo string) (map[string]float64, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	duration := now.Sub(startOfDay)
	// 构造查询
	query := ""
	for i, metric := range metrics {
		if i != 0 {
			query += " or "
		}
		query += fmt.Sprintf(`avg_over_time(%s{exported_instance=&#126;"%s"}[%ss])`, metric, deviceNo, int(duration.Seconds()))
	}
	result, _, err := p.api.Query(ctx, query, now)
	if err != nil {
		return nil, err
	}
	vec, ok := result.(model.Vector)
	if !ok {
		return nil, errors.New("数据类型有误:" + result.String())
	}
	// 将结果存入map中
	results := make(map[string]float64)
	for _, sample := range vec {
		metricName := string(sample.Metric["name"])
		results[metricName] = float64(sample.Value)
	}
	return results, nil
}
