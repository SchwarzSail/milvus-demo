package milvus

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
	"github.com/pkg/errors"
	"milvus-demo/pkg/constants"
)

func CreateCollection(ctx context.Context) (ok bool, err error) {
	client := MClient.milvus
	exist, err := client.HasCollection(ctx, constants.CollectionName)
	if err != nil {
		return false, errors.WithMessage(err, "milvus error")
	}
	if exist {
		return true, nil
	}
	schema := &entity.Schema{
		CollectionName: constants.CollectionName,
		Fields: []*entity.Field{
			{
				Name:       "id",
				DataType:   entity.FieldTypeInt64,
				PrimaryKey: true,
				AutoID:     false,
			},
			{
				Name:       "vector",
				DataType:   entity.FieldTypeFloatVector,
				TypeParams: map[string]string{"dim": "512"},
			},
		},
	}
	err = client.CreateCollection(ctx, schema, 2)
	if err != nil {
		return false, errors.WithMessage(err, "CreateCollection failed")
	}
	return false, nil
}

func CreateIndex(ctx context.Context) (err error) {
	client := MClient.milvus
	idx, err := entity.NewIndexIvfFlat(
		entity.L2,
		1024,
	)
	if err != nil {
		return errors.WithMessage(err, "Create Idx failed")
	}
	err = client.CreateIndex(
		ctx,
		constants.CollectionName,
		"vector",
		idx,
		false,
	)
	if err != nil {
		return errors.WithMessage(err, "CreateIndex failed")
	}
	return nil
}

func InsertVector(ctx context.Context, vector []float32, id int64) (err error) {
	client := MClient.milvus
	entities := []entity.Column{
		entity.NewColumnInt64("id", []int64{id}),
		entity.NewColumnFloatVector("vector", 512, [][]float32{vector}),
	}
	if _, err := client.Insert(ctx, constants.CollectionName, "", entities...); err != nil {
		return errors.WithMessage(err, "insert failed")
	}
	return nil
}

func Search(ctx context.Context, vector []float32) (ids []int64, err error) {
	client := MClient.milvus
	//先加载一下Collection
	err = client.LoadCollection(ctx, constants.CollectionName, false)
	if err != nil {
		return nil, errors.WithMessage(err, "load collection failed")
	}
	//定义聚类中心，越高说明搜索越精准，但是时间消耗更多
	sp, _ := entity.NewIndexIvfFlatSearchParam(
		10,
	)
	res, err := client.Search(
		ctx,
		constants.CollectionName,
		[]string{},
		"",
		[]string{"id"},
		[]entity.Vector{entity.FloatVector(vector)},
		"vector",
		entity.L2,
		10,
		sp,
	)
	if err != nil {
		return nil, errors.WithMessage(err, "search failed")
	}
	ids = make([]int64, 0, len(res))
	for _, v := range res {
		for i := 0; i < v.ResultCount; i++ {
			id, err := v.IDs.GetAsInt64(i)
			if err != nil {
				return nil, errors.WithMessage(err, "get id failed")
			}
			ids = append(ids, id)
		}
	}
	return ids, nil
}
