# sap-api-integrations-purchase-requisition-reads
sap-api-integrations-purchase-requisition-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で 購買依頼データを取得するマイクロサービスです。    
sap-api-integrations-purchase-requisition-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-purchase-requisition-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_PURCHASEREQ_PROCESS_SRV_0001/overview   

## 動作環境  
sap-api-integrations-purchase-requisition-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用
sap-api-integrations-purchase-requisition-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-purchase-requisition-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PURCHASEREQ_PROCESS_SRV_0001/overview    
* APIサービス名(=baseURL): API_PURCHASEREQ_PROCESS_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-purchase-requisition-reads には、次の API をコールするためのリソースが含まれています。  

* A_PurchaseRequisitionHeader（購買依頼 - ヘッダ）
* A_PurchaseRequisitionItem（購買依頼 - 明細）
* A_PurReqAddDelivery（購買依頼 - 納入先住所）
* A_PurReqnAcctAssgmt（購買依頼 - 勘定設定）

## API への 値入力条件 の 初期値
sap-api-integrations-purchase-requisition-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.PurchaseRequisition.PurchaseRequisition（購買依頼）
* inoutSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchaseRequisitionItem（購買依頼明細）
* inoutSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchasingDocument（購買発注）
* inoutSDC.PurchaseRequisition.PurchaseRequisitionItem.PurchasingDocumentItem（購買発注明細）
