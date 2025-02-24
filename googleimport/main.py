from organisationparser import OrganizationParser, SCOPES
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from export import Exporter
from organizationexport import OrganizationExporter
import os
import logging


def pidori(spred_id, lost_data, log_file):
#     SAMPLE_SPREADSHEET_ID = "1PyUB06AV8JGpPj8Z9Cw_NZqlVXCZi-xRtSG7i7M3cvs" = spred_id
#     LOST_DATA_JSON = 'BBC.json' = lost_data
#     LOG_FILE = 'BBC.log' = log_file
    SAMPLE_SPREADSHEET_ID = spred_id
    LOST_DATA_JSON = lost_data
    LOG_FILE = log_file
    creds = None
    # The file token.json stores the user's access and refresh tokens, and is
    # created automatically when the authorization flow completes for the first
    # time.
    if os.path.exists("token.json"):
        creds = Credentials.from_authorized_user_file("token.json", SCOPES)
    # If there are no (valid) credentials available, let the user log in.
    if not creds or not creds.valid:
        if creds and creds.expired and creds.refresh_token:
            creds.refresh(Request())
        else:
            flow = InstalledAppFlow.from_client_secrets_file(
            "creds.json", SCOPES)
            creds = flow.run_local_server(port=0)
    # Save the credentials for the next run
        with open("token.json", "w") as token:
            token.write(creds.to_json())


    l = logging.getLogger("OrganizationParser")
    l.setLevel(logging.INFO)
    logging.basicConfig(level=logging.DEBUG, filename=LOG_FILE, filemode="w")
    # logging.basicConfig(level=logging.WARNING, filename="warning.log", filemode="w")

    org = OrganizationParser(SAMPLE_SPREADSHEET_ID, creds, l)
    org.ParseAndDownload()
    orgas = org.GetOrganisation()

    e = Exporter("http://localhost", 5000, "TestHeadMaster", "12345678")

    oe = OrganizationExporter(l, orgas, e, LOST_DATA_JSON)
    oe.StartExport()


