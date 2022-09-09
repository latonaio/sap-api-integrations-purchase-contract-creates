# sap-api-integrations-purchase-contract-creates  
sap-api-integrations-purchase-contract-creates は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 購買基本契約データ を登録するマイクロサービスです。  
sap-api-integrations-purchase-contract-creates には、サンプルのAPI Json フォーマットが含まれています。  
sap-api-integrations-purchase-contract-creates は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。  
https://api.sap.com/api/OP_API_PURCHASECONTRACT_PROCESS_SRV/overview  

## 動作環境  
sap-api-integrations-purchase-contract-creates は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）   
・ AION のリソース （推奨)   
・ OS: LinuxOS （必須）   
・ CPU: ARM/AMD/Intel（いずれか必須）  

## クラウド環境での利用
sap-api-integrations-purchase-contract-creates は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。


## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchase-contract-creates が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PURCHASECONTRACT_PROCESS_SRV/overview  
* APIサービス名(=baseURL): API_PURCHASECONTRACT_PROCESS_SRV  

## 本レポジトリ に 含まれる API名
sap-api-integrations-creates には、次の API をコールするためのリソースが含まれています。  

* A_PurchaseContract（購買基本契約 - ヘッダ）
* A_PurchaseContractItem（購買基本契約 - 明細）  
* A_PurContrItemCondition（購買基本契約 - 明細価格条件）
* A_PurCtrAddress（購買基本契約 - 明細住所）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"Header" が指定されています。    
  
```
	"api_schema": "SAPPurchaseContractCreates",
	"accepter": ["Header"],
	"purchase_contract_no": "",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "SAPPurchaseContractCreates",
	"accepter": ["All"],
	"purchase_contract_no": "",
	"deleted": false
```
## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncPostPurchaseContract(
	header            *requests.Header,
	item              *requests.Item,
	accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(header)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(item)
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
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 購買基本契約 の ヘッダデータ が登録された結果の JSON の例です。  
以下の項目のうち、"PurchaseContract" ～ "to_PurchaseContractItem" は、/SAP_API_Output_Formatter/type.go 内 の Type Header {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-purchase-contract-reads/SAP_API_Caller/caller.go#L58",
	"function": "sap-api-integrations-purchase-contract-reads/SAP_API_Caller.(*SAPAPICaller).Header",
	"level": "INFO",
	"message": [
		{
			"PurchaseContract": "4600000001",
			"PurchaseContractType": "MK",
			"CompanyCode": "1710",
			"PurchasingDocumentDeletionCode": "",
			"CreationDate": "/Date(1471305600000)/",
			"Supplier": "USSU-VSF10",
			"PurchasingOrganization": "1710",
			"PurchasingGroup": "001",
			"PaymentTerms": "0001",
			"NetPaymentDays": "0",
			"DocumentCurrency": "USD",
			"ExchangeRate": "1.00000",
			"ValidityStartDate": "/Date(1467331200000)/",
			"ValidityEndDate": "/Date(1469923200000)/",
			"SupplierRespSalesPersonName": "",
			"SupplierPhoneNumber": "",
			"IncotermsClassification": "EXW",
			"PurchaseContractTargetAmount": "4000.00",
			"InvoicingParty": "",
			"ReleaseCode": "",
			"LastChangeDateTime": "",
			"PurchasingProcessingStatus": "02",
			"PurchasingProcessingStatusName": "Active",
			"PurgContractIsInPreparation": false,
			"to_PurchaseContractItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PURCHASECONTRACT_PROCESS_SRV/A_PurchaseContract('4600000001')/to_PurchaseContractItem"
		}
	],
	"time": "2022-01-28T11:39:14.908297+09:00"
```
