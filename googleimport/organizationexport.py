import logging
from organisation import Organization
from member import Member
from export import Exporter
import os
import json
import utils

DEFAULT_PASSWORD = 'qwerty12345678'

class OrganizationExporter:
    def __init__(self, logger: logging.Logger, organization : Organization, exporter: Exporter, lostDataFileName: str):
        self.logger = logger
        self.organization = organization
        self.exporter = exporter
        self.lostData = dict()
        self.lostDataFileName = lostDataFileName
    
    def StartExport(self):
        clubId = self.exportOrganisation()
        memberMap = self.exportMembers()
        self.exportClubOrgs(memberMap, clubId)
        self.exportAchievements(clubId)
        self.exportClubPhoto(clubId)
        self.saveLostData()
        

    
    def exportOrganisation(self):
        self.logger.info(f"Start export club '{self.organization.Name}'")
        logo = None
        if self.organization.OsLogoPath != None:
            try:
                print(self.organization.OsLogoPath)
                logo = self.exporter.ExportPhotoPrivate(self.organization.OsLogoPath)
            except Exception as e:
                self.logger.error(f"Error exporting logo '{self.organization.Name}': {e}")
                
        
        if logo == None:
            if self.organization.OsLogoPath == None:
                self.logger.info("No logo path specified for club '{self.organization.Name}. Using default.")
            else:
                self.logger.info(f"Using default Media for club '{self.organization.Name}")
            try:
                logo = self.exporter.GetRandomDefaultMedia()
            except BaseException as e:
                self.logger.error(f"Error getting default media: {str(e)}")
                raise

        

        self.logger.info(f"Got logo for '{self.organization.Name}' with ID: {logo}")
        try:
            id = self.exporter.ExportOrganisationOnly(self.organization, logo)
        except Exception as e:
            self.logger.error(f"Error posting club: {str(e)}")
            raise
        
        return id
            
    def exportAchievements(self, clubId):
        self.logger.info(f"Exporting achievements for '{self.organization.Name}'")
        self.lostData['achievements'] = []
        countSuccess = 0
        countFailed = 0
        for a in self.organization.Achievements:
            try:
                self.exporter.ExportAchievement(a, clubId)
                countSuccess += 1
            except BaseException as e:
                self.logger.error(f"Error exporting achievement: {a.__dict__}.")
                countFailed += 1
                self.lostData['achievments'].append(a.ToDict())
        
        self.logger.info(f"Done exporting achievements: success: {countSuccess}, failed: {countFailed}")
        if countFailed != 0:
            self.logger.error(f"Lost some achievement. Search in lostdata json")
    
    def exportMembers(self) -> dict[int, Member]:
        self.logger.info(f"Exporting members for '{self.organization.Name}'")
        self.lostData['members'] = []
        self.lostData['member_photos'] = []
        memberMap = dict()
        countSuccess = 0
        countFailed = 0
        for m in self.organization.Members:
            m.SetLogin(utils.CreateLoginFromName(m.GetName()))
            m.SetPassword(DEFAULT_PASSWORD)
            try:
                memberId = self.exporter.GetMemberIdByName(m.GetName())
                memberMap[memberId] = m
                countSuccess += 1
                self.logger.info(f"Found member {m.GetName()} in database.")
                continue

            except Exception as e:
                self.logger.info(f"Not found member for '{m.GetName()}': {str(e)}")
            photoPath = m.GetOsPhotoPath()
            photoId = None
            if photoPath != None:
                try:
                    photoId = self.exporter.ExportPhotoPrivate(photoPath)
                except Exception as e:
                    self.logger.error(f"Error exporting photo for member '{m.GetName()}': {str(e)}")
                    self.lostData['member_photos'].append(photoPath)

            if photoId == None:
                self.logger.info(f"No photo found for member '{m.GetName()}'. Using default photo")
                try:
                    photoId = self.exporter.GetRandomDefaultMedia()
                except BaseException as e:
                    self.logger.error(f"Error getting default media: {str(e)}")
                    self.lostData['members'].append(m.ToDict())
                    countFailed += 1
                    continue
            
            try:
                mid = self.exporter.ExportMember(m, photoId)
                self.logger.info(f"Member '{m.GetName()}' exported")
                memberMap[mid] = m
            except Exception as e:
                self.logger.error(f"Error exporting member: {m.ToDict()}.")
                self.lostData['members'].append(m.ToDict())
                countFailed += 1
                continue
        
        self.logger.info(f"Done exporting members: success: {countSuccess}, failed: {countFailed}")
        if countFailed!= 0:
            self.logger.error(f"Lost some members. Search in lostdata json")
        return memberMap

    def exportClubOrgs(self, memberMap: dict[int, Member], clubId):
        self.logger.info(f"Exporting club organizators for '{self.organization.Name}'")
        try:
            self.exporter.AddOrgsToClub(clubId, memberMap)
            self.logger.info(f"Club organizations exported for '{self.organization.Name}'")
        except Exception as e:
            self.logger.error(f"Error adding club organizators: {str(e)}")
            self.lostData['club_orgs'] = {
                "club_id" : clubId,
                "members" : [
                    {
                        "member_id":mid,
                        "role_name": member.GetRoleName(),
                        "role_spec": member.GetRoleSpec()
                    }
                    for mid, member in memberMap.items()
                ]
            }

    def exportClubPhoto(self, clubId):
        countSuccess = 0
        countFailed = 0

        self.lostData['club_photos'] = []
        self.logger.info(f"Exporting club photo for '{self.organization.Name}'")
        photoPath = self.organization.OsPhotosPath
        if photoPath == None or not os.path.isdir(photoPath):
            self.logger.error(f"No photo path specified for club '{self.organization.Name}'.")
            return
        photoPaths = os.listdir(photoPath)
        if len(photoPaths) == 0:
            self.logger.error(f"No photos found in '{photoPath}'.")
            return
        ref_num = 1
        for photo in photoPaths:
            curPath = os.path.join(photoPath, photo)
            try:
                photoId = self.exporter.ExportPhotoPrivate(curPath)
                self.exporter.ExportClubPhoto(photoId, clubId, ref_num)
                self.logger.info(f"Photo '{photo}' exported for '{self.organization.Name}'")
                ref_num +=1
            except Exception as e:
                self.logger.error(f"Error exporting photo '{photo}': {str(e)}")
                self.lostData['club_photos'].append(photo)
        self.logger.info(f"Done exporting club photo for '{self.organization.Name}'")
        self.logger.info(f"Photo '{self.organization.Name}: success: {countSuccess}, failed: {countFailed}")
        if countFailed!= 0:
            self.logger.error(f"Lost some photos. Search in lostdata json")
         
    def saveLostData(self):
        try:
            with open(self.lostDataFileName, 'w') as outfile:
                json.dump(self.lostData, outfile, indent=4)
        except Exception as e:
            self.logger.error(f"Error saving lost data: {str(e)}. Printing in log:")
            self.logger.error(self.lostData)

        
                 

            

                
            

                    

                

        
            