data = [
# {
# "spred_id": "1PV5aoeHyZ1wjFtLolPE55RkH-WmcTeNPPrY7FQ8dbr8",
# "lost_data":"БКП.json",
# "log_file": "БКП.log"
# },
{
"spred_id": "1BvAHep1Qnha5Cl4h8yJ4YS8OfsWHRL0R6SgZu0yJhYE",
"lost_data":"БФЛ.json",
"log_file": "БФЛ.log"
},
{
"spred_id": "1CO78_OiEEmj4GG4LSc4t67r_GIpXT8cqgCfJKVns6HY",
"lost_data":"ВЦ.json",
"log_file": "ВЦ.log"
},
# {
# "spred_id": "1jDldCvvX59uWP7F7eg2TxilcLP_h_NA8vN9vychlmmg",
# "lost_data":"Вектор.json",
# "log_file": "Вектор.log"
# },
# {
# "spred_id": "1BDFfd-3kWQszl1fZ0nmcoZhuwUi2IcdmqQRkc215wbs",
# "lost_data":"Кафедра_Юмора.json",
# "log_file": "Кафедра_Юмора.log"
# },
# {
# "spred_id": "1MGBk3kBpOPKMI8tk5w9YwjwriA02KfkbSQJ-NGEkmME",
# "lost_data":"КвизON.json",
# "log_file": "КвизON.log"
# },
# {
# "spred_id": "1tbxEUTY85cZqFlMTNmPYIwy9vWKPpJI36KyaOvVyAw8",
# "lost_data":"Киношки.json",
# "log_file": "Киношки.log"
# },
# {
# "spred_id": "1zzxhVZLxPvl-sWnpvIU5sQOQTKsx70ya7UDPCI83vCE",
# "lost_data":"Литквартирник.json",
# "log_file": "Литквартирник.log"
# },
{
"spred_id": "1fWG4DWPtoWHJk-l_Fp34dtDPrj-NXHd63iECu7IVZsk",
"lost_data":"Внешние_коммуникации.json",
"log_file": "Внешние_коммуникации.log"
},
{
"spred_id": "1nXZ6OKiJW0Ps2oKfY77HphIkXTxn03EzuPJIS7XUUOs",
"lost_data":"Внутренние_коммуникации.json",
"log_file": "Внутренние_коммуникации.log"
},
# {
# "spred_id": "1Q3LIlGXNFmWhFzRjKChfzg7yL55iDjuIRr0R1LYLJEU",
# "lost_data":"Проектный.json",
# "log_file": "Проектный.log"
# },
# {
# "spred_id": "1pwhc8Jzo1lu78w3XX5RG2S9RIwKdkYMOdTJyxjYHw6U",
# "lost_data":"Секретариат.json",
# "log_file": "Секретариат.log"
# },
{
"spred_id": "1J_RtMnHaUBMtyswhgw0eluO7ohnbM818brOOgUI-iPA",
"lost_data":"ЮР.json",
"log_file": "ЮР.log"
},

{
"spred_id": "1xGJOxKfDQqxh7i7Hsc9ISTjkdujUJuDfks_OP5SbM2I",
"lost_data":"№10.json",
"log_file": "№10.log"
},
# {
# "spred_id": "1-FV0bVGlbmx6j9pgiqCHcQfg4zDXcd84-nTBW-OTMME",
# "lost_data":"№11.json",
# "log_file": "№11.log"
# },
{
"spred_id": "1liyOfA3H8RtxCMDcDO_67UGj-PdjxjmG3CCZCCj4Ejg",
"lost_data":"АК.json",
"log_file": "АК.log"
},
{
"spred_id": "1GcnQmd4OHMYRXa6ku5J34qWkITfM9CWgOLFvNzmz7bA",
"lost_data":"ИБМ.json",
"log_file": "ИБМ.log"
},
{
"spred_id": "1_fy4EderTZOVEq1Sv3eV_o3WbAl0YviKTPW1pFeqUwE",
"lost_data":"ИУ.json",
"log_file": "ИУ.log"
},
{
"spred_id": "1LxmcWkumhtCtdL_PmbROOPhmm0D7MAugU40Sz3TGL9g",
"lost_data":"МТ.json",
"log_file": "МТ.log"
},
{
"spred_id": "1HD99KLlb5rAund_pmS4U5sz3LX-4pToJZJKf6GQUXvY",
"lost_data":"РК.json",
"log_file": "РК.log"
},
# {
# "spred_id": "1xcf2aZaDKzeQzzgtgYx6HOQndt6WqvmqnFGF2JeaZas",
# "lost_data":"РЛМ.json",
# "log_file": "РЛМ.log"
# },
# {
# "spred_id": "1Y2WRuk8mmrZdKXsnmloQTxqw8vdVq0oN26NJBg2USq0",
# "lost_data":"ССФ_СМ.json",
# "log_file": "ССФ_СМ.log"
# },
{
"spred_id": "1-JHj1RSDyCcJ7cphHinfbR8SYZbFoDlzhp72f8oj1Go",
"lost_data":"ССФ_ФН.json",
"log_file": "ССФ_ФН.log"
},
{
"spred_id": "1DKT9Z86LvOYtwBo-lpYZ23vO0U9RvwJF1BxuylaeIew",
"lost_data":"ССФ_Э.json",
"log_file": "ССФ_Э.log"
},
{
"spred_id": "1PRFbQcYXrADoi59Jj5Hn7leVmr5HiCd7j3Z5ltnpKpw",
"lost_data":"ЦК.json",
"log_file": "ЦК.log"
},
# {
# "spred_id": "15Th9tX3gbWC5T2RRxe1TuVcsEu95Zr08qqKrMHYvnts",
# "lost_data":"ЦСР.json",
# "log_file": "ЦСР.log"
# },
{
"spred_id": "1sFLfnuZg6sccsByCAWgJW8ktK0uwTQx6BIEqcl9XPbg",
"lost_data":"Штаб.json",
"log_file": "Штаб.log"
},
{
"spred_id": "1UhB0alaWuiCtSkyYcbcbJuP445btwV_VRr0gDNuUQ1E",
"lost_data":"BAS.json",
"log_file": "BAS.log"
},
{
"spred_id": "1PyUB06AV8JGpPj8Z9Cw_NZqlVXCZi-xRtSG7i7M3cvs",
"lost_data":"BBC.json",
"log_file": "BBC.log"
},
{
"spred_id": "1-Jk_JNWwyx_T81vPlL5ohkPv6MgbWPo-0Cg3_CqeuDo",
"lost_data":"BEST.json",
"log_file": "BEST.log"
},
{
"spred_id": "1zB8VQwJM41MwyIJnp5e5AQILv6oNDNFSXfHVTD2G6Ro",
"lost_data":"HR.json",
"log_file": "HR.log"
},
{
"spred_id": "1WXv9qFUh6RT87vfbZuPUtdxSrLeEi7QUHDOZRE3FCvk",
"lost_data":"ISCRA.json",
"log_file": "ISCRA.log"
},
# {
# "spred_id": "1NZoyoBpehhzwpV-Db8k6Q5YXNH4oy6HfwkuH8sHf0cA",
# "lost_data":"Media.json",
# "log_file": "Media.log"
# },
{
"spred_id": "1-HvTREFw-ozeZYSGm5Zj_y_nfNd8kmLoT1Mn1sB3Q4M",
"lost_data":"UNIAL.json",
"log_file": "UNIAL.log"
},
]

#     pidori(dat["spred_id"], dat["lost_data"], dat["log_file"])
# dat = {"spred_id": "1UhB0alaWuiCtSkyYcbcbJuP445btwV_VRr0gDNuUQ1E",
# "lost_data":"BAS.json",
# "log_file": "BAS.log"
# }
# pidori(dat["spred_id"], dat["lost_data"], dat["log_file"])

for dt in data:
#     print(dt)
    pidori(dt["spred_id"], dt["lost_data"], dt["log_file"])