import utils
class Member:
    def __init__(self, name, photoUrl, telegram, vk, roleName, roleSpec, roleField):
        try:
            self.photoGoogleId = utils.ParseSharedFileID(photoUrl)
        except ValueError: # Incorrect photo url
            self.photoGoogleId = None
        self.osPhotoPath = None
        self.name = name
        self.telegram = telegram
        self.vk = vk
        self.roleName = roleName
        self.roleSpec = roleSpec
        self.roleField = roleField
    
    def GetPhotoGoogleId(self):
        return self.photoGoogleId
    
    def SetOsPhotoPath(self, path):
        self.osPhotoPath = path
    
    def GetName(self):
        return self.name