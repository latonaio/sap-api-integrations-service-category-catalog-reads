# sap-api-integrations-service-category-catalog-reads  
sap-api-integrations-service-category-catalog-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API サービスカテゴリーカタログデータを取得するマイクロサービスです。  
sap-api-integrations-service-category-catalog-reads には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-service-category-catalog-reads は、オンプレミス版である（＝クラウド版ではない）SAPC4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/serviceissuecategorycatalogue/overview  

## 動作環境
sap-api-integrations-service-category-catalog-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。   
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。   
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須） 

## クラウド環境での利用  
sap-api-integrations-service-category-catalog-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-service-category-catalog-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/serviceissuecategorycatalogue/overview    
* APIサービス名(=baseURL): c4codataapi

## 本レポジトリ に 含まれる API名
sap-api-integrations-service-category-catalog-reads には、次の API をコールするためのリソースが含まれています。  

* ServiceCategoryCatalogCollection（サービスカテゴリーカタログ - サービスカテゴリーカタログ）  

## API への 値入力条件 の 初期値
sap-api-integrations-service-category-catalog-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ServiceCategoryCatalogCollection.ID（ID）  


## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"ServiceCategoryCatalogCollection" が指定されています。    
  
```
	"api_schema": "ServiceCategoryCatalog",
	"accepter": ["ServiceCategoryCatalogCollection"],
	"service_category_catalog_code": "SCC_1",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "ServiceCategoryCatalog",
	"accepter": ["All"],
	"service_category_catalog_code": "SCC_1",
	"deleted": false
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetServiceCategoryCatalog(iD string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "ServiceCategoryCatalogCollection":
			func() {
				c.ServiceCategoryCatalogCollection(iD)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-sap](https://github.com/latonaio/golang-logging-library-for-sap) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP サービスカテゴリーカタログ の サービスカテゴリーカタログデータ が取得された結果の JSON の例です。  
以下の項目のうち、"ObjectID" ～ "EntityLastChangedOn" は、/SAP_API_Output_Formatter/type.go 内 の Type ServiceCategoryCatalogCollection {} による出力結果です。"cursor" ～ "time"は、golang-logging-library-for-sap による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-service-category-catalog-reads/SAP_API_Caller/caller.go#L53",
	"function": "sap-api-integrations-service-category-catalog-reads/SAP_API_Caller.(*SAPAPICaller).ServiceCategoryCatalogCollection",
	"level": "INFO",
	"message": [
		{
			"ObjectID": "00163E03A0701ED28A846E6DDBA46A34",
			"ETag": "2013-06-24T15:26:36+09:00",
			"ID": "SCC_1",
			"VersionID": "1",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2013-06-24T15:27:26+09:00",
			"StartDateTime": "2012-11-23T18:48:52+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2013-06-24T15:26:36+09:00"
		},
		{
			"ObjectID": "00163E04B6021EE2B6E7F5F5EE02E4F7",
			"ETag": "2013-06-25T08:47:56+09:00",
			"ID": "SCC_1",
			"VersionID": "2",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2013-06-25T08:27:10+09:00",
			"StartDateTime": "2013-06-24T15:27:27+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2013-06-25T08:47:56+09:00"
		},
		{
			"ObjectID": "00163E04B6021EE2B7A18DA73A11521D",
			"ETag": "2013-06-26T03:39:34+09:00",
			"ID": "SCC_1",
			"VersionID": "3",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2013-06-26T03:40:29+09:00",
			"StartDateTime": "2013-06-25T08:27:11+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2013-06-26T03:39:34+09:00"
		},
		{
			"ObjectID": "00163E04B6021EE2B7B885E531F8AC75",
			"ETag": "2014-01-27T18:47:01+09:00",
			"ID": "SCC_1",
			"VersionID": "4",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2014-01-27T18:47:54+09:00",
			"StartDateTime": "2013-06-26T03:40:30+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2014-01-27T18:47:01+09:00"
		},
		{
			"ObjectID": "00163E0664D31ED3A1E6C0D9FE64C946",
			"ETag": "2014-01-28T11:21:58+09:00",
			"ID": "SCC_1",
			"VersionID": "5",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2014-01-28T11:22:29+09:00",
			"StartDateTime": "2014-01-27T18:47:55+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2014-01-28T11:21:58+09:00"
		},
		{
			"ObjectID": "00163E0664D31ED3A1F8132D168C85A8",
			"ETag": "2014-02-13T16:27:17+09:00",
			"ID": "SCC_1",
			"VersionID": "6",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2014-02-13T16:28:04+09:00",
			"StartDateTime": "2014-01-28T11:22:30+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2014-02-13T16:27:17+09:00"
		},
		{
			"ObjectID": "00163E0664D31EE3A58FF007AA794D07",
			"ETag": "2015-10-05T15:47:38+09:00",
			"ID": "SCC_1",
			"VersionID": "7",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2015-10-05T15:48:20+09:00",
			"StartDateTime": "2014-02-13T16:28:05+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2015-10-05T15:47:38+09:00"
		},
		{
			"ObjectID": "00163E0C6CDB1ED59AE56F47A543E8B0",
			"ETag": "2016-02-12T20:50:20+09:00",
			"ID": "SCC_1",
			"VersionID": "8",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2016-02-12T20:51:15+09:00",
			"StartDateTime": "2015-10-05T15:48:21+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2016-02-12T20:50:20+09:00"
		},
		{
			"ObjectID": "00163E0C6CDB1ED89586E833310C1E11",
			"ETag": "2019-07-30T19:48:01+09:00",
			"ID": "SCC_1",
			"VersionID": "13",
			"LifeCycleStatusCode": "2",
			"LifeCycleStatusCodeText": "Released",
			"EndDateTime": "2019-07-30T13:59:59+09:00",
			"StartDateTime": "2018-05-11T14:00:00+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2019-07-30T19:48:01+09:00"
		},
		{
			"ObjectID": "00163E10F31E1EE5B4AEBD8498F484DD",
			"ETag": "2016-02-12T20:15:58+09:00",
			"ID": "SCC_1",
			"VersionID": "9",
			"LifeCycleStatusCode": "1",
			"LifeCycleStatusCodeText": "In Preparation",
			"EndDateTime": "2015-02-13T20:00:00+09:00",
			"StartDateTime": "2015-02-13T09:00:00+09:00",
			"ServiceCategoryCatalogueName": "Almika SCC",
			"ServiceCategoryCatalogueNamelanguageCode": "EN",
			"ServiceCategoryCatalogueNamelanguageCodeText": "English",
			"EntityLastChangedOn": "2016-02-12T20:15:58+09:00"
		}
	],
	"time": "2022-07-27T13:22:22+09:00"
}

```