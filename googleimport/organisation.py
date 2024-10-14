import os.path
import utils
from member import Member
from achievement import Achievement

class Organization:
    def __init__(self):
        self.ClubId = None
        self.ClubType = ""
        self.Name = ""
        self.ShortName = ""
        self.ShortDescription = ""
        self.Description = ""
        self.Telegram = ""
        self.Vk = ""
        self.PhotoFolderGoogleID = ""
        self.LogoGoogleID = None
        self.OsLogoPath = None
        self.OsPhotosPath = None
        self.Members = []
        self.Achievements = []
    
    def AddMember(self, member: Member):
        self.Members.append(member)
    
    def AddAchievement(self, achievement: Achievement):
        self.Achievements.append(achievement)
    
    def DeleteMember(self, member: Member):
        self.Members.remove(member)
    
    def DeleteAchievement(self, achievement: Achievement):
        self.Achievements.remove(achievement)

    
