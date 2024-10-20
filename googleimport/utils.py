import transliterate as tr


def ParseSharedFolderID(url: str) -> str:
    """Parses the Shared Folder ID from the Google Drive URL."""
    folderIndex = 0
    splitURL = url.split('/')
    while folderIndex < len(splitURL) - 1:
        if splitURL[folderIndex] == 'folders':
            return splitURL[folderIndex + 1]
        folderIndex += 1
    raise ValueError('Could not parse Shared Folder ID: no folders found')
    
def ParseSharedFileID(url: str) -> str:
    """Parses the Shared File ID from the Google Drive URL."""
    fileIndex = 0
    splitURL = url.split('/')
    while fileIndex < len(splitURL) - 1:
        if splitURL[fileIndex] == 'file':
            return splitURL[fileIndex + 2]
        fileIndex += 1
    raise ValueError('Could not parse Shared File ID: no files found')

def BytesToIntList(b : bytes) -> list[int]:
    """Converts a byte array to a list of integers(every byte to one integer)."""
    return [int(el) for el in b]

def CreateLoginFromName(name : str) -> str:
    nameParts = name.split()
    login = ""
    loginParts = []

    loginParts.append(tr.translit(nameParts[0], 'ru', reversed=True))
    for part in nameParts[1:]:
        loginParts.append(tr.translit(part, 'ru', reversed=True)[0])
    login = '_'.join(loginParts)
    return login.lower()