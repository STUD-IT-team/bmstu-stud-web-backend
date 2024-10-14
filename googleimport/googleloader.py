from googleapiclient.discovery import build
from googleapiclient.http import MediaIoBaseDownload
import os

from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow

SCOPES = ["https://www.googleapis.com/auth/drive.readonly"]

class GoogleFile:
    def __init__(self, id: str, name: str, mimeType: str):
        self.Id = id
        self.Name = name
        self.MimeType = mimeType

class GoogleLoader:
    def __init__(self, creds):
        self.service = build("drive", "v3", credentials=creds)

    def DownloadBlobFile(self, googleFileId, filename):
        req = self.service.files().get_media(
            fileId = googleFileId,
            supportsAllDrives=True,
            )
        with open(filename, 'wb') as saveFile:
            downloader = MediaIoBaseDownload(saveFile, req)
            done = False
            while not done:
                _, done = downloader.next_chunk()
    
    def GetFileInfo(self, googleFileId) -> GoogleFile:
        req = self.service.files().get(
            fileId=googleFileId,
            supportsAllDrives=True,
            ).execute()
        return GoogleFile(
            id=req["id"],
            name=req["name"],
            mimeType=req["mimeType"]
        )
    
    def GetSharedDriveFiles(self, googleDriveId) -> list[GoogleFile] :
        req = (self.service.files()
        .list(q=f"'{googleDriveId}' in parents", supportsAllDrives=True, includeItemsFromAllDrives=True,)
        .execute())
        files = req.get("files", [])
        resFiles = []
        for f in files:
            resFiles.append(GoogleFile(
                id=f['id'],
                name=f['name'],
                mimeType=f['mimeType']
            ))
        return resFiles
    


