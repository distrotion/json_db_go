package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"GqlIconTest/graph/generated"
	"GqlIconTest/graph/model"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func (r *mutationResolver) CallIcon(ctx context.Context, input model.IconName) (*model.ReturnIcon, error) {

	fmt.Println(input.Icon)
	file, _ := ioutil.ReadFile("pic_lib_B64.json")
	//fmt.Println(file)

	var result map[string]interface{}
	json.Unmarshal([]byte(file), &result)

	fmt.Println(reflect.TypeOf(result))
	//fmt.Println(reflect.TypeOf(result["icon-activity"]))

	iconout := result[input.Icon]

	//fmt.Println(reflect.TypeOf(iconout))

	var out *model.ReturnIcon
	dummyOut := model.ReturnIcon{
		Icon: iconout.(string),
	}
	out = &dummyOut
	return out, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
