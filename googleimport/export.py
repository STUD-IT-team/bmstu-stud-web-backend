from organisation import Organization
from member import Member
from achievement import Achievement
import requests as re
import os
import utils
import random as r

class AuthorizationError(re.HTTPError):
    pass

class UpdatePhotoError(re.HTTPError):
    pass

class Exporter:

    def autoLogin(func):
        def wrapper(self, *args, **kwargs):
            try:
                return func(self, *args, **kwargs)
            except AuthorizationError as e:
                self.login()
                return func(self, *args, **kwargs)
            except:
                raise
        
        return wrapper
            
    def __init__(self, ip, port, adminLogin, adminPassword):
        self.ip = ip
        self.port = port
        self.adminLogin = adminLogin
        self.adminPassword = adminPassword
        self.loginSession = re.Session()
    
    def login(self):
        response = self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/guard/login/",
            json={
                "login": self.adminLogin,
                "password": self.adminPassword
            }
        )

        if response.status_code != 200:
            raise re.HTTPError("Can't login:", response=response)
    
    def logout(self):
        self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/guard/logout/"
        )

    @autoLogin
    def GetRandomDefaultMedia(self) -> int:
        response = self.loginSession.get(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/media/default/"
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 200:
            raise re.HTTPError("Can't get random default media:", response=response)
        return r.choice(response.json()['media'])['id']
    
    @autoLogin
    def ExportPhotoPrivate(self, filepath) -> int:
        if not os.path.exists(filepath):
            raise FileNotFoundError(f"File '{filepath}' not found.")
        
        with open(filepath, 'rb') as file:
            b = file.read()
            name = os.path.basename(filepath)
        
        data = utils.BytesToIntList(b)
        response = self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/media/private/",
            json={
                "name": name,
                "data" : data
            }
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code != 200:
            raise re.HTTPError(f"Can't export photo: {response.status_code}", response=response)
        return response.json()['id']
    
    @autoLogin
    def GetClubIdByName(self, clubName):
        response = self.loginSession.get(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/clubs/",
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 200:
            raise re.HTTPError("Can't get club id by name:", response=response)
        
        clubs = [club for club in response.json()['clubs']]
        for c in clubs:
            if c['name'] == clubName:
                return c['id']
        else:
            raise ValueError(f"Club '{clubName}' not found.")

    @autoLogin
    def ExportOrganisationOnly(self, club: Organization, photoId: int) -> int:
        response = self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/clubs/",
            json={
                "name" : club.Name,
                "short_name" : club.ShortName,
                "description" : club.Description,
                "short_description": club.ShortDescription,
                "type" : club.ClubType,
                "logo_id" : photoId,
                "vk_url" : club.Vk,
                "tg_url" : club.Telegram,
                "parent_id" : 0,
                "orgs" : []
            }
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 200:
            raise re.HTTPError("Can't organisation:", response=response)
    

        return self.GetClubIdByName(club.Name)

    @autoLogin
    def GetMemberIdByName(self, name):
        response = self.loginSession.get(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/members/search/{name}",
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 200:
            raise re.HTTPError("Can't get member id by name:", response=response)
        
        return response.json()['members'][0]['id']

    # default password: qwerty12345678
    # Транслитом Фамилия и инициалы.
    # Если коллизия, то к фамилиии инициалам добавить цифру.
    # На наличие в базе смотреть по имени.
    def ExportMember(self, member: Member, photoId: int) -> int:
        self.logout()
        response = self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/guard/register/",
            json={
                "login": member.GetLogin(),
                "password" : member.GetPassword(),
                "name" : member.GetName(),
                "telegram": member.GetTelegram(),
                "vk": member.GetVk(),
            }
        )
        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 201:
            raise re.HTTPError("Can't export member:", response=response)
        
        self.logout()
        
        id = self.GetMemberIdByName(member.GetName())
        response = self.loginSession.put(
            url = f"{self.ip}:{self.port}/bmstu-stud-web/api/members/{id}",
            json={
                "is_admin": False,
                "login": member.GetLogin(),
                "media_id": photoId,
                "name": member.GetName(),
                "telegram": member.GetTelegram(),
                "vk": member.GetVk()
            }
        )

        if response.status_code != 200:
            raise UpdatePhotoError(f"Can't update photo for member id = {id}", response=response)

        return id
        

    @autoLogin
    def AddOrgsToClub(self, clubID: int, members: dict[int, Member]):
        response = self.loginSession.get(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/clubs/{clubID}"
        )
        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 200:
            raise re.HTTPError("Can't get club info:", response=clubInfo)
        
        clubInfo = response.json()
        orgs = []
        for memberID, member in members.items():
            orgs.append({
                "member_id": memberID,
                "role_name": member.GetRoleName(),
                "role_spec": member.GetRoleSpec()
            })
        
        response = self.loginSession.put(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/clubs/{clubID}",
            json={
                "description": clubInfo["description"],
                "logo_id": clubInfo["logo"]['id'],
                "name": clubInfo["name"],
                "orgs": orgs,
                "parent_id": clubInfo["parent_id"],
                "short_description": clubInfo["short_description"],
                "short_name": clubInfo["short_name"],
                "tg_url": clubInfo["tg_url"],
                "type": clubInfo["type"],
                "vk_url": clubInfo["vk_url"]
            }
        )
        if response.status_code != 200:
            raise re.HTTPError(f"Can't update club info for club id = {clubID}", response=response)
    
        

    @autoLogin
    def ExportClubPhoto(self, photoId: int, clubId: int, refNumber: int):
        response = self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/clubs/media/{clubId}",
            json={
                "photos": [
                {
                    "ref_number": refNumber,
                    "media_id" : photoId
                }
                ]  
            }
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code != 201:
            raise re.HTTPError("Can't export club photo:", response=response)


    @autoLogin
    def ExportAchievement(self, achievement: Achievement, clubId: int):
        response = self.loginSession.post(
            url=f"{self.ip}:{self.port}/bmstu-stud-web/api/feed/encounters/",
            json={
                "club_id": clubId,
                "count" : achievement.count,
                "description" : achievement.description
            }
        )

        if response.status_code == 401:
            raise AuthorizationError("Not authorized.")
        if response.status_code!= 201:
            raise re.HTTPError("Can't export achievement:", response=response)

