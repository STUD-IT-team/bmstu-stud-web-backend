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
        self.login = None
        self.password = None
    
    def GetPhotoGoogleId(self):
        return self.photoGoogleId
    
    def SetOsPhotoPath(self, path):
        self.osPhotoPath = path
    
    def GetOsPhotoPath(self):
        return self.osPhotoPath
    
    def GetName(self):
        return self.name

    def GetTelegram(self):
        return self.telegram
    
    def GetVk(self):
        return self.vk
    
    def GetRoleName(self):
        return self.roleName
    
    def GetRoleSpec(self):
        return self.roleSpec
    
    def GetRoleField(self):
        return self.roleField
    
    def GetLogin(self):
        return self.login
    
    def GetPassword(self):
        return self.password
    
    def SetLogin(self, login):
        self.login = login
    
    def SetPassword(self, password):
        self.password = password
    
    def ToDict(self):
        return {
            'name': self.name,
            'password': self.password,
            'role_name' : self.roleName,
            'role_spec' : self.roleSpec,
            'role_field' : self.roleField,
            'login': self.login,
            'photo_path': self.osPhotoPath,
            'photo_id': self.photoGoogleId,
            'telegram': self.telegram,
            'vk': self.vk
        }