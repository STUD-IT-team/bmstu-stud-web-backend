import json
from googleapiclient.discovery import build
import os.path
import utils
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError
from member import Member
from achievement import Achievement
from googlesheet import GoogleSheetRange, GoogleSpreadsheet
from googleloader import GoogleLoader
from organisation import Organization
import googleloader
import googlesheet
import logging

SCOPES = googlesheet.SCOPES + googleloader.SCOPES
SAMPLE_SPREADSHEET_ID = "1lfzZuoui21E78wYmF6fBPXHimgREkx81iSSFF45dOw8"

Settings = json.load(open("settings.json"))
class OrganizationParser:
    def __init__(self, spreadsheetID: str, creds, logger: logging.Logger):
        self.spreadsheetGoogleID = spreadsheetID
        self.creds = creds
        self.spreadsheet = GoogleSpreadsheet(spreadsheetID, creds)
        self.loader = GoogleLoader(creds)
        self.organisation = Organization()
        self.logger = logger
    

    def ParseSpreadsheet(self):
        self.logger.info(f"Start loading spreadsheet id={self.spreadsheetGoogleID}.")


        self.logger.info(f"Parsing sheet {Settings["organisations_table"]}.")
        orgRange = self.spreadsheet.GetTableRange(Settings["organisations_table"],
            f"{Settings["organisation_start_column"]}{Settings["organisation_row"]}",
            f"{Settings["organisation_end_column"]}{Settings["organisation_row"]}")
        self.logger.info(f"Parse {Settings["organisations_table"]} successful.")

        self.organisation.ClubType = orgRange[f"{Settings["club_type_column"]}{Settings["organisation_row"]}"]
        self.organisation.Name = orgRange[f"{Settings['name_column']}{Settings["organisation_row"]}"]
        self.organisation.ShortName = orgRange[f"{Settings['short_name_column']}{Settings["organisation_row"]}"]
        self.organisation.ShortDescription = orgRange[f"{Settings['short_description_column']}{Settings["organisation_row"]}"]
        self.organisation.Description = orgRange[f"{Settings['description_column']}{Settings["organisation_row"]}"]
        self.organisation.PhotoFolderGoogleID = utils.ParseSharedFolderID(orgRange[f"{Settings['photos_url_column']}{Settings["organisation_row"]}"])
        self.organisation.LogoGoogleID = utils.ParseSharedFileID(orgRange[f"{Settings['logo_url_column']}{Settings["organisation_row"]}"])
        
        self.logger.info(f"Parsing sheet {Settings["urls_table"]}.")
        urlRange = self.spreadsheet.GetTableRange(Settings["urls_table"],
            f"{Settings["urls_column"]}{Settings["vk_url_row"]}",
            f"{Settings["urls_column"]}{Settings["telegram_url_row"]}")
        
        self.logger.info(f"Parse {Settings["urls_table"]} successful.")
        
        self.organisation.Telegram = urlRange[f"{Settings["urls_column"]}{Settings["telegram_url_row"]}"]
        self.organisation.Vk = urlRange[f"{Settings["urls_column"]}{Settings["vk_url_row"]}"]

        self.logger.debug(f"Parsed organisation {self.organisation.__dict__}.")

        self.ParseMembers()
        self.ParseAchievements()
    
    def ParseMembers(self):
        self.logger.info(f"Parsing members of organisation {self.organisation.Name}, id={self.spreadsheetGoogleID}.")
        self.logger.info(f"Parsing {Settings["members_table"]}.")

        rows, _ = self.spreadsheet.GetTableSizes(Settings['members_table'])
        startRow = Settings['members_start_row']
        data = self.spreadsheet.GetTableRange(Settings['members_table'], f"{Settings['members_start_column']}{startRow}", f"{Settings['members_end_column']}{rows}")
        self.logger.info(f"Parse {Settings['members_table']} successful.")
        self.logger.info(f"Got {data.NonEmptyRows()} member rows.")

        for i in range(data.NonEmptyRows()):
            try:
                row = i + startRow
                self.organisation.AddMember(Member(
                            name=data[f"{Settings['members_name_column']}{row}"],
                            photoUrl=data[f"{Settings['members_photo_url_column']}{row}"],
                            telegram=data[f"{Settings['members_telegram_url_column']}{row}"],
                            vk=data[f"{Settings['members_vk_url_column']}{row}"],
                            roleName=data[f"{Settings['members_role_name_column']}{row}"],
                            roleSpec=data[f"{Settings['members_role_spec_column']}{row}"],
                            roleField=data[f"{Settings['members_role_field_column']}{row}"],
                        )
                    )
                self.logger.debug(f"Parsed member {self.organisation.Members[-1].__dict__}")
            except BaseException as e:
                self.logger.error(f"Error parsing member at row {row}: {e}")
                self.logger.error(f"Do not include member at row {row}.")

    
    def ParseAchievements(self):
        self.logger.info(f"Parsing achievements of organisation {self.organisation.Name}, id={self.spreadsheetGoogleID}.")
        self.logger.info(f"Parsing {Settings["achievments_table"]}.")

        rows, _ = self.spreadsheet.GetTableSizes(Settings['achievments_table'])
        startRow = Settings['achievments_start_row']
        data = self.spreadsheet.GetTableRange(Settings['achievments_table'], f"{Settings['achievments_start_column']}{startRow}", f"{Settings['achievments_end_column']}{rows}")
        self.logger.info(f"Parse {Settings['achievments_table']} successful.")
        self.logger.info(f"Got {data.NonEmptyRows()} achievements rows.")

        for i in range(data.NonEmptyRows()):
            try:
                row = i + startRow
                self.organisation.AddAchievement(Achievement(
                            count=data[f"{Settings['achievments_count_column']}{row}"],
                            description=data[f"{Settings['achievments_description_column']}{row}"]
                        )
                    )
                self.logger.debug(f"Parsed member {self.organisation.Achievements[-1].__dict__}")
            except BaseException as e:
                self.logger.error(f"Error parsing achievement at row {row}: {e}")
                self.logger.error(f"Do not include achievement at row {row}.")
    
    def DownloadLogo(self):
        self.logger.info(f"Downloading organisation {self.organisation.Name} logo id={self.organisation.LogoGoogleID}...")
        self.CreateOrganisationSubdir(Settings["logo_dir"])
        if self.organisation.LogoGoogleID == None:
            self.logger.info("No logo found.")
            return

        fileInfo = self.loader.GetFileInfo(self.organisation.LogoGoogleID)
        self.logger.info(f"Logo file: {fileInfo}.")
        logoFileName = os.path.join(self.GetOrganisationSubdir(Settings["logo_dir"]), fileInfo.Name)
        self.logger.info(f"Downloading logo file id={self.organisation.LogoGoogleID} to {logoFileName}")
        try:
            self.loader.DownloadBlobFile(self.organisation.LogoGoogleID, logoFileName)
        except BaseException as e:
            self.logger.error(f"Error downloading logo: {e}")
        self.logger.info(f"Logo downloaded to {logoFileName}.")
        self.organisation.OsLogoPath = logoFileName
    
    def DownloadOrganizationPhotos(self):
        self.logger.info(f"Downloading organisation {self.organisation.Name} photos...")
        self.CreateOrganisationSubdir(Settings["photos_dir"])
        self.organisation.OsPhotosPath = self.GetOrganisationSubdir(Settings["photos_dir"])

        try:
            files = self.loader.GetSharedDriveFiles(self.organisation.PhotoFolderGoogleID)
        except BaseException as e:
            self.logger.error(f"Error getting photo files: {e}")
            return
        self.logger.info(f"Got {len(files)} photo files.")
        for f in files:
            if f.MimeType == 'application/vnd.google-apps.folder':
                continue
            self.logger.info(f"Photo file: {f}.")
            photoFileName = os.path.join(self.GetOrganisationSubdir(Settings['photos_dir']), f.Name)
            self.logger.info(f"Downloading photo id={f.Id} to {photoFileName}")
            try:
                self.loader.DownloadBlobFile(f.Id, photoFileName)
            except BaseException as e:
                self.logger.error(f"Error downloading photo: {e}")
            self.logger.info(f"Photo downloaded to {photoFileName}.")
    
    def DownloadMemberPhotos(self):
        self.logger.info(f"Downloading organisation {self.organisation.Name} members photos...")
        self.CreateOrganisationSubdir(Settings["member_photos_dir"])

        for member in self.organisation.Members:
            fId = member.GetPhotoGoogleId()
            if fId == None:
                self.logger.info(f"No photo found for member {member.GetName()}.")
                continue
            self.logger.info(f"Downloading member {member.GetName()} photo id={fId}")
            try:
                fileInfo = self.loader.GetFileInfo(fId)
            except BaseException as e:
                self.logger.error(f"Error getting member photo file: {e}")
                continue
            self.logger.info(f"Member {member.GetName()} photo file: {fileInfo}.")
            photoFileName = os.path.join(self.GetOrganisationSubdir(Settings['member_photos_dir']), fileInfo.Name)
            self.logger.info(f"Downloading member photo id={fId} to {photoFileName}")
            try:
                self.loader.DownloadBlobFile(fId, photoFileName)
            except BaseException as e:
                self.logger.error(f"Error downloading member photo: {e}")
            self.logger.info(f"Member {member.GetName()} photo downloaded to {photoFileName}.")
            member.SetOsPhotoPath(photoFileName)

    def CreateOrganisationDirectory(self):
        if not os.path.exists(Settings["data_dir"]):
            self.logger.info("Creating organisation directory {0}".format(Settings["data_dir"]))
            os.mkdir(Settings["data_dir"])
        if not os.path.exists(self.GetOrganisationDir()):
            self.logger.info("Creating organisation subdir {0}".format(self.GetOrganisationDir()))
            os.mkdir(self.GetOrganisationDir())
    
    def CreateOrganisationSubdir(self, subdir):
        self.CreateOrganisationDirectory()
        if not os.path.exists(self.GetOrganisationSubdir(subdir)):
            self.logger.info("Creating {0} subdir {1}".format(subdir, self.GetOrganisationSubdir(subdir)))
            os.mkdir(self.GetOrganisationSubdir(subdir))
    
    def GetOrganisationDir(self):
        return os.path.join(os.path.join(Settings["data_dir"], self.organisation.Name))

    def GetOrganisationSubdir(self, subdir):
        return os.path.join(os.path.join(Settings["data_dir"], self.organisation.Name, subdir))


if __name__ == "__main__":
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
    logging.basicConfig(level=logging.DEBUG, filename="info.log", filemode="w")
    # logging.basicConfig(level=logging.WARNING, filename="warning.log", filemode="w")

    org = OrganizationParser(SAMPLE_SPREADSHEET_ID, creds, l)
    org.ParseSpreadsheet()
    # print(org.organisation.__dict__)
    org.DownloadLogo()
    org.DownloadOrganizationPhotos()
    org.DownloadMemberPhotos()