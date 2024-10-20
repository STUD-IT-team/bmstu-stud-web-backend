from organisationparser import OrganizationParser, SCOPES
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from export import Exporter
from organizationexport import OrganizationExporter
import os
import logging

SAMPLE_SPREADSHEET_ID = "1lfzZuoui21E78wYmF6fBPXHimgREkx81iSSFF45dOw8"
LOST_DATA_JSON = 'BAS.json'
LOG_FILE = 'BAS.log'
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



