package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"os"
)

var (
	subsID string
	groupName string
)

var (
	ctx = context.Background()
)

func main() {
	subsID = os.Getenv("SUBSCRIPTION_ID")
	if subsID == "" {
		panic("Empty SUBSCRIPTION_ID")
	}
	groupName = os.Getenv("GROUP_NAME")
	if groupName == "" {
		panic("Empty GROUP_NAME")
	}
	if authorizer, err := auth.NewAuthorizerFromCLI(); err != nil {
		panic(err)
	} else {
		for iterator, err := listInResourceGroup(groupName, authorizer);
		iterator.NotDone();
		err = iterator.NextWithContext(ctx) {
			if err != nil {
				panic(err)
			}
			fmt.Println(*iterator.Value().ID)
		}
	}
}

func listInResourceGroup(groupName string, authorizer autorest.Authorizer) (resources.ListResultIterator, error) {
	client := resources.NewClient(subsID)
	client.Authorizer = authorizer
	return client.ListByResourceGroupComplete(ctx, groupName, "(resourceType eq 'Microsoft.Compute/virtualMachines' or resourceType eq 'Microsoft.Compute/virtualMachineScaleSets')", "", nil)
}
