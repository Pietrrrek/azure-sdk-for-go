//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armlogic_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/logic/armlogic"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_List.json
func ExampleIntegrationAccountMapsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewIntegrationAccountMapsClient().NewListPager("testResourceGroup", "testIntegrationAccount", &armlogic.IntegrationAccountMapsClientListOptions{Top: nil,
		Filter: nil,
	})
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.IntegrationAccountMapListResult = armlogic.IntegrationAccountMapListResult{
		// 	Value: []*armlogic.IntegrationAccountMap{
		// 		{
		// 			Name: to.Ptr("IntegrationAccountMap9943"),
		// 			Type: to.Ptr("Microsoft.Logic/integrationAccounts/maps"),
		// 			ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/maps/IntegrationAccountMap9943"),
		// 			Properties: &armlogic.IntegrationAccountMapProperties{
		// 				ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T18:34:32.390576Z"); return t}()),
		// 				ContentLink: &armlogic.ContentLink{
		// 					ContentHash: &armlogic.ContentHash{
		// 						Algorithm: to.Ptr("md5"),
		// 						Value: to.Ptr("A2avz/M0ov2FPI3+Je8vDw=="),
		// 					},
		// 					ContentSize: to.Ptr[int64](3056),
		// 					ContentVersion: to.Ptr("\"0x8D45CE3C6D23B4B\""),
		// 					URI: to.Ptr("<Uri>"),
		// 				},
		// 				CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T18:34:32.3902373Z"); return t}()),
		// 				MapType: to.Ptr(armlogic.MapTypeXslt),
		// 				Metadata: map[string]any{
		// 				},
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_Get.json
func ExampleIntegrationAccountMapsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountMapsClient().Get(ctx, "testResourceGroup", "testIntegrationAccount", "testMap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountMap = armlogic.IntegrationAccountMap{
	// 	Name: to.Ptr("testMap"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/maps"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/testResourceGroup/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/maps/testMap"),
	// 	Properties: &armlogic.IntegrationAccountMapProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T18:45:23.4137139Z"); return t}()),
	// 		ContentLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("A2avz/M0ov2FPI3+Je8vDw=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](3056),
	// 			ContentVersion: to.Ptr("\"0x8D45CE54B058881\""),
	// 			URI: to.Ptr("<Uri>"),
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-02-24T18:45:23.4129778Z"); return t}()),
	// 		MapType: to.Ptr(armlogic.MapTypeXslt),
	// 		Metadata: map[string]any{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_CreateOrUpdate.json
func ExampleIntegrationAccountMapsClient_CreateOrUpdate_createOrUpdateAMap() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountMapsClient().CreateOrUpdate(ctx, "testResourceGroup", "testIntegrationAccount", "testMap", armlogic.IntegrationAccountMap{
		Location: to.Ptr("westus"),
		Properties: &armlogic.IntegrationAccountMapProperties{
			Content:     to.Ptr("<?xml version=\"1.0\" encoding=\"UTF-16\"?>\r\n<xsl:stylesheet xmlns:xsl=\"http://www.w3.org/1999/XSL/Transform\" xmlns:msxsl=\"urn:schemas-microsoft-com:xslt\" xmlns:var=\"http://schemas.microsoft.com/BizTalk/2003/var\" exclude-result-prefixes=\"msxsl var s0 userCSharp\" version=\"1.0\" xmlns:ns0=\"http://BizTalk_Server_Project4.StringFunctoidsDestinationSchema\" xmlns:s0=\"http://BizTalk_Server_Project4.StringFunctoidsSourceSchema\" xmlns:userCSharp=\"http://schemas.microsoft.com/BizTalk/2003/userCSharp\">\r\n  <xsl:import href=\"http://btsfunctoids.blob.core.windows.net/functoids/functoids.xslt\" />\r\n  <xsl:output omit-xml-declaration=\"yes\" method=\"xml\" version=\"1.0\" />\r\n  <xsl:template match=\"/\">\r\n    <xsl:apply-templates select=\"/s0:Root\" />\r\n  </xsl:template>\r\n  <xsl:template match=\"/s0:Root\">\r\n    <xsl:variable name=\"var:v1\" select=\"userCSharp:StringFind(string(StringFindSource/text()) , &quot;SearchString&quot;)\" />\r\n    <xsl:variable name=\"var:v2\" select=\"userCSharp:StringLeft(string(StringLeftSource/text()) , &quot;2&quot;)\" />\r\n    <xsl:variable name=\"var:v3\" select=\"userCSharp:StringRight(string(StringRightSource/text()) , &quot;2&quot;)\" />\r\n    <xsl:variable name=\"var:v4\" select=\"userCSharp:StringUpperCase(string(UppercaseSource/text()))\" />\r\n    <xsl:variable name=\"var:v5\" select=\"userCSharp:StringLowerCase(string(LowercaseSource/text()))\" />\r\n    <xsl:variable name=\"var:v6\" select=\"userCSharp:StringSize(string(SizeSource/text()))\" />\r\n    <xsl:variable name=\"var:v7\" select=\"userCSharp:StringSubstring(string(StringExtractSource/text()) , &quot;0&quot; , &quot;2&quot;)\" />\r\n    <xsl:variable name=\"var:v8\" select=\"userCSharp:StringConcat(string(StringConcatSource/text()))\" />\r\n    <xsl:variable name=\"var:v9\" select=\"userCSharp:StringTrimLeft(string(StringLeftTrimSource/text()))\" />\r\n    <xsl:variable name=\"var:v10\" select=\"userCSharp:StringTrimRight(string(StringRightTrimSource/text()))\" />\r\n    <ns0:Root>\r\n      <StringFindDestination>\r\n        <xsl:value-of select=\"$var:v1\" />\r\n      </StringFindDestination>\r\n      <StringLeftDestination>\r\n        <xsl:value-of select=\"$var:v2\" />\r\n      </StringLeftDestination>\r\n      <StringRightDestination>\r\n        <xsl:value-of select=\"$var:v3\" />\r\n      </StringRightDestination>\r\n      <UppercaseDestination>\r\n        <xsl:value-of select=\"$var:v4\" />\r\n      </UppercaseDestination>\r\n      <LowercaseDestination>\r\n        <xsl:value-of select=\"$var:v5\" />\r\n      </LowercaseDestination>\r\n      <SizeDestination>\r\n        <xsl:value-of select=\"$var:v6\" />\r\n      </SizeDestination>\r\n      <StringExtractDestination>\r\n        <xsl:value-of select=\"$var:v7\" />\r\n      </StringExtractDestination>\r\n      <StringConcatDestination>\r\n        <xsl:value-of select=\"$var:v8\" />\r\n      </StringConcatDestination>\r\n      <StringLeftTrimDestination>\r\n        <xsl:value-of select=\"$var:v9\" />\r\n      </StringLeftTrimDestination>\r\n      <StringRightTrimDestination>\r\n        <xsl:value-of select=\"$var:v10\" />\r\n      </StringRightTrimDestination>\r\n    </ns0:Root>\r\n  </xsl:template>\r\n</xsl:stylesheet>"),
			ContentType: to.Ptr("application/xml"),
			MapType:     to.Ptr(armlogic.MapTypeXslt),
			Metadata:    map[string]any{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountMap = armlogic.IntegrationAccountMap{
	// 	Name: to.Ptr("IntegrationAccountMap291"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/maps"),
	// 	ID: to.Ptr("/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/<resourceGroup>/providers/Microsoft.Logic/integrationAccounts/<IntegrationAccount>/maps/testMap"),
	// 	Properties: &armlogic.IntegrationAccountMapProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T18:41:04.4088605Z"); return t}()),
	// 		ContentLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("A2avz/M0ov2FPI3+Je8vDw=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](3056),
	// 			ContentVersion: to.Ptr("\"0x8D464C057F22E5F\""),
	// 			URI: to.Ptr("<Uri>"),
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2017-03-06T18:41:03.7366103Z"); return t}()),
	// 		MapType: to.Ptr(armlogic.MapTypeXslt),
	// 		Metadata: map[string]any{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountLargeMaps_CreateOrUpdate.json
func ExampleIntegrationAccountMapsClient_CreateOrUpdate_createOrUpdateAMapLargerThan4Mb() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountMapsClient().CreateOrUpdate(ctx, "testResourceGroup", "testIntegrationAccount", "testMap", armlogic.IntegrationAccountMap{
		Location: to.Ptr("westus"),
		Properties: &armlogic.IntegrationAccountMapProperties{
			ContentLink: &armlogic.ContentLink{
				URI: to.Ptr("<blob-SAS-URL-for-map>"),
			},
			ContentType: to.Ptr("application/xml"),
			MapType:     to.Ptr(armlogic.MapTypeXslt),
			Metadata:    map[string]any{},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.IntegrationAccountMap = armlogic.IntegrationAccountMap{
	// 	Name: to.Ptr("testMap"),
	// 	Type: to.Ptr("Microsoft.Logic/integrationAccounts/maps"),
	// 	ID: to.Ptr("/subscriptions/<Azure-subscription-ID>/resourceGroups/refresh/providers/Microsoft.Logic/integrationAccounts/testIntegrationAccount/maps/testMap"),
	// 	Properties: &armlogic.IntegrationAccountMapProperties{
	// 		ChangedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-09T07:40:10.2906473Z"); return t}()),
	// 		ContentLink: &armlogic.ContentLink{
	// 			ContentHash: &armlogic.ContentHash{
	// 				Algorithm: to.Ptr("md5"),
	// 				Value: to.Ptr("GxQRrFCYoyH58kMyu34ISg=="),
	// 			},
	// 			ContentSize: to.Ptr[int64](7888419),
	// 			ContentVersion: to.Ptr("\"0x8D9EB9F6691E7A2\""),
	// 			URI: to.Ptr("<URI>"),
	// 		},
	// 		CreatedTime: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2022-02-09T07:40:10.2863459Z"); return t}()),
	// 		MapType: to.Ptr(armlogic.MapTypeXslt),
	// 		Metadata: map[string]any{
	// 		},
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_Delete.json
func ExampleIntegrationAccountMapsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	_, err = clientFactory.NewIntegrationAccountMapsClient().Delete(ctx, "testResourceGroup", "testIntegrationAccount", "testMap", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/d55b8005f05b040b852c15e74a0f3e36494a15e1/specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_ListContentCallbackUrl.json
func ExampleIntegrationAccountMapsClient_ListContentCallbackURL() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armlogic.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewIntegrationAccountMapsClient().ListContentCallbackURL(ctx, "testResourceGroup", "testIntegrationAccount", "testMap", armlogic.GetCallbackURLParameters{
		KeyType:  to.Ptr(armlogic.KeyTypePrimary),
		NotAfter: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-04-19T16:00:00Z"); return t }()),
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.WorkflowTriggerCallbackURL = armlogic.WorkflowTriggerCallbackURL{
	// 	Method: to.Ptr("GET"),
	// 	BasePath: to.Ptr("https://prod-00.brazilus.logic.azure.com/integrationAccounts/0fdabc3a76514ca48dede71c73d9fe97/maps/testmap/contents/Value"),
	// 	Queries: &armlogic.WorkflowTriggerListCallbackURLQueries{
	// 		APIVersion: to.Ptr("2015-08-01-preview"),
	// 	},
	// 	Value: to.Ptr("https://prod-00.westus.logic.azure.com:443/integrationAccounts/0fdabc3a76514ca48dede71c73d9fe97/maps/testMap/contents/Value?api-version=2015-08-01-preview&sp=%2Fmaps%2Ftestmap%2Fread&sv=1.0&sig=VK_mbQPTHTa3ezhsrI8IctckwjlL3GdJmroQH_baYj4"),
	// }
}
