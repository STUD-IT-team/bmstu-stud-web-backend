from googleapiclient.discovery import build
import os.path
import utils
from google.auth.transport.requests import Request
from google.oauth2.credentials import Credentials
from google_auth_oauthlib.flow import InstalledAppFlow
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError

SCOPES = ["https://www.googleapis.com/auth/spreadsheets.readonly"]
ALPHABET = {chr(i): i - ord('A') + 1 for i in range(ord('A'), ord('Z') + 1)}
REVERSE_ALPHABET = {i + 1: chr(i + ord('A')) for i in range(26)}
class GoogleCellAddress:
    def __init__(self, addr : str):
        addr = addr.upper()
        self.column = 0
        self.row = 0
        columnPart = ""
        for i in range(len(addr)):
            if addr[i] in ALPHABET:
                columnPart += addr[i]
            else:
                self.row = int(addr[i:])
                break
        pow = 26 ** (len(columnPart) - 1)
        for el in columnPart:
            self.column += pow * ALPHABET[el]
            pow //= 26

    
    def ToGoogleCell(self):
        column = self.column
        addr = ""
        while column > 0:
            if column % 26 == 0:
                addr = REVERSE_ALPHABET[26] + addr
                column //= 26
                column -= 1
            else:
                addr = REVERSE_ALPHABET[column % 26] + addr
                column //= 26
        addr += str(self.row)
        return addr
    
    def GetColumn(self):
        return self.column
    
    def GetRow(self):
        return self.row
    
    @staticmethod
    def Difference(cell1, cell2):
        return cell1.column - cell2.column, cell1.row - cell2.row

    def __str__(self):
        return f"{'{'}Column: {self.column}, Row: {self.row}.{'}'}"

class GoogleSheetRange:
    def __init__(self, values, cellStart: GoogleCellAddress, cellEnd: GoogleCellAddress):
        self. valuesMatrix = values
        if values == None:
            raise ValueError("values must be specified")
        self.cellStart = cellStart
        self.cellEnd = cellEnd
        self.cols, self.rows = GoogleCellAddress.Difference(self.cellEnd, self.cellStart)
        self.cols += 1
        self.rows += 1
        if self.cols <= 0 or self.rows <= 0:
            raise ValueError("Incorrect range values")

    def Rows(self):
        return self.rows
    
    def Cols(self):
        return self.cols
    
    def NonEmptyRows(self):
        return len(self.valuesMatrix)
    
    def GetCell(self, address: str) -> str:
        googleAddress = GoogleCellAddress(address)
        col, row = GoogleCellAddress.Difference(googleAddress, self.cellStart)
        if row < 0 or row >= self.rows or col < 0 or col >= self.cols:
            raise IndexError(f"Address out of range: {address}")
        if row >= len(self.valuesMatrix):
            return ""
        if col >= len(self.valuesMatrix[row]):
            return ""
        return self.valuesMatrix[row][col]
    
    def __getitem__(self, index: str):
        return self.GetCell(index)

class GoogleSpreadsheet:
    def __init__(self, spreadsheetId, creds):
        self.service = build("sheets", "v4", credentials=creds)
        self.spreadsheetGoogleID = spreadsheetId
    
    def GetTableSizes(self, tableName: str) -> tuple[int, int]:
        spreadsheets = self.service.spreadsheets().get(spreadsheetId=self.spreadsheetGoogleID).execute()
        for spreadsheet in spreadsheets['sheets']:
            if spreadsheet['properties']['title'] == tableName:
                sheet = spreadsheet
                break
        else:
            raise ValueError(f"No {tableName} table found.")
        
        rowCount = sheet['properties']['gridProperties']['rowCount']
        colCount = sheet['properties']['gridProperties']['columnCount']
        return rowCount, colCount
    
    def GetTableRange(self, tableName: str, cellStart: str, cellEnd: str):
        cellS = GoogleCellAddress(cellStart)
        cellE = GoogleCellAddress(cellEnd)
        if any([el < 0 for el in GoogleCellAddress.Difference(cellE, cellS)]):
            raise ValueError(f"invalid range {cellS.ToGoogleCell()}:{cellE.ToGoogleCell()}")
        
        dataRange = f"{tableName}!{cellS.ToGoogleCell()}:{cellE.ToGoogleCell()}"
        
        data = self.service.spreadsheets().values().get(
            spreadsheetId=self.spreadsheetGoogleID,
            range=dataRange).execute()
        
        if data == None:
            raise RuntimeError(f"can't get table {dataRange} from spreadsheet {self.spreadsheetGoogleID}")

        return GoogleSheetRange(values=data['values'], cellStart=cellS, cellEnd=cellE)



    



        
        


